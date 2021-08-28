package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func reportError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"error": err.Error(),
	})
}

func stateEndpoint(c *gin.Context, cfg *Config) {
	c.JSON(200, cfg)
}

func evaluateQuery(cfg *Config, query string) (string, bool, error) {
	link, ok := cfg.Links[query]
	if !ok {
		return "", false, nil
	} else {
		tmpl, err := template.New(link.Name).Parse(link.Url)
		if err != nil {
			return "", false, err
		}
		var out bytes.Buffer
		if err = tmpl.Execute(&out, cfg.Variables); err != nil {
			return "", false, err
		}
		result := out.String()
		return result, true, nil
	}
}

// handleEvaluate pulls out the common parts of go/ and eval/
func handleEvaluation(c *gin.Context, cfg *Config) (string, bool) {
	query := c.Param("query")
	result, found, err := evaluateQuery(cfg, query)
	if err != nil {
		reportError(c, err)
	} else if found == false {
		c.JSON(404, gin.H{"error": fmt.Sprintf("query %s could not be resolved: %s", query, err)})
	} else {
		fmt.Printf("eval: %s => %s\n", query, result)
		return result, true
	}
	return "", false
}

func evalEndpoint(c *gin.Context, cfg *Config) {
	result, cont := handleEvaluation(c, cfg)
	if cont {
		c.JSON(200, result)
	}
}

func goEndpoint(c *gin.Context, cfg *Config) {
	result, cont := handleEvaluation(c, cfg)
	if cont {
		c.Redirect(http.StatusFound, result)
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
	r.GET("/eval/:query", endpoint(evalEndpoint))
	r.GET("/go/:query", endpoint(goEndpoint))

	r.Run()
}
