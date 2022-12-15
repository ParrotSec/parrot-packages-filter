package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	filter "package-filter/src"
)

func getPackages(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "packages.json")
}

func main() {
	p := new(filter.Package)

	if _, err := os.Stat("./packages"); err == nil {
		fmt.Println("Parsing packages...")
		p.Parser()
		fmt.Println("Done. Check packages.json")
	} else {
		fmt.Println("Downloading packages...")
		err := filter.GetPackages("./packages", "https://download.parrot.sh/parrot/dists/parrot/main/binary-amd64/Packages")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done! Filtering...")
		defer p.Parser()
	}

	http.HandleFunc("/packages/", getPackages)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
