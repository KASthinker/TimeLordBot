package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/KASthinker/TimeLordBot/configs"
	_ "github.com/go-sql-driver/mysql"
)

var (
	err  error
	db   *sql.DB
	once sync.Once
)

func main() {
	once.Do(func() {
		conf := configs.Config()
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", conf.User, conf.Password, conf.DBname))
	})
	defer db.Close()

	if err != nil {
		log.Println(err)
	}
	//defer db.Close()

	db.Exec("ALTER DATABASE TimeLordBot charset=utf8;")

	_, err = db.Exec(
		`CREATE TABLE Users (
			user_id INT NOT NULL,
			language VARCHAR(6) NOT NULL,
            type_account VARCHAR(6) DEFAULT 'User',
			timezone VARCHAR(3) NOT NULL,
			time_format INT NOT NULL,
            group_id VARCHAR(255),
            PRIMARY KEY (user_id)
		);`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("OK")

	_, err = db.Exec(
		`
		CREATE TABLE Groups (
			id INT NOT NULL AUTO_INCREMENT,
			group_name VARCHAR(255) NOT NULL,
			group_timezone VARCHAR(3) NOT NULL,
			time_format INT NOT NULL,
			creator INT NOT NULL,
			PRIMARY KEY (id)
		);`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("OK")
}
