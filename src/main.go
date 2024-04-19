package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes defined in the routes package
	routes := r.Group("/fizzybuzzy")
	{
		routes.GET("/", replace)
		routes.POST("/", replace)
		routes.GET("/stats", stats)
	}

	if err := r.Run(":8080"); err != nil {
		return
	}

}

func replace(c *gin.Context) {
	var input Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	message, err := buildMessage(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": message,
	})
}

var ErrorOnDivider = errors.New("Int1 and Int2 should be positive non zero integers")

func buildMessage(input Input) (string, error) {
	if input.Int1 == 0 || input.Int2 == 0 {
		return "", ErrorOnDivider
	}
	message := ""
	for i := 1; i <= input.Limit; i++ {
		toAdd := ""
		if i%input.Int1 == 0 {
			toAdd = input.Str1
		}
		if i%input.Int2 == 0 {
			toAdd = toAdd + input.Str2
		}
		if toAdd != "" {
			message = message + toAdd + ","
		} else {
			message = message + strconv.Itoa(i) + ","
		}
	}
	// Remove last comma
	if message != "" {
		message = message[:len(message)-1]
	}
	return message, nil
}

func stats(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API!",
	})
}

type Input struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}
