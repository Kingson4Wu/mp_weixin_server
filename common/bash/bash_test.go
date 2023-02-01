package bash_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/bash"
	"testing"
)

func TestExistProcName(t *testing.T) {

	fmt.Println(bash.ExecShellCmd("ls -al"))
}
