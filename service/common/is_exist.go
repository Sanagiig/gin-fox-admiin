package common

import (
	"gin-one/global"
	"reflect"
)

// IsExist
//
//	@Description: 判断 ID 是否存在
//	@param model
//	@param id
//	@return bool
//	@return error
func IsExist(model any, id string) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where("id = ?", id).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

// IsModelDataExist
//
//	@Description: 判断删选条件是否能查到数据
//	@param model
//	@return bool
//	@return error
func IsModelDataExist(model any) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where(model).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

// IsAllExist
//
//	@Description: 判断 ids 是否全部都存在
//	@param model
//	@param ids
//	@return bool
//	@return error
func IsAllExist(model any, ids []string) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where("id in ?", ids).Count(&count).Error
	if err != nil || int(count) != len(ids) {
		return false, err
	}

	return true, nil
}

// IsExistParentModel
//
//	@Description: 判断是否存在 ParentModel
//	@param model
//	@return bool
//	@return error
func IsExistParentModel(model any) bool {
	typ := reflect.Indirect(reflect.ValueOf(model)).Type()
	_, hasId := typ.FieldByName("ParentID")
	_, hasName := typ.FieldByName("ParentName")
	return hasId && hasName
}
