package controllers

import (
	"github.com/gin-gonic/gin"
	"rest-api/service"
)

type TestController struct {


}

var Tests TestController

func (self TestController) Test(c *gin.Context){
	c.JSON(200, gin.H{

		"message":  service.Service1.Test(),
	})
}