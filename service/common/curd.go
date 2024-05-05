package common

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	"gin-one/utils/helper"
	"gorm.io/gorm"
	"reflect"
)

type GetChildrenListByParentIds = func(parentIds []string, model any) (msgCode int, res any, err error)

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

func GetChildrenByParents(db *gorm.DB, parentIds []string, res any) (codeMsg int, err error) {
	if parentIds != nil && len(parentIds) > 0 {
		db = db.Where("parent_id in ?", parentIds)
	} else {
		db = db.Where("parent_id IS NULL")
	}

	err = db.Find(res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}
	return message.QUERY_OK, nil
}

func GetTreeList() {

}

func GetModelNameByIds(model any, ids []string) (codeMsg int, data []map[string]string, err error) {
	err = global.DB.Model(model).Select("id,name").Where("id in ?", ids).Find(&data).Error
	if err != nil {
		return message.QUERY_ERR, nil, err
	}

	return message.QUERY_OK, data, nil
}

func GetModelByParentIds(model any, parentIds []string, res any) (codeMsg int, err error) {
	var isExist bool
	isExist, err = IsAllExist(model, parentIds)
	if err != nil || !isExist {
		return message.SOME_PARENT_DATA_NOT_EXIST, err
	}

	err = global.DB.Model(model).Where("parent_id in ?", parentIds).Find(res).Error
	if err != nil {
		return message.QUERY_ERR, err
	}

	return message.QUERY_OK, nil
}

func GetParentName(model any, data any) (codeMsg int, err error) {
	var eleTyp reflect.Type
	var dataLength = 1
	var isSlice = false
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	typ := reflect.Indirect(reflect.ValueOf(data)).Type()
	dataVal := reflect.Indirect(reflect.ValueOf(data))

	switch typ.Kind() {
	case reflect.Slice, reflect.Array:
		dataLength = dataVal.Len()
		eleTyp = typ.Elem()
		isSlice = true
	default:
		eleTyp = typ
	}

	if dataLength == 0 {
		return message.QUERY_OK, nil
	}

	// 入参不匹配 或 不存在 parent 字段
	if eleTyp.Kind() != modelType.Kind() || !IsExistParentModel(model) {
		return message.DATA_STRUCT_ERR, nil
	}

	var parentList = make([]map[string]interface{}, 0, dataLength)
	if isSlice {
		//var modelList = make([]map[string]string, 0, dataLength)
		var parentIdList = make([]string, 0, dataLength)
		var ids = make([]string, 0, dataLength)
		// 已有 ID
		for i := 0; i < dataLength; i++ {
			id := dataVal.Index(i).FieldByName("ID").String()
			if !helper.HasEle(parentIdList, id) {
				ids = append(ids, id)
				// 将已存在的 Id-name 对应到 map
				parentList = append(parentList, map[string]interface{}{
					"id":   id,
					"name": dataVal.Index(i).FieldByName("Name").String(),
				})
			}
		}

		// 筛选需要查找的 ID
		for i := 0; i < dataLength; i++ {
			pid := dataVal.Index(i).FieldByName("ParentID").String()
			if !helper.HasEle(parentIdList, pid) && !helper.HasEle(ids, pid) {
				parentIdList = append(parentIdList, pid)
			}
		}

		err = global.DB.Model(model).Select("id,name").Where("id in ?", parentIdList).Limit(dataLength).Find(&parentList).Error
		if err != nil {
			return message.QUERY_ERR, err
		}

		// 更新值
		for i := 0; i < dataLength; i++ {
			pid := dataVal.Index(i).FieldByName("ParentID").String()
			for j := 0; j < len(parentList); j++ {
				if pid == parentList[j]["id"] {
					dataVal.Index(i).FieldByName("ParentName").Set(reflect.ValueOf(parentList[j]["name"]))
				}
			}
		}
	} else {
		parentId := dataVal.FieldByName("ParentID").String()
		err = global.DB.Model(model).Select("id", "name").
			Where("id = ?", parentId).
			Limit(dataLength).Find(&parentList).Error
		if err != nil {
			return message.QUERY_ERR, err
		}

		if len(parentList) > 0 {
			dataVal.FieldByName("ParentName").Set(reflect.ValueOf(parentList[0]["name"]))
		}
	}

	return message.QUERY_OK, nil
}
