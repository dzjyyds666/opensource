package captcah

import (
	"github.com/mojocn/base64Captcha"
)

type CaptchaServer struct {
	Id     string `json:"id"`
	B64s   string `json:"b64s"`
	Answer string `json:"answer"`
}

func GetCaptcha(height, width, length, dotCount int, maxSkew float64) (captchaServer *CaptchaServer, err error) {

	driver := base64Captcha.NewDriverDigit(height, width, length, maxSkew, dotCount)

	store := base64Captcha.DefaultMemStore

	//生成图形验证码
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &CaptchaServer{
		Id:     id,
		B64s:   b64s,
		Answer: answer,
	}, nil
}

func VerifyCaptcha(id, answer string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, answer, clear)
}
