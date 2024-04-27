package validator

var (
	PageInfoVerify = Rules{
		"Page":     {NotEmpty(), Ge("1"), Le("1000000")},
		"PageSize": {NotEmpty(), Ge("1"), Le("10000")},
	}
	LoginVerify = Rules{
		"Username":  {NotEmpty()},
		"Password":  {NotEmpty()},
		"Captcha":   {NotEmpty()},
		"CaptchaId": {NotEmpty()},
	}
)
