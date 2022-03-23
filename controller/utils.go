package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UtilsController struct {
}

func NewUtilsController() *UtilsController {
	return &UtilsController{}
}

func (utilsController UtilsController) ReturnResult(c *gin.Context, success bool, message string) {
	c.JSON(http.StatusOK, gin.H{
		"success": success,
		"message": message,
	})
}
