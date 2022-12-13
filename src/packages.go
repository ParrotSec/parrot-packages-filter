package src

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetPackages(filepath string, url string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Bad status: %s", res.Status)
	}

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}
