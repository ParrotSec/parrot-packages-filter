// This program is continuously WIP.
// The goal is to filter each Packages package (which contains the list of all packages available in the ParrotOS
// repository) for each architecture and each branch.

package main

import (
	"log"
	"net/http"
	"os"
	filter "package-filter/src"
)

// These functions take care of showing the respective .json files contained in the json folder on the browser
func getAMD64Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "json/amd64-packages.json")
}

func getARM64Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "json/arm64-packages.json")
}

func getARMHFPackages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "json/armhf-packages.json")
}

func geti386Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "json/i386-packages.json")
}

func handleFunctions() {
	http.HandleFunc("/packages/main/amd64/", getAMD64Packages)
	http.HandleFunc("/packages/main/arm64/", getARM64Packages)
	http.HandleFunc("/packages/main/armhf/", getARMHFPackages)
	http.HandleFunc("/packages/main/i386/", geti386Packages)
}

// Here the three phases of the program are carried out:
// 1. Download the Packages for all the architectures.
// 2. Filter and return them as JSON files.
// 3. Start the HTTP server to show them in the browser.
func main() {
	const port = "8080"
	const url = "https://download.parrot.sh/parrot/dists/parrot"

	f := new(filter.Package)

	// Create temporary dir called "packages"
	errMkdir := os.Mkdir("packages", os.ModePerm)
	if errMkdir != nil {
		log.Fatal(errMkdir)
	}

	// Start downloading packages for all architectures available
	log.Println("Downloading packages...")
	arch := []string{
		"amd64",
		"arm64",
		"armhf",
		"i386",
	}

	// Use the DownloadPackages function to download Packages for each architecture
	for i := range arch {
		errDownload := filter.DownloadPackages(
			"packages/"+arch[i]+"-packages",
			url+"/main/binary-"+arch[i]+"/Packages")

		if errDownload != nil {
			log.Fatal(errDownload)
		}
	}

	// The filter phase begins.
	log.Println("[!] Filtering...")
	f.Parser()

	// The packages folder which contains Packages for each architecture
	// is deleted as it is no longer useful.
	errRmdir := os.RemoveAll("packages")
	if errRmdir != nil {
		log.Fatal(errRmdir)
	}
	log.Println("Deleted all Packages files.")

	// The HTTP server to show the JSON files is started.
	handleFunctions()
	log.Printf("[!] Starting HTTP server to serve json files at port: %s\n", port)
	errHttp := http.ListenAndServe(":"+port, nil)
	if errHttp != nil {
		log.Fatal(errHttp)
	}
}
