package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Number Game Session Manager is running")
	})

	r.GET("/session/find", func(c *gin.Context) {
		data := ResponseData{
			Message: "Session found",
			//URL:     "ws://192.168.0.6:8080",
			URL: "ws://34.168.132.186:8080",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
