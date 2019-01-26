package controllers

import (
	"github.com/gin-gonic/gin"
	)

type NewController struct {

}

var News NewController

func (self NewController) News(c *gin.Context){

	c.JSON(200, gin.H{
		"message": "new",
	})
}