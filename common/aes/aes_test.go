package aes_test

import (
	"fmt"
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

	fmt.Println(aes.DecryptByAesWithKey("6lsTe3Hj3wkPki7tK2Ns3/61f1TBT5irX6PBau8X4eo=", "labali1234labali"))
}
