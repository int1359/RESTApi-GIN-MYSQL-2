package config

import (
	"database/sql"
	"fmt"
	"rest/api/util"
)

var DB *sql.DB

func DbURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.GetEnvVariable("DB_USER"),
		util.GetEnvVariable("DB_PASSWORD"),
		util.GetEnvVariable("DB_HOST"),
		util.GetEnvVariable("DB_PORT"),
		util.GetEnvVariable("DB_NAME_STUDENT"))

}
