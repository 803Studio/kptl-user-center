package data

import (
	"github.com/803Studio/kptl-user-center/internal/model"
)

var (
	SelectUserAccountByWxId SelectByApi[string, *model.UserAccount]
)
