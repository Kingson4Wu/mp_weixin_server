package file_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/common/file"
	"testing"
)

func TestGetAllFile(t *testing.T) {

	fmt.Println(file.GetAllFile("."))
}
