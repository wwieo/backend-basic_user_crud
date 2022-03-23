package view

import (
	controller "backend-user_crud/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	s := controller.NewUserController()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Hello world!"))
	})

	router.POST("/user", s.CreateUser)

	return router
}
