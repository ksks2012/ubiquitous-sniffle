package v1

import "github.com/gin-gonic/gin"

type Shorten struct{}

func NewShorten() Shorten {
	return Shorten{}
}

func (a Shorten) Create(c *gin.Context) {
	return
}
