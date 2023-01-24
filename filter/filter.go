// filter.go takes care of filtering the desired information from Packages for each architecture.

package filter

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	c "package-filter/constants"
	"strings"
)

// Structs to manage the components of each single Package.
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

	// Each Packages file is contained within a temporary directory called packages.
	for b := range c.Branch {

		architecture := map[string]string{
			"amd64": "packages/" + c.Branch[b] + "/amd64",
			"arm64": "packages/" + c.Branch[b] + "/arm64",
			"armhf": "packages/" + c.Branch[b] + "/armhf",
			"i386":  "packages/" + c.Branch[b] + "/i386",
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

			// Scan every line within every Packages file for every architecture.
			for scanner.Scan() {
				line := scanner.Text()

				// Each line is scanned and filtered according to prefixes.
				switch {

				case strings.HasPrefix(line, c.PrefixName):
					name := strings.TrimPrefix(line, c.PrefixName)
					p.Name = name

				case strings.HasPrefix(line, c.PrefixDesc):
					desc := strings.TrimPrefix(line, c.PrefixDesc)
					p.Description = desc

				case strings.HasPrefix(line, c.PrefixVersion):
					version := strings.TrimPrefix(line, c.PrefixVersion)
					p.Version = version

				case strings.HasPrefix(line, c.PrefixMaintainer):
					maintainer := strings.TrimPrefix(line, c.PrefixMaintainer)
					p.Maintainer = maintainer

				case strings.HasPrefix(line, c.PrefixArch):
					arch := strings.TrimPrefix(line, c.PrefixArch)
					p.Architecture = arch

				case strings.HasPrefix(line, c.PrefixSection):
					section := strings.TrimPrefix(line, c.PrefixSection)
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
			s := strings.TrimPrefix(architecture[a], "packages/"+c.Branch[b]+"/")

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
		}
	}
}
