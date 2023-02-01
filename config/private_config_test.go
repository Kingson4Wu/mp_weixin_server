package config_test

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/config"
	"testing"
)

func TestGetConfig(t *testing.T) {

	fmt.Println(config.GetDatabaseConfig().Host)
	fmt.Println(config.GetDatabaseConfig().Host)
}
