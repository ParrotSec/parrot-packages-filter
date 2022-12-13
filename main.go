package main

import (
	"fmt"
	"os"
	filter "package-filter/src"
)

func main() {
	if _, err := os.Stat("./packages"); err == nil {
		fmt.Println("Parsing packages...")
		p := new(filter.Package)
		p.Parser()
	} else {
		fmt.Println("Downloading packages...")
		filter.GetPackages("./packages", "https://download.parrot.sh/parrot/dists/parrot/main/binary-amd64/Packages")
		fmt.Println("Done! Parsing...")
		//defer f.Parser()
	}
}
