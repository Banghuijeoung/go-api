package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
	"time"
)

func Route(){

	router := gin.Default()
	//cors도메인을 방지하는것 : 모든 브라우저의 정책, 보안의 이슈떄문에 필요
	router.Use(cors.New(cors.Config{
		//Origins? * 모든 요청을 허용하겠다.
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		//AllowCredentials 비밀스러운 정보까지 허용하겠다
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	//router.Group이란? 기본적인 도메인 그룹
	test := router.Group("/api/v1/test")
	{
		test.GET("", controllers.Tests.Test)
	}

	member := router.Group("/api/v1/new")
	{
		member.GET("", controllers.News.News)
	}
	router.Run()
}
