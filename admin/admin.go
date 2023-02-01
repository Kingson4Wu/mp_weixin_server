package admin

import "github.com/kingson4wu/mp_weixin_server/config"

func IsAdministrator(account string) bool {

	admins :=
		config.GetAdminConfig().Accounts

	if len(admins) == 0 {
		return false
	}

	for _, a := range admins {
		if account == a {
			return true
		}
	}
	return false
}
