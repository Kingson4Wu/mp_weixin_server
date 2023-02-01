package admin_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/admin"
	"testing"
)

func TestIsAdministrator(t *testing.T) {

	fmt.Println(admin.IsAdministrator("llll"))
}
