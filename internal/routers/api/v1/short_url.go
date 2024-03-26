package v1

import "github.com/gin-gonic/gin"

type ShortURL struct{}

func NewShortUrl() ShortURL {
	return ShortURL{}
}

func (a ShortURL) Get(c *gin.Context) {
	return
}
