package internal

import (
	"io"
	"log"
	"net/http"
	"os"
)

// DownloadPackages This function just manages the download of each Packages file.
func DownloadPackages(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Bad status: %s", res.Status)
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
