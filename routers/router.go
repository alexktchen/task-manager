package routers

import (
	v1 "github.com/alexktchen/task-manager/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/task", v1.AddTask)
		apiv1.GET("/tasks", v1.GetTasks)
		apiv1.PUT("/task/:id", v1.UpdateTask)
		apiv1.DELETE("/task/:id", v1.DeleteTask)

	}

	return r
}
