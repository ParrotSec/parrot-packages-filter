package fileops

import (
	"log"
	"os"
)

func Mkdir() {
	err := os.Mkdir("packages", os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func Rmdirs() {
	err := os.RemoveAll("packages")
	if err != nil {
		log.Fatal(err)
	}
}
