package data

import (
	"database/sql"
	"fmt"
	"github.com/803Studio/kptl-user-center/internal/config"
	"github.com/803Studio/kptl-user-center/internal/model"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func fmtDbConfig() string {
	dbConfig := &config.AppConfig.Maria
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Addr,
		dbConfig.Database,
	)
}

func Init() error {
	var err error = nil
	db, err = sql.Open("mysql", fmtDbConfig())
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(config.AppConfig.Maria.Conn)
	db.SetMaxIdleConns(config.AppConfig.Maria.Conn)

	SelectUserAccountByWxId = createSelectByApi[string, *model.UserAccount](
		config.AppConfig.Maria.Tables.Users,
		"wxid",
		func() *model.UserAccount {
			return new(model.UserAccount)
		},
	)

	InsertIntoUserAccount = createInsertApi[*model.UserAccount](
		config.AppConfig.Maria.Tables.Users,
	)

	return nil
}
