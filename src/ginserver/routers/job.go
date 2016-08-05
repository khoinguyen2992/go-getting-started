package routers

import (
	"ginserver/handlers"
	"ginserver/handlers/job"

	"github.com/gin-gonic/gin"
)

func addJobHandlers(router *gin.Engine) {
	jobHandlers := router.Group("/jobs")
	jobHandlers.GET("/:id", handlers.NewHandler(&job.GetHandler{}))
}
