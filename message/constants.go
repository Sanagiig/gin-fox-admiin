package message

const (
	ZH = "zh"
	EN = "en"
)

const (
	// Base
	CAPTCHA_OK         = iota + 1 //验证通过
	CAPTCHA_FAIL                  // 验证码校验失败
	CAPTCHA_ERR                   // 验证码校验错误
	CAPTCHA_TIME_OUT              // 校验超市
	CAPTCHA_OVER_TIMES            // 验证码验证次数过多

	// 用户相关
	USER_IS_EXIST             // 用户已存在
	USERNAME_OR_PASS_FAILED   // 账号密码错误
	PASSWORD_NOT_INCONSISTENT //  密码不一致
	LOGIN_SUCCESS             // 登录失败
	LOGIN_ERR                 // 登录失败

	// 角色
	ROLE_NOT_EXIST       // 角色不存在
	SOME_ROLES_NOT_EXIST // 部分角色不存在

	// 权限
	AUTH_NOT_EXIST      // 权限不存在
	SOME_AUTH_NOT_EXIST // 部分权限不存在
	AUTH_FAILED
	TOKEN_IS_BLOCK   // TOKEN 被禁用
	TOKEN_IS_EXPIRED //TOKEN 已过期,
	// 请求相关
	REQ_DATA_ERR               // 请求数据错误
	OPER_DB_ERR                // 操作数据错误
	OPER_OK                    // 操作成功
	OPER_FAILED                // 操作失败
	OPER_ERR                   // 操作错误
	QUERY_OK                   // 查询成功
	QUERY_FAILED               // 查询失败
	QUERY_ERR                  // 查询错误
	DATA_NOT_EXIST             // 数据不存在
	DATA_EXIST                 // 数据已存在
	SOME_DATA_NOT_EXIST        // 部分数据不存在
	SOME_PARENT_DATA_NOT_EXIST // 部分父节点数据不存在
	DATA_STRUCT_ERR            // 数据结构错误
)
