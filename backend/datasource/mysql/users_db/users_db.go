package users_db

import (
	"database/sql"
	"log"
)

func init() {
	_ "github.com/go-sql-driver/mysql"
}


var (
	Client *sql.DB
	username = "root"
	password = "password"
	host = "127.0.0.1:3306"
	schema = "users_db_02"
)

func init(){
	// username:password@tcp(host)/user_schema
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", uusername, papassword, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err := Client.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}