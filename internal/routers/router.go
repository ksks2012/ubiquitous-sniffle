package routers

import (
	v1 "github.com/ksks2012/ubiquitous-sniffle/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	var shorten = v1.NewShorten()
	var shortUrl = v1.NewShortUrl()

	apiv1 := r.Group("/api/v1")
	// long url -> short url
	apiv1.POST("/shorten", shorten.Create)
	apiv1.GET("/shortUrl", shortUrl.Get)

	return r
}
