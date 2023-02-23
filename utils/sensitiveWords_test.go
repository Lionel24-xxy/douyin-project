package utils

import (
	"testing"
)

func TestCheckAndReplace(t *testing.T) {
	if err := SensitiveWordInit(); err != nil {
		t.Errorf(err.Error())
	}
	text := "你傻/逼"
	isContain := SensitiveWordCheck(text, 1)
	if isContain != true {
		t.Errorf("error")
	}
	replace := SensitiveWordReplace(text)
	if replace != "你**" {
		t.Errorf("error")
	}
}