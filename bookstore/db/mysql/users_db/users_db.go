package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_username = "root"
	mysql_password = "rootroot"
	mysql_host = "127.0.0.1:3306"
	mysql_schema = "users_db"
)

var (
	DBClient *sql.DB
	username = os.Getenv(mysql_username)
	password = os.Getenv(mysql_password)
	host = os.Getenv(mysql_host)
	schema = os.Getenv(mysql_schema)
)

func init()  {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	DBClient, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = DBClient.Ping(); err != nil {
		panic(err)
	}
	log.Println("Mysql successfully configured")
}