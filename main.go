// This program is continuously WIP.
// The goal is to filter each Packages package (which contains the list of all packages available in the ParrotOS
// repository) for each architecture and each branch.

package main

import (
	"log"
	"package-filter/fileops"
	"package-filter/filter"
)

// Here the three phases of the program are carried out:
// 1. Download the Packages for all the architectures.
// 2. Filter and return them as JSON files (~70 mb).
func main() {
	f := new(filter.Package)

	// Create temporary dir called "packages"
	fileops.Mkdir()

	// Start the downloading phase
	log.Println("[info] Downloading packages...")

	// Use the DownloadPackages function to download Packages for each branch and architecture
	fileops.GetJSONPackages()

	// The filter phase begins.
	log.Println("[info] Filtering...")
	f.Parser()

	// The packages folder which contains Packages for each architecture
	// is deleted as it is no longer useful.
	fileops.Rmdirs()

	log.Println("[info] All Packages files deleted.")
	log.Println("[success] Check the json folder.")
}
