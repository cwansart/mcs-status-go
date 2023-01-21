package main

import (
	"net/http"

	"de.cwansart.mcss/settings"
	"de.cwansart.mcss/status"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		url := settings.Get(settings.ServerUrlKey)
		s := status.Get(url)
		c.JSON(http.StatusOK, s)
	})
	r.Run()
}
