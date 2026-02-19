package transfer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DownloadFile downloads a file from a URL to the specified destination
func DownloadFile(url, destination string) string {
	// Validate URL
	if !isValidURL(url) {
		return fmt.Sprintf("[!] Invalid URL: %s\n", url)
	}

	// Create destination directory if it doesn't exist
	destDir := filepath.Dir(destination)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Sprintf("[!] Failed to create destination directory: %v\n", err)
	}

	// Check if file already exists
	if _, err := os.Stat(destination); err == nil {
		return fmt.Sprintf("[!] File already exists: %s\n"+
			"    Remove it first or choose a different destination.\n", destination)
	}

	// Start download
	output := fmt.Sprintf("[*] Downloading from: %s\n", url)
	output += fmt.Sprintf("[*] Destination: %s\n", destination)

	startTime := time.Now()

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 5 * time.Minute, // 5 minute timeout for large files
	}

	// Make GET request
	resp, err := client.Get(url)
	if err != nil {
		return output + fmt.Sprintf("[!] Download failed: %v\n", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return output + fmt.Sprintf("[!] Server returned status: %d %s\n",
			resp.StatusCode, resp.Status)
	}

	// Get file size if available
	fileSize := resp.ContentLength
	if fileSize > 0 {
		output += fmt.Sprintf("[*] File size: %s\n", formatBytes(fileSize))
	}

	// Create destination file
	outFile, err := os.Create(destination)
	if err != nil {
		return output + fmt.Sprintf("[!] Failed to create file: %v\n", err)
	}
	defer outFile.Close()

	// Download with progress tracking
	written, err := io.Copy(outFile, resp.Body)
	if err != nil {
		// Clean up partial file on error
		os.Remove(destination)
		return output + fmt.Sprintf("[!] Download interrupted: %v\n", err)
	}

	duration := time.Since(startTime)

	// Verify file was created
	fileInfo, err := os.Stat(destination)
	if err != nil {
		return output + fmt.Sprintf("[!] File verification failed: %v\n", err)
	}

	// Success message
	output += fmt.Sprintf("[+] Download complete!\n")
	output += fmt.Sprintf("    Bytes written: %s\n", formatBytes(written))
	output += fmt.Sprintf("    Final size: %s\n", formatBytes(fileInfo.Size()))
	output += fmt.Sprintf("    Duration: %s\n", duration.Round(time.Millisecond))

	if duration.Seconds() > 0 {
		speed := float64(written) / duration.Seconds()
		output += fmt.Sprintf("    Speed: %s/s\n", formatBytes(int64(speed)))
	}

	return output
}

// DownloadFileWithProgress downloads a file with detailed progress reporting
// This is an alternative for large files (not used by default to keep output clean)
func DownloadFileWithProgress(url, destination string) string {
	// Similar to DownloadFile but with progress bar
	// Implementation omitted for simplicity
	// Could be added as "download --progress <url> <dest>"
	return DownloadFile(url, destination)
}

// isValidURL performs basic URL validation
func isValidURL(url string) bool {
	// Basic validation - should start with http:// or https://
	if len(url) < 10 {
		return false
	}

	hasHTTP := len(url) >= 7 && url[:7] == "http://"
	hasHTTPS := len(url) >= 8 && url[:8] == "https://"

	return hasHTTP || hasHTTPS
}

// formatBytes converts bytes to human-readable format
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}

	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	// KB, MB, GB, TB
	units := []string{"KB", "MB", "GB", "TB"}
	if exp >= len(units) {
		exp = len(units) - 1
	}

	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}

// VerifyDownload checks if a downloaded file exists and matches expected size
func VerifyDownload(filepath string, expectedSize int64) (bool, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}

	if expectedSize > 0 && info.Size() != expectedSize {
		return false, fmt.Errorf("size mismatch: expected %d, got %d",
			expectedSize, info.Size())
	}

	return true, nil
}
