package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func stateEndpoint(c *gin.Context, cfg *Config) {
	c.JSON(200, cfg)
}

func goEndpoint(c *gin.Context, cfg *Config) {
	linkname := c.Param("linkname")
	link, ok := cfg.Links[linkname]
	if !ok {
		c.JSON(404, gin.H{"error": fmt.Sprintf("Link %s not found", linkname)})
	} else {
		c.Redirect(http.StatusFound, link.Url)
	}
}

func serveCommand(flags commonFlags) {
	r := gin.Default()

	endpoint := func(cb func(c *gin.Context, cfg *Config)) func(c *gin.Context) {
		return func(c *gin.Context) {
			cfg, err := loadConfig(flags.ConfigPath)
			if err != nil {
				reportError(c, err)
				return
			}
			cb(c, cfg)
		}
	}

	r.GET("/state", endpoint(stateEndpoint))
	r.GET("/go/:linkname", endpoint(goEndpoint))

	r.Run()
}
