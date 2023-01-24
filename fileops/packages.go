package fileops

import (
	"errors"
	"log"
	"os"
	c "package-filter/constants"
	"package-filter/network"
)

func GetJSONPackages() {
	const url = "https://deb.parrot.sh/parrot/dists/parrot"

	for b := range c.Branch {

		errBranchDir := os.Mkdir("packages/"+c.Branch[b], os.ModePerm)
		if errBranchDir != nil {
			log.Fatal(errBranchDir)
		}

		for a := range c.Arch {
			// Check and if not exists create a new JSON folder where to store each new JSON file for each branch and architecture
			jsonPath := "json/packages/" + c.Branch[b] + "/" + c.Arch[a] + "/"

			if _, errStatJson := os.Stat(jsonPath); errors.Is(errStatJson, os.ErrNotExist) {

				errJsonFolder := os.MkdirAll(jsonPath, os.ModePerm)
				if errJsonFolder != nil {
					log.Fatal(errJsonFolder)
				}

			}

			// Start downloading packages for all branches and architectures available
			errDownload := network.DownloadPackages(
				"packages/"+c.Branch[b]+"/"+c.Arch[a],
				url+"/"+c.Branch[b]+"/binary-"+c.Arch[a]+"/Packages",
			)
			if errDownload != nil {
				log.Fatal(errDownload)
			}

		}
	}
}
