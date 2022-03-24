package controller

import (
	dao "backend-user_crud/dao"

	"github.com/gin-gonic/gin"
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

func (apiRecordController ApiRecordController) GetApiRecord(context *gin.Context) {
	if records := apiRecordDao.GetApiRecord(); records != nil {
		utilsController.ReturnResult(context, true, records)
	} else {
		utilsController.ReturnResult(context, false, "Get api records failed")
	}
}
