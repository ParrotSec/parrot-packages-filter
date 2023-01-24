package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadPackages This function just manages the download of each Packages file.
func DownloadPackages(filepath string, url string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer out.Close()

	// Make the GET request
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer res.Body.Close()

	// Check the status code
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", res.StatusCode)
	}

	// Copy the response body to the file
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return fmt.Errorf("error copying response body to file: %w", err)
	}

	return nil
}
