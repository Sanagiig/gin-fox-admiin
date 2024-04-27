package initialize

import (
	"gin-one/global"
	"gin-one/message"
)

func Msg() *message.Message {
	return message.New(global.Config.System.Lang)
}
