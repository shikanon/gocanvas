package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("./src", "./src")
	r.Static("./assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		r.LoadHTMLGlob("./*.html")
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
