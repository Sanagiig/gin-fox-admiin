package common

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	"gin-one/model/common/request"
	"gorm.io/gorm"
)

func CreateModel(model any) (err error) {
	err = global.DB.Create(model).Error
	return err
}

func UpdateModel(model any) (err error) {
	err = global.DB.Updates(model).Error
	return
}

func DeleteModelByID(model any, id string) (msgCode int, err error) {
	var isExist bool
	isExist, err = IsExist(model, id)
	if err != nil {
		return message.OPER_ERR, err
	} else if !isExist {
		return message.DATA_NOT_EXIST, nil
	}

	err = global.DB.Where("id = ?", id).Delete(model).Error
	if err != nil {
		return message.OPER_DB_ERR, err
	}
	return message.OPER_OK, nil
}

func DeleteModelList(model any, ids []string) (msgCode int, err error) {
	var isExist bool

	// 校验数据完整性
	isExist, err = IsAllExist(model, ids)
	if err != nil {
		return message.REQ_DATA_ERR, err
	} else if !isExist {
		return message.SOME_DATA_NOT_EXIST, nil
	}

	err = global.DB.Model(model).Delete("id in ?", ids).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return message.DATA_NOT_EXIST, err
		}
		return message.OPER_DB_ERR, err
	}
	return message.OPER_OK, nil
}

func GetCountByID(model any, ids []string) (msgCode int, err error) {
	return
}

func GetModel(model any) (int, error) {
	err := global.DB.Where(model).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return message.DATA_NOT_EXIST, err
		}
		return message.OPER_DB_ERR, err
	}
	return message.QUERY_OK, nil
}

func GetModelByID(model any, id string) (codeMsg int, err error) {
	err = global.DB.Where("id = ?", id).First(model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return message.DATA_NOT_EXIST, err
	}
	return message.QUERY_OK, err
}

func GetModelByIds(model any, ids []string, res any) (codeMsg int, err error) {
	var isExist bool
	isExist, err = IsAllExist(model, ids)
	if err != nil || !isExist {
		return message.SOME_DATA_NOT_EXIST, err
	}

	err = global.DB.Model(model).Where("id in ?", ids).Find(res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}

	return message.QUERY_OK, nil
}

func GetModelPagination(model any, pageInfo request.PageInfo, columns []string, vals []string, res any) (codeMsg int, err error) {
	err = global.DB.Model(model).
		Scopes(PaginateAndLike(
			pageInfo.Page, pageInfo.PageSize, columns, vals,
		)).Find(res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}
	return message.QUERY_OK, err
}

func GetAndLikeDB(model any, columns []string, vals []string) *gorm.DB {
	return global.DB.Model(model).Scopes(AndLike(columns, vals))
}
