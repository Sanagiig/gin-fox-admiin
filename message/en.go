package message

type en struct {
	msgMap map[int]string
}

func (e *en) init() {
	e.msgMap = map[int]string{
		USER_IS_EXIST: "用户已存在",
	}
}

func (e *en) Msg(code int) string {
	msg, ok := e.msgMap[code]
	if ok {
		return msg
	}
	return ""
}

func newEn() *en {
	msg := &en{}
	msg.init()
	return msg
}
