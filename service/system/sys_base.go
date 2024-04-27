package system

import (
	"gin-one/global"
	"gin-one/message"
	"gin-one/model/common/response"
	"github.com/mojocn/base64Captcha"
	"time"
)

var store = base64Captcha.DefaultMemStore

// var defaultDriver =
type BaseService struct{}

func (service *BaseService) GetCaptcha(key string) (response.CaptchaRes, error) {
	openCaptcha := global.Config.Captcha.OpenCaptcha
	openCaptchaTimeout := global.Config.Captcha.OpenCaptchaTimeOut

	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeout))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}

	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight, global.Config.Captcha.ImgWidth, global.Config.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64Str, code, err := cp.Generate()

	return response.CaptchaRes{
		ID:          id,
		Base64Str:   base64Str,
		CaptchaLen:  global.Config.Captcha.KeyLong,
		OpenCaptcha: oc,
		CaptchaCode: code,
	}, err
}

func (service *BaseService) CheckCaptchaCode(key string, id string, code string) (int, bool, error) {
	var times int = 1
	openCaptcha := global.Config.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.Config.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	} else {
		times = interfaceToInt(v)
		times++
		global.BlackCache.Set(key, times, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc = openCaptcha == 0 || openCaptcha >= times
	if !oc {
		return message.CAPTCHA_OVER_TIMES, false, nil
	}

	ok = store.Verify(id, code, true)
	if !ok {
		return message.CAPTCHA_FAIL, false, nil
	}
	global.BlackCache.Delete(key)
	return message.CAPTCHA_OK, true, nil
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

var BaseServiceApp = new(BaseService)
