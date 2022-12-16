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

	architecture := map[string]string{
		"amd64": "packages/amd64-packages",
		"arm64": "packages/arm64-packages",
		"armhf": "packages/armhf-packages",
		"i386":  "packages/i386-packages",
	}

	/*
		errJsonDir := os.Chdir("../json")
		if errJsonDir != nil {
			log.Fatal(errJsonDir)
		}
		log.Println("Dir changed successfully")
		newDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current Working Direcotry: %s\n", newDir)
	*/

	for i := range architecture {
		file, _ := os.Open(architecture[i])

		scanner := bufio.NewScanner(file)
		buf := make([]byte, 0, 64*1024)
		scanner.Buffer(buf, 1024*1024)

		var P PackageSet
		P.Packages = make(map[string]Package)

		lineNumber := 0

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

		errScanner := scanner.Err()
		if errScanner != nil {
			log.Fatalf("Error on line %v: %v", lineNumber, errScanner)
		}

		data, _ := json.MarshalIndent(P, "", "\t")

		s := strings.TrimPrefix(architecture[i], "packages/")
		jsonData := s + ".json"

		errWriteFile := os.WriteFile(jsonData, data, 0644)
		if errWriteFile != nil {
			log.Fatalf("Can't write and %s", errWriteFile)
		}

		errJsonData := os.Rename(jsonData, "./json/"+jsonData)
		if errJsonData != nil {
			log.Fatal(errJsonData)
		}
	}
}
