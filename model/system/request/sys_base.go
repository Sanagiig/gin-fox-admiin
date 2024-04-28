package request

type CheckCaptchaReq struct {
	ID          string `json:"id" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}
