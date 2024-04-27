package common

import (
	"gin-one/global"
)

func IsExist(model any, id string) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where("id = ?", id).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func IsModelDataExist(model any) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where(model).Count(&count).Error
	if err != nil || count == 0 {
		return false, err
	}

	return true, nil
}

func IsAllExist(model any, ids []string) (bool, error) {
	var count int64

	err := global.DB.Model(model).Where("id in ?", ids).Count(&count).Error
	if err != nil || int(count) != len(ids) {
		return false, err
	}

	return true, nil
}
