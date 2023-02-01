package aes_test

import (
	"github.com/kingson4wu/mp_weixin_server/common/aes"
	"testing"
)

func TestEncryptByAesWithKey(t *testing.T) {
	var (
		in       = "kxw"
		expected = "aDBA5g4yEryqeDty8fuW3A=="
	)
	actual, _ := aes.EncryptByAesWithKey(in, "ABCDABCDABCDABCD")
	if actual != expected {
		t.Errorf("EncryptByAesWithKey(%s) = %s; expected %s", in, actual, expected)
	}
}

func TestDecryptByAesWithKey(t *testing.T) {
	var (
		in       = "aDBA5g4yEryqeDty8fuW3A=="
		expected = "kxw"
	)
	actual, _ := aes.DecryptByAesWithKey(in, "ABCDABCDABCDABCD")
	if actual != expected {
		t.Errorf("DecryptByAesWithKey(%s) = %s; expected %s", in, actual, expected)
	}
}
