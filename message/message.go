package message

import "gin-one/utils/helper"

type Msg interface {
	Msg(int) string
}

type Message struct {
	language Msg
}

var okList = []int{CAPTCHA_OK,
	LOGIN_SUCCESS,
	OPER_OK,
	QUERY_OK}

func (m *Message) Msg(code int) string {
	return m.language.Msg(code)
}

func (m *Message) IsMsg(code int, msg string) bool {
	return m.language.Msg(code) == msg
}

func (m *Message) IsOkCode(code int) bool {
	return helper.HasEle(okList, code)
}

func New(lang string) *Message {
	msg := &Message{}
	switch lang {
	case ZH:
		msg.language = newZh()
	case EN:
		msg.language = newEn()
	default:
		panic("unknown language " + lang)
	}
	return msg
}
