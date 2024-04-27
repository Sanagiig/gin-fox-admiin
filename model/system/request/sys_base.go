package request

type CheckCaptchaReq struct {
	ID          string `json:"id"`
	CaptchaCode string `json:"captchaCode"`
}
