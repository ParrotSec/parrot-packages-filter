package main

import (
	"log"
	"net/http"
	"os"
	filter "package-filter/src"
)

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

	for i := range arch {
		errDownload := filter.DownloadPackages(
			"packages/"+arch[i]+"-packages",
			url+"/main/binary-"+arch[i]+"/Packages")

		if errDownload != nil {
			log.Fatal(errDownload)
		}
	}

	log.Println("[!] Filtering...")
	f.Parser()

	errRmdir := os.RemoveAll("packages")
	if errRmdir != nil {
		log.Fatal(errRmdir)
	}
	log.Println("Deleted all Packages files.")

	log.Printf("[!] Starting HTTP server to serve json files at port: %s\n", port)
	handleFunctions()
	errHttp := http.ListenAndServe(":"+port, nil)
	if errHttp != nil {
		log.Fatal(errHttp)
	}
}
