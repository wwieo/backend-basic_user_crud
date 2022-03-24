package view

import (
	controller "backend-user_crud/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	userController := controller.NewUserController()
	apiRecordController := controller.NewApiRecordController()

	apiRecordController.InitApiRecord()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/plain", []byte("Hello world!"))
	})

	router.POST("/user", userController.CreateUser)

	router.GET("/user/:username", userController.FindUser)

	router.PUT("/user/:username", userController.UpdateUser)

	router.DELETE("/user/:username", userController.DeleteUser)

	router.GET("/api_records", apiRecordController.GetApiRecord)

	return router
}
