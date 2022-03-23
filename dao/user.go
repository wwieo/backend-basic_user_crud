package dao

import (
	model "backend-user_crud/model"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (userDao *UserDao) Find(user *model.User) (rstUser *model.User, err error) {
	rstUser = &model.User{}
	rst := Orm.Where("username = ?", user.Username).First(rstUser)

	if rst.Error != nil {
		err = rst.Error
		return
	}
	return
}

func (userDao *UserDao) Insert(user *model.User) (username string, err error) {
	rst := Orm.Create(user)
	username = user.Username

	if rst.Error != nil {
		err = rst.Error
		return
	}
	return
}

func (userDao *UserDao) Update(user *model.User) (username string, err error) {
	rst := Orm.Model(&user).Updates(model.User{
		Password: user.Password,
	})
	username = user.Username

	if rst.Error != nil {
		err = rst.Error
		return
	}
	return
}

func (userDao *UserDao) Delete(user *model.User) (username string, err error) {
	rst := Orm.Where("username = ?", user.Username).Delete(user)
	username = user.Username

	if rst.Error != nil {
		err = rst.Error
		return
	}
	return
}
