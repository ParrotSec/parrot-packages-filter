// filter.go takes care of filtering the desired informations
// from Packages for each architecture.

package src

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

// Prefixes and structs to manage the components of each single Package.
const (
	prefixName       = "Package: "
	prefixDesc       = "Description: "
	prefixVersion    = "Version: "
	prefixMaintainer = "Maintainer: "
	prefixArch       = "Architecture: "
	prefixSection    = "Section: "
)

type Package struct {
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Version      string `json:"Version"`
	Maintainer   string `json:"Maintainer"`
	Architecture string `json:"Architecture"`
	Section      string `json:"Section"`
}

type PackageSet struct {
	Packages map[string]Package
}

func (p *Package) Parser() {

	branch := []string{
		"contrib",
		"main",
		"non-free",
	}

	// Each Packages is contained within a temporary directory called packages.
	for b := range branch {

		architecture := map[string]string{
			"amd64": "packages/" + branch[b] + "/amd64",
			"arm64": "packages/" + branch[b] + "/arm64",
			"armhf": "packages/" + branch[b] + "/armhf",
			"i386":  "packages/" + branch[b] + "/i386",
		}

		// For each architecture the filter phase takes place here.
		for a := range architecture {
			file, _ := os.Open(architecture[a])

			// Increase of the buffer because the size of each single Packages file is large.
			scanner := bufio.NewScanner(file)
			buf := make([]byte, 0, 64*1024)
			scanner.Buffer(buf, 1024*1024)

			var P PackageSet
			P.Packages = make(map[string]Package)

			lineNumber := 0

			// Scan every line within every Packages for every architecture.
			for scanner.Scan() {
				line := scanner.Text()

				// Each line is scanned and filtered according to prefixes.
				if strings.HasPrefix(line, prefixName) {
					name := strings.TrimPrefix(line, prefixName)
					p.Name = name
				} else if strings.HasPrefix(line, prefixDesc) {
					desc := strings.TrimPrefix(line, prefixDesc)
					p.Description = desc
				} else if strings.HasPrefix(line, prefixVersion) {
					version := strings.TrimPrefix(line, prefixVersion)
					p.Version = version
				} else if strings.HasPrefix(line, prefixMaintainer) {
					maintainer := strings.TrimPrefix(line, prefixMaintainer)
					p.Maintainer = maintainer
				} else if strings.HasPrefix(line, prefixArch) {
					arch := strings.TrimPrefix(line, prefixArch)
					p.Architecture = arch
				} else if strings.HasPrefix(line, prefixSection) {
					section := strings.TrimPrefix(line, prefixSection)
					p.Section = section
				}

				// Each filtered line is stored in the Package struct above.
				P.Packages[p.Name] = Package{
					Name:         p.Name,
					Description:  p.Description,
					Version:      p.Version,
					Maintainer:   p.Maintainer,
					Architecture: p.Architecture,
					Section:      p.Section,
				}
				lineNumber++
			}

			errScanner := scanner.Err()
			if errScanner != nil {
				log.Fatalf("Error on line %v: %v", lineNumber, errScanner)
			}

			// Once the filtering stage is complete, the data is indented in a JSON file.
			data, _ := json.MarshalIndent(P, "", "\t")

			// For simplicity, the word "packages" has been removed from the architecture map
			// in order to better manage the movement of new JSON files within the program.
			// Check architecture variable.
			s := strings.TrimPrefix(architecture[a], "packages/"+branch[b]+"/")

			// The filtered and indented JSON file is correctly written in its format.
			jsonData := s + ".json"
			errWriteFile := os.WriteFile(jsonData, data, os.ModePerm)
			if errWriteFile != nil {
				log.Fatalf("Can't %s", errWriteFile)
			}

			// Each JSON file is now placed in a specific directory, "json".
			errJsonData := os.Rename(jsonData, "json/"+architecture[a]+"/"+jsonData)
			if errJsonData != nil {
				log.Fatal(errJsonData)
			}
			defer file.Close()
		}
	}
}
