package common

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	"reflect"
)

func SetStructField(data any, name string, val any) error {
	refDataVal := reflect.Indirect(reflect.ValueOf(val))
	if refDataVal.Kind() != reflect.Struct {
		return errors.New(global.Msg.Msg(message.DATA_STRUCT_ERR))
	}

	fieldVal := refDataVal.FieldByName(name)
	if fieldVal.Kind() != reflect.ValueOf(val).Kind() {
		return errors.New("field type err: " + global.Msg.Msg(message.DATA_STRUCT_ERR))
	}
	if !fieldVal.CanSet() {
		return errors.New("Can not set err: " + global.Msg.Msg(message.DATA_STRUCT_ERR))
	}

	fieldVal.Set(reflect.ValueOf(val))
	return nil
}
