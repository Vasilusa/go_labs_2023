package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl",
			gin.H{
				"title": "Main website",
			})
	})

	r.GET("/birthday", func(c *gin.Context) {
		c.HTML(http.StatusOK, "birth.tmpl", gin.H{
			"message":   "Wishlist",
			"wish_list": []string{"толстовка", "телефон", "торт"},
		})
	})

	r.Static("/image", "./resources")

	r.Run() //localhost:8080

	log.Println("Server down")
}
