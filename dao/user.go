package dao

import (
	model "backend-user_crud/model"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (userDao *UserDao) Insert(user *model.User) (id int64, err error) {
	rst := Orm.Create(user)
	id = user.ID
	if rst.Error != nil {
		err = rst.Error
		return
	}
	return
}
