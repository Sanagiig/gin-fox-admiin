package data

import (
	"errors"
	"gin-one/global"
	"gin-one/message"
	"reflect"
)

func Copy[T any](dst, src T) error {
	var typ = reflect.TypeOf(dst)
	switch typ.Kind() {
	case reflect.Struct:
		return CopyStruct(dst, src)
	default:
		return errors.New(global.Msg.Msg(message.DATA_STRUCT_ERR))
	}
}

func CopyStruct(dst, src any) error {
	return nil
}
