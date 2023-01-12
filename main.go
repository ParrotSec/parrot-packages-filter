// This program is continuously WIP.
// The goal is to filter each Packages package (which contains the list of all packages available in the ParrotOS
// repository) for each architecture and each branch.

package main

import (
	"errors"
	"log"
	"os"
	filter "package-filter/internal"
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

// Here the three phases of the program are carried out:
// 1. Download the Packages for all the architectures.
// 2. Filter and return them as JSON files.
func main() {
	const url = "https://download.parrot.sh/parrot/dists/parrot"

	f := new(filter.Package)

	// Create temporary dir called "packages"
	errMkdir := os.Mkdir("packages", os.ModePerm)
	if errMkdir != nil {
		log.Fatal(errMkdir)
	}

	// Start the downloading phase
	log.Println("[info] Downloading packages...")

	// Use the DownloadPackages function to download Packages for each branch and architecture
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
			errDownload := filter.DownloadPackages(
				"packages/"+branch[b]+"/"+arch[a],
				url+"/"+branch[b]+"/binary-"+arch[a]+"/Packages")

			if errDownload != nil {
				log.Fatal(errDownload)
			}

		}
	}

	// The filter phase begins.
	log.Println("[info] Filtering...")
	f.Parser()

	// The packages folder which contains Packages for each architecture
	// is deleted as it is no longer useful.
	errRmdirs := os.RemoveAll("packages")
	if errRmdirs != nil {
		log.Fatal(errRmdirs)
	}

	log.Println("[info] All Packages files deleted.")
	log.Println("[success] Check the json folder.")
}
