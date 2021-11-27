package math_test

import (
	"fmt"
	"github.com/mrminglang/go-captcha/captcha_api/math"
	"testing"
)

func TestGenerateMathCaptcha(t *testing.T) {
	id, b64s, err := math.NewMathCaptchaDriver().
		SetRequiredParam(100, 60).
		GenerateMathCaptcha()
	fmt.Println(id, b64s, err)
	if err != nil {
		t.Errorf("GenerateMathCaptcha: %v\n", err)
	}
}

func TestVerify(t *testing.T) {
	id, b64s, err := math.NewMathCaptchaDriver().
		SetRequiredParam(100, 60).
		GenerateMathCaptcha()
	if err != nil {
		t.Errorf("Verify: %v\n", err)
	}
	_ = b64s
	if math.Verify(id, "5") {
		t.Errorf("Verify: want <false>, get <true>")
	}
}
