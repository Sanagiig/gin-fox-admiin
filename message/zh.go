package message

type zh struct {
	msgMap map[int]string
}

func (z *zh) init() {
	z.msgMap = map[int]string{
		// Base
		CAPTCHA_OK:         "验证通过",
		CAPTCHA_FAIL:       "验证码校验失败",
		CAPTCHA_ERR:        "验证码校验错误",
		CAPTCHA_TIME_OUT:   "校验码超时",
		CAPTCHA_OVER_TIMES: "验证码验证次数过多",
		// 用户相关
		USER_IS_EXIST:             "用户已存在",
		USERNAME_OR_PASS_FAILED:   "用户名或密码错误",
		PASSWORD_NOT_INCONSISTENT: "密码不一致",
		LOGIN_SUCCESS:             "登录成功",
		LOGIN_ERR:                 "登录失败",
		// 请求相关
		REQ_DATA_ERR:        "请求数据错误",
		OPER_DB_ERR:         "操作数据错误",
		OPER_OK:             "操作成功",
		OPER_FAILED:         "操作失败",
		OPER_ERR:            "操作错误",
		QUERY_OK:            "查询成功",
		QUERY_FAILED:        "查询失败",
		QUERY_ERR:           "查询错误",
		DATA_NOT_EXIST:      "数据不存在",
		SOME_DATA_NOT_EXIST: "部分数据不存在",
		DATA_STRUCT_ERR:     "数据结构错误",
	}
}

func (z *zh) Msg(code int) string {
	msg, ok := z.msgMap[code]
	if ok {
		return msg
	}
	return ""
}

func newZh() *zh {
	msg := &zh{}
	msg.init()
	return msg
}
