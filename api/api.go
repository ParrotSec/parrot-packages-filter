package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	p "package-filter/src"
)

func getPackages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, new(p.Package))
}

func PackagesAPI() {
	const port = "8080"

	router := gin.Default()

	branch := []string{
		"contrib",
		"main",
		"non-free",
	}

	arch := []string{
		"amd64",
		"arm64",
		"armhf",
		"i386",
	}

	for b := range branch {
		for a := range arch {
			router.GET("/packages/"+branch[b]+"/"+arch[a], getPackages)
		}
	}
	// router.GET("/packages/:id", getAlbumByID)

	errHttpServer := router.Run("localhost:" + port)
	if errHttpServer != nil {
		log.Fatal(errHttpServer)
	}
}
