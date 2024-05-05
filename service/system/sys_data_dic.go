package system

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	comReq "gin-one/model/common/request"
	"gin-one/model/system"
	"gin-one/service/common"
	"gorm.io/gorm"
)

type DataDicService struct{}

func (service *DataDicService) CreateDataDic(dataDic *system.SysDataDic) (int, error) {
	ok, err := common.IsModelDataExist(&system.SysDataDic{Code: dataDic.Code})
	if err != nil {
		return message.OPER_ERR, err
	} else if ok {
		return message.DATA_EXIST, nil
	}

	err = global.DB.Create(dataDic).Error
	if err != nil {
		return message.OPER_ERR, err
	}
	return message.OPER_OK, nil
}

func (service *DataDicService) UpdateDataDic(dataDic *system.SysDataDic) error {
	return global.DB.Updates(dataDic).Error
}

func (service *DataDicService) DeleteDataDicByID(id string) (msgCode int, err error) {
	return common.DeleteModelByID(&system.SysDataDic{}, id)
}

func (service *DataDicService) DeleteDataDicList(ids []string) (msgCode int, err error) {
	return common.DeleteModelList(&system.SysDataDic{}, ids)
}

func (service *DataDicService) GetDataDicByID(id string) (codeMsg int, dataDic system.SysDataDic, err error) {
	codeMsg, err = common.GetModelByID(&dataDic, id)
	return
}

func (service *DataDicService) GetDataDicByIds(ids []string) (codeMsg int, dataDics []system.SysDataDic, err error) {
	codeMsg, err = common.GetModelByIds(&system.SysDataDic{}, ids, &dataDics)
	return
}

func (service *DataDicService) GetDataDicByParentIds(parentIds []string) (codeMsg int, dataDics []system.SysDataDic, err error) {
	codeMsg, err = common.GetModelByParentIds(&system.SysDataDic{}, parentIds, &dataDics)
	return
}

func (service *DataDicService) GetDataDicListByCode(code string) (codeMsg int, dataDics []system.SysDataDic, err error) {
	var model system.SysDataDic
	err = global.DB.Select("id").Where("code = ?", code).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return message.QUERY_OK, dataDics, nil
		}
		return message.QUERY_ERR, dataDics, err
	}

	err = global.DB.Where("parent_id = ?", model.ID).Find(&dataDics).Error
	if err != nil {
		return message.QUERY_ERR, dataDics, err
	}

	return message.QUERY_OK, dataDics, nil
}

func (service *DataDicService) GetDataDic(dataDic *system.SysDataDic) (int, error) {
	return common.GetModel(dataDic)
}

func (service *DataDicService) GetDataDicPagination(pageInfo comReq.PageInfo, dataDic *system.SysDataDic, parentIds []string) (_ int, dataDics []system.SysDataDic, count int64, err error) {
	columns := []string{"name", "code", "description"}
	vals := []string{dataDic.Name, dataDic.Code, dataDic.Description}
	db := global.DB.Model(&system.SysDataDic{})

	if dataDic.Status != "" {
		db = db.Where("status = ?", dataDic.Status)
	}

	if parentIds != nil && len(parentIds) > 0 {
		db = db.Where("parent_id in (?)", parentIds)
	}

	err = db.Scopes(common.AndLike(columns, vals)).Count(&count).Scopes(
		common.Paginate(pageInfo.Page, pageInfo.PageSize),
	).Find(&dataDics).Error

	if err != nil {
		return message.QUERY_ERR, dataDics, count, err
	}

	return message.QUERY_OK, dataDics, count, err
}

var DataDicServiceApp = new(DataDicService)
