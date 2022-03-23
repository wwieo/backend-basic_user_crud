package controller

import (
	"fmt"

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

func (userController UserController) FindUser(context *gin.Context) {
	user := &model.User{
		Username: context.Param("username"),
	}

	if rstUser, err := userDao.Find(user); err != nil {
		utilsController.ReturnResult(context, false, fmt.Sprintf("%s has not found", user.Username))
	} else {
		rstUser.Password = ""
		utilsController.ReturnResult(context, true, rstUser)
	}
}

func (userController UserController) CreateUser(context *gin.Context) {
	user := &model.User{
		Username: context.Request.FormValue("username"),
		Password: context.Request.FormValue("password"),
	}

	if err := context.Bind(&user); err != nil {
		return
	}
	if _, err := userDao.Insert(user); err != nil {
		utilsController.ReturnResult(context, false, "Add failed")
	} else {
		user.Password = ""
		utilsController.ReturnResult(context, true, user)
	}
}

func (userController UserController) UpdateUser(context *gin.Context) {
	user := &model.User{
		Username: context.Param("username"),
		Password: context.Request.FormValue("password"),
	}

	if err := context.Bind(&user); err != nil {
		return
	}

	if _, err := userDao.Find(user); err != nil {
		utilsController.ReturnResult(context, false, fmt.Sprintf("%s has not found", user.Username))
		return
	}

	if _, err := userDao.Update(user); err != nil {
		utilsController.ReturnResult(context, false, "Update failed")
	} else {
		user.Password = ""
		utilsController.ReturnResult(context, true, user)
	}
}

func (userController UserController) DeleteUser(context *gin.Context) {
	user := &model.User{
		Username: context.Param("username"),
	}

	_, err := userDao.Find(user)
	if err != nil {
		utilsController.ReturnResult(context, false, fmt.Sprintf("%s has not found", user.Username))
		return
	}

	if username, err := userDao.Delete(user); err != nil {
		utilsController.ReturnResult(context, false, "Delete failed")
	} else {
		utilsController.ReturnResult(context, true, username)
	}
}
