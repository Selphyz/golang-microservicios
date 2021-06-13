package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)
var(
	DBClient *sql.DB
)
func init()  {
	os.Setenv("mysql_username", "root")
	os.Setenv("mysql_password", "rootroot")
	os.Setenv("mysql_host", "127.0.0.1:3306")
	os.Setenv("mysql_schema", "users_db")
	var (
		username 	= os.Getenv("mysql_username")
		password 	= os.Getenv("mysql_password")
		host 		= os.Getenv("mysql_host")
		schema 		= os.Getenv("mysql_schema")
	)
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