package controllers

import (
	"fmt"

	"github.com/alochym01/ftp-api/api/models"
	"github.com/gin-gonic/gin"
)

type alochym struct {
	Ping string
}

func GetPing(c *gin.Context) {
	c.String(200, "pong")
}

func PostPingJson(c *gin.Context) {
	var tmp alochym
	err := c.BindJSON(&tmp)
	if err != nil {
		fmt.Println("cannot get json object")
	}
	c.JSON(200, gin.H{
		"ping": models.Ping(tmp.Ping),
	})
}

func PostPingForm(c *gin.Context) {
	var tmp alochym
	err := c.Bind(&tmp)
	if err != nil {
		fmt.Println("cannot get json object")
	}
	fmt.Println(tmp.Ping)
	c.String(200, models.Ping(tmp.Ping))
}
