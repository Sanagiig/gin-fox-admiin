package response

type CaptchaRes struct {
	ID          string `json:"id"`
	Base64Str   string `json:"base64Str"`
	CaptchaLen  int    `json:"captchaLen"`
	OpenCaptcha bool   `json:"openCaptcha"`
	CaptchaCode string `json:"-"`
}
