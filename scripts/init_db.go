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
	conf := configs.Configs()
	once.Do(func() {
		db, err = sql.Open("mysql", 
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", 
				conf.User, conf.Password, conf.Host, conf.DBname))
	})

	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	} else {
		log.Println("DB open!")
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("ALTER DATABASE %v charset=utf8;", conf.DBname))
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("OK! -> '%v'", fmt.Sprintf("ALTER DATABASE %v charset=utf8;", conf.DBname))
	}

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
		log.Println(err)
	} else {
		log.Println("OK! -> 'CREATE TABLE Users'")
	}

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
		log.Println(err)
	} else {
		log.Println("OK! -> 'CREATE TABLE Groups'")
	}
}
