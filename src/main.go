package main

import (
	"errors"
	"fmt"
	"lbc/fizzy/src/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var client *gorm.DB

func main() {
	var err error
	client, err = db.ConnectToMariaDB()
	if err != nil {
		log.Fatal(err)
	}

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

	err = db.Incr(client, input.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
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

	val, err := db.GetMaxQuery(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(200, gin.H{
		val.Id: val.Queries,
	})

}

type Input struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func (i *Input) String() string {
	s := fmt.Sprintf("Int1: %d, Int2: %d, Limit: %d, Str1: %s, Str2 %s", i.Int1, i.Int2, i.Limit, i.Str1, i.Str2)
	return s
}
