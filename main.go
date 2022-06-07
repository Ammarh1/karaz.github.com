package main

import (
	controllers "karaz/src/Controllers"
	"karaz/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.SQLconnect()

	Router := gin.Default()
	Router.GET("/", controllers.GetUsers)
	Router.POST("/post", controllers.PostUser)
	Router.PUT("/update", controllers.UpdateUser)
	Router.DELETE("/delete", controllers.DeletUser)

	Router.Run()

}
