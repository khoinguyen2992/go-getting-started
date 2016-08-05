package routers

import "github.com/gin-gonic/gin"

var router *gin.Engine

func init() {
	if router == nil {
		router = gin.New()
		router.Use(gin.Logger())
	}
}

func addHandlers() {
	addJobHandlers(router)
}

func Handle(url string) {
	addHandlers()
	router.Run(url)
}
