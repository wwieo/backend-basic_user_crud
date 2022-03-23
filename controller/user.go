package controller

import (
	"github.com/gin-gonic/gin"

	dao "backend-user_crud/dao"
	model "backend-user_crud/model"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

var (
	userDao         = dao.NewUserDao()
	utilsController = NewUtilsController()
)

func (userController UserController) CreateUser(c *gin.Context) {

	user := &model.User{
		Username: c.Request.FormValue("username"),
		Password: c.Request.FormValue("password"),
	}

	if err := c.Bind(&user); err != nil {
		return
	}
	if _, err := userDao.Insert(user); err != nil {
		utilsController.ReturnResult(c, false, "Add failed")
	} else {
		utilsController.ReturnResult(c, true, "Add successfully")
	}
}

func (userController UserController) UpdateUser() {

}

func (userController UserController) FindUser() {

}

func (userController UserController) DeleteUser() {

}
