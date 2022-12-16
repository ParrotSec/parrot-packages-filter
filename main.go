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

func handleFunctions() {
	http.HandleFunc("/packages/main/amd64/", getAMD64Packages)
	http.HandleFunc("/packages/main/arm64/", getARM64Packages)
}

func main() {
	const port = 8080

	f := new(filter.Package)

	log.Println("Downloading packages...")
	amd64, arm64 := "amd64", "arm64"
	filter.DownloadPackages("./"+amd64+"-packages", "https://download.parrot.sh/parrot/dists/parrot/main/binary-"+amd64+"/Packages")
	filter.DownloadPackages("./"+arm64+"-packages", "https://download.parrot.sh/parrot/dists/parrot/main/binary-"+arm64+"/Packages")

	log.Println("Filtering...")
	f.Parser()
	log.Printf("[!] Starting HTTP server at port: %d\n", port)

	handleFunctions()
	errHttp := http.ListenAndServe(":8080", nil)
	if errHttp != nil {
		log.Fatal(errHttp)
	}
}
