package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	//	"github.com/k0kubun/pp"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wmw9/go-makaba"
	yarb "github.com/wmw9/yarb-struct"
	"net/http"
)

func handlePing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func handlePost(c *gin.Context) {
	// Unmarshal payload into struct
	var p yarb.Payload
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println("Couldn't bind json")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, num, err := sendToMakaba(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"subject": subject, "num": num, "board": p.Board})

}

func sendToMakaba(p yarb.Payload) (string, string, error) {
	// Find thread number
	num, subject, err := makaba.Get(p.Board).Thread(p.Thread)
	if err != nil {
		return "", "", err
	}
	log.Debugf("Found '%v' thread #%v", subject, num)

	// Post in thread we found earlier
	num, err = makaba.Post().Board(p.Board).Thread(num).Name(p.Person).Comment(p.Caption).File(p.Files).Do(Passcode)
	if err != nil {
		return "", "", err
	}
	return subject, num, nil
}

func UpdateIGStoriesTs(p yarb.Payload) error {
	url := fmt.Sprintf("http://%v/yarb/user/name/%v/date/instagram_stories/%v", YarbDBApiURL, p.Person, p.Timestamp)
	log.Debugf("%v\n", url)
	client := resty.New()
	client.SetBasicAuth(yarbBasicAuthUser, yarbBasicAuthPass)
	resp, err := client.R().Get(url)
	if err != nil {
		return err
	}
	log.Infof("%v:\n%v", resp.String())
	return nil
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
