/*
	Created by Bontus Mayor
	Contact <bontus.doku@gmail.com>
*/

package main

import (
	"Doku/config"
	"Doku/internal/app"
	"Doku/internal/routes"
	"Doku/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	APPNAME    = "Doku"
	APPVERSION = "1.0.0"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/index.html")
	r.Static("/dist", "./templates/dist")

	// Load manifest file
	manifest, err := utils.LoadManifest("templates/dist/.vite/manifest.json")
	if err != nil {
		log.Fatalf("could not load manifest file: %v", err)
	}

	r.GET("/", func(ctx *gin.Context) {
		entry, ok := manifest["index.html"].(map[string]interface{})
		if !ok {
			log.Fatalf("could not find entry for index.html in the manifest")
		}

		script, ok := entry["file"].(string)
		if !ok {
			log.Fatalf("could not find the script file in the manifest entry")
		}

		css, ok := entry["css"].([]interface{})
		var cssFile string
		if ok && len(css) > 0 {
			cssFile = css[0].(string)
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"script": script,
			"css":    cssFile,
		})
	})

	// Load env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf(".env file is missing: %v", err)
	}


	// Load Configurations
	cfg := config.LoadConig()

	if cfg.MongoUri == "" {
		log.Fatal("MONGO_URI is not set")
	}

	// Initialize the application
	doku := app.NewApp(APPNAME, APPVERSION, cfg.MongoUri)

	// Initialize routes
	routes.InitializeRoutes(doku)

	// Start server
	if err := doku.Router.Run(":9000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
