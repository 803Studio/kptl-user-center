package data

import (
	"github.com/803Studio/kptl-user-center/internal/config"
	"github.com/803Studio/kptl-user-center/internal/model"
)

var SelectUserAccountByWxId = createSelectByApi[string, *model.UserAccount](
	config.AppConfig.Maria.Tables.Users,
	"wxid",
	func() *model.UserAccount {
		return new(model.UserAccount)
	},
)
