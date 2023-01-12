package internal

import (
	"errors"
	"log"
	"os"
)

var branch = [3]string{
	"contrib",
	"main",
	"non-free",
}

var arch = [4]string{
	"amd64",
	"arm64",
	"armhf",
	"i386",
}

func GetJsonPackages() {
	const url = "https://deb.parrot.sh/parrot/dists/parrot"

	for b := range branch {

		errBranchDir := os.Mkdir("packages/"+branch[b], os.ModePerm)
		if errBranchDir != nil {
			log.Fatal(errBranchDir)
		}

		for a := range arch {
			// Check and if not exists create a new JSON folder where to store each new JSON file for each branch and architecture
			jsonPath := "json/packages/" + branch[b] + "/" + arch[a] + "/"

			if _, errStatJson := os.Stat(jsonPath); errors.Is(errStatJson, os.ErrNotExist) {

				errJsonFolder := os.MkdirAll(jsonPath, os.ModePerm)
				if errJsonFolder != nil {
					log.Fatal(errJsonFolder)
				}

			}

			// Start downloading packages for all branches and architectures available
			errDownload := DownloadPackages(
				"packages/"+branch[b]+"/"+arch[a],
				url+"/"+branch[b]+"/binary-"+arch[a]+"/Packages",
			)
			if errDownload != nil {
				log.Fatal(errDownload)
			}

		}
	}
}
