package math

import (
	bs64Captcha "github.com/mojocn/base64Captcha"
)

var store = bs64Captcha.DefaultMemStore

type MathCaptchaDriver struct {
	*bs64Captcha.DriverMath
}

func NewMathCaptchaDriver() *MathCaptchaDriver {
	return &MathCaptchaDriver{
		&bs64Captcha.DriverMath{},
	}
}

func (driver *MathCaptchaDriver) SetRequiredParam(height, width int) *MathCaptchaDriver {
	driver.Height = height
	driver.Width = width
	return driver
}

func (driver *MathCaptchaDriver) GenerateMathCaptcha() (id, b64s string, err error) {
	driver.ConvertFonts()
	id, b64s, err = bs64Captcha.NewCaptcha(driver, store).Generate()
	return
}

func Verify(id, verifyValue string) bool {
	return store.Verify(id, verifyValue, true)
}
