package main

import (
	"fmt"
	"zuoguai/internal/config"
	"zuoguai/internal/db"
)

func main() {
	config := config.GetConfigs()
	fmt.Println(
		"user:", config.Mysql.User,
		"password:", config.Mysql.Password,
		"host:", config.Mysql.Host,
		"port:", config.Mysql.Port,
		"database:", config.Mysql.Database,
	)

	db := db.GetDB()
	fmt.Println(db)

}
