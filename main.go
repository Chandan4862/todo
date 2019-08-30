package main

import (
	"todo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	task := handler.New()
	route.GET("/get", handler.GetTask(task))
	route.POST("/add", handler.PostTask(task))
	route.DELETE("/delete/:name", handler.Delete(task))
	route.PUT("/update/:name",handler.Update(task) )
	route.Run() // listen and serve on 0.0.0.0:8080
}
