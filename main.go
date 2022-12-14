package main

import (
	"fmt"
	"os"
	filter "package-filter/src"
)

func main() {
	p := new(filter.Package)

	if _, err := os.Stat("./packages"); err == nil {
		fmt.Println("Parsing packages...")
		p.Parser()
	} else {
		fmt.Println("Downloading packages...")
		filter.GetPackages("./packages", "https://download.parrot.sh/parrot/dists/parrot/main/binary-amd64/Packages")
		fmt.Println("Done! Filtering...")
		defer p.Parser()
	}
}
