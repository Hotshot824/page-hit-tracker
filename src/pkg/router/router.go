package router

import (
	"fmt"
	"text/template"
	"page-hit-tracker/pkg/db"
	"page-hit-tracker/pkg/config"
	"github.com/gin-gonic/gin"
)

var cfg = config.NewConfig()

func Run() error {
	// Initialize
	cfg.Initlog()
	fmt.Println("TrackerURL: ", cfg.GetURL() + "/tracker.js")

	gin.SetMode(gin.ReleaseMode)

	// Run router
	router := gin.Default()
	setupRoutes(router)
	err := router.Run(":8080")

	return err
}

func setupRoutes(router *gin.Engine) {
	router.SetTrustedProxies(nil)
	router.Use(handleCors)
	router.GET("/", serveCounter)
	router.GET("/tracker.js", serveTracker)
}

func handleCors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type")
}

func serveCounter(c *gin.Context) {
	href := c.Query("href")
	if href == "" {
		return
	} else {
		url := murmur3Hasher(href)
		handler, err := db.NewHandler()
		if err != nil {
			panic(err)
		}

		count, err := handler.Count(url)
		if err != nil {
			panic(err)
		}

		defer handler.Close()

		c.JSON(200, gin.H{"href": murmur3Hasher(href), "count": count})
	}
}

func serveTracker(c *gin.Context) {
	c.Header("Cache-Control", "max-age=0, no-cache, no-store, must-revalidate")

	template, err := template.ParseFiles("./assets/tracker.js")
	if err != nil {
		panic(err)
	}

	mapping := map[string]string{
		"PUBLISHING_URL": cfg.GetURL(),
	}
	err = template.Execute(c.Writer, mapping)
	if err != nil {
		panic(err)
	}
}