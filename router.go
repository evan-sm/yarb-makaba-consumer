package main

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/wmw9/go-makaba"
	"log"
	"net/http"
)

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func handlePost(c *gin.Context) {
	// Unmarshal payload into struct
	var p Payload
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println("Couldn't bind json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pp.Println(p)

	// Find thread number
	num, subject, err := makaba.Get(p.Board).Thread(p.Thread)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Post in thread we found earlier
	num, err = makaba.Post().Board(p.Board).Thread(num).Subject(p.Person).Comment(p.Caption).File(p.Files).Do(Passcode)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subject": subject, "num": num, "board": p.Board})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	authorized := r.Group("/makaba", gin.BasicAuth(gin.Accounts{
		yarbBasicAuthUser: yarbBasicAuthPass,
	}))

	authorized.GET("/ping", handlePing)
	authorized.POST("/post", handlePost)

	return r
}
