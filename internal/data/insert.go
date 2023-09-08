package data

import (
	"github.com/803Studio/kptl-user-center/internal/config"
	"github.com/803Studio/kptl-user-center/internal/model"
)

var InsertIntoUserAccount = createInsertApi[*model.UserAccount](
	config.AppConfig.Maria.Tables.Users,
)
