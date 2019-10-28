package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/KASthinker/TimeLordBot/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("ALTER DATABASE TimeLordBot charset=utf8;")

	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
	
	_, err = db.Exec(
		`CREATE TABLE Users (
            user_id INT NOT NULL,
            type_account VARCHAR(6) NOT NULL DEFAULT 'User',
            timezone VARCHAR(3) NOT NULL,
            group_id VARCHAR(255),
            PRIMARY KEY (user_id)
		);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")

	_, err = db.Exec(
		`
		CREATE TABLE Groups (
			id INT NOT NULL AUTO_INCREMENT,
			group_name VARCHAR(255) NOT NULL,
			group_timezone VARCHAR(3) NOT NULL,
			creator INT NOT NULL,
			PRIMARY KEY (id)
		);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
}
