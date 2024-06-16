package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksks2012/ubiquitous-sniffle/internal/model"
)

type ShortURL struct{}

func NewShortUrl() ShortURL {
	return ShortURL{}
}

func (a ShortURL) Get(c *gin.Context) {
	shortURL := c.Param("url")

	url, err := model.GetURL(shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}
