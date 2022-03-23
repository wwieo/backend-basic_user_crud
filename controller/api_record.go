package controller

import (
	dao "backend-user_crud/dao"
)

var (
	apiRecordDao = dao.NewApiRecordDao()
)

type ApiRecordController struct {
}

func NewApiRecordController() *ApiRecordController {
	return &ApiRecordController{}
}

func (apiRecordController ApiRecordController) InitApiRecord() {
	apiRecordDao.InitApiRecord()
}

func (apiRecordController ApiRecordController) IncrApiRecord(key string) {
	apiRecordDao.IncrApiRecord(key)
}
