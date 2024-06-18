package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksks2012/ubiquitous-sniffle/internal/model"
	"github.com/ksks2012/ubiquitous-sniffle/pkg/hash"
)

type Shorten struct{}

func NewShorten() Shorten {
	return Shorten{}
}

// Create handles the creation of a shortened URL.
// It retrieves the URL and method from the request's form data,
// hashes the URL using the specified method, and returns the hashed URL.
// If the method is invalid, it returns without further processing.
func (a Shorten) Create(c *gin.Context) {
	url := c.DefaultPostForm("url", "missing url in input")
	method := c.DefaultPostForm("method", "missing method in input")

	var hashMethod func(string) string
	if method == "md5" {
		hashMethod = hash.HashMD5
	} else if method == "sha1" {
		hashMethod = hash.HashSHA1
	} else if method == "crc32" {
		hashMethod = hash.HashCRC32
	} else {
		// handle invalid method here
		return
	}

	hashedURL := hashMethod(url)
	// continue with the rest of the code

	// Save URL in redis cache
	err := model.SetKey(hashedURL, url)
	if err != nil {
		// handle error here
		return
	}

	// Save the URL and shortened URL in the database
	err = model.SaveURL(url, hashedURL)
	if err != nil {
		// handle error here
		return
	}

	c.JSON(http.StatusOK, gin.H{"hashedURL": hashedURL})
}
