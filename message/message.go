package message

type Msg interface {
	Msg(int) string
}

type Message struct {
	language Msg
}

func (m *Message) Msg(code int) string {
	return m.language.Msg(code)
}

func (m *Message) IsMsg(code int, msg string) bool {
	return m.language.Msg(code) == msg
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
