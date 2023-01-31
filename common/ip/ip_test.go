package ip_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/ip"
	"testing"
)

func TestGetExtranetIp(t *testing.T) {

	fmt.Println(ip.GetExtranetIp())
}
