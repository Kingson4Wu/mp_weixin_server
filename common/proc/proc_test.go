package proc_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/proc"
	"testing"
)

func TestExistProcName(t *testing.T) {

	fmt.Println(proc.ExistProcName("java"))
}
