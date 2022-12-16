package src

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

const (
	prefixName       = "Package: "
	prefixDesc       = "Description: "
	prefixVersion    = "Version: "
	prefixMaintainer = "Maintainer: "
	prefixArch       = "Architecture: "
)

type Package struct {
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Version      string `json:"Version"`
	Maintainer   string `json:"Maintainer"`
	Architecture string `json:"Architecture"`
}

type PackageSet struct {
	Packages map[string]Package
}

func (p *Package) Parser() {
	version := map[string]string{
		"amd64": "./amd64-packages",
		"arm64": "./arm64-packages",
		"armhf": "./armhf-packages",
		"i386":  "./i386-packages",
	}

	for i := range version {
		file, _ := os.Open(version[i])

		scanner := bufio.NewScanner(file)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)

		lineNumber := 0

		var P PackageSet
		P.Packages = make(map[string]Package)

		for scanner.Scan() {
			line := scanner.Text()

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
			}

			P.Packages[p.Name] = Package{
				Name:         p.Name,
				Description:  p.Description,
				Version:      p.Version,
				Maintainer:   p.Maintainer,
				Architecture: p.Architecture,
			}

			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("Error on line %v: %v", lineNumber, err)
		}
		data, _ := json.MarshalIndent(P, "", "\t")
		errWriteFile := os.WriteFile(version[i]+".json", data, 0644)
		if errWriteFile != nil {
			log.Fatalf("Can't create and write to %s", errWriteFile)
		}
	}
}
