package dbSql

import (
	"fmt"
	"tirupatiBE/config"
	"tirupatiBE/dal/dbModel"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type dbAccess struct {
	db *sqlx.DB
}

func SqlConnection(cnf *config.Config) (dbModel.RouterFunc, error) {

	var dba dbAccess

	dsn := cnf.Db.EngineConfigStruct.User + ":" + cnf.Db.EngineConfigStruct.Password + "@tcp(" + cnf.Db.EngineConfigStruct.Host + ":)/" + cnf.Db.EngineConfigStruct.DB + "?parseTime=true"

	db := sqlx.MustConnect("mysql", dsn)

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		panic("Error to connetion with db")
	}

	db.SetMaxOpenConns(cnf.Db.EngineConfigStruct.MaxOpenConn)
	db.SetMaxIdleConns(cnf.Db.EngineConfigStruct.MaxIdelConn)

	dba.db = db

	return &dba, nil

}
