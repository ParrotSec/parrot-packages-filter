package main

import (
	"log"
	"net/http"
	filter "package-filter/src"
)

func getAMD64Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "amd64-packages.json")
}

func getARM64Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "arm64-packages.json")
}

func getARMHFPackages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "armhf-packages.json")
}

func geti386Packages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "i386-packages.json")
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

	log.Println("Downloading packages...")
	arch := []string{
		"amd64",
		"arm64",
		"armhf",
		"i386",
	}

	for i := range arch {
		errDownload := filter.DownloadPackages(
			"./"+arch[i]+"-packages",
			url+"/main/binary-"+arch[i]+"/Packages")

		if errDownload != nil {
			log.Fatal(errDownload)
		}
	}

	log.Println("Filtering...")
	f.Parser()
	log.Printf("[!] Starting HTTP server at port: %s\n", port)

	handleFunctions()
	errHttp := http.ListenAndServe(":"+port, nil)
	if errHttp != nil {
		log.Fatal(errHttp)
	}
}
