package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	if err := router.SetTrustedProxies(nil); err != nil {
		panic(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	return router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := newRouter().Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
