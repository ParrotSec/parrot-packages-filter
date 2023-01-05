package internal

import (
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadPackages This function just manages the download of each Packages file.
func DownloadPackages(filepath string, url string) (err error) {
	out, errCreate := os.Create(filepath)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	defer out.Close()

	res, errGet := http.Get(url)
	if errGet != nil {
		log.Fatal(errGet)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", res.Status)
	}

	_, errCopy := io.Copy(out, res.Body)
	if errCopy != nil {
		log.Fatal(errCopy)
	}

	return nil
}
