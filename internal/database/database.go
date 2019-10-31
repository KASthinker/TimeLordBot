package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/KASthinker/TimeLordBot/configs"
)

// Users ...
type Users struct {
	UserID      int
	TypeAccount string
	TimeZone    string
	GroupID     string
}

var (
	err  error
	db   *sql.DB
	once sync.Once
)

//Connect ...
func Connect() (*sql.DB, error) {
	once.Do(func() {
		conf := configs.Config()
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", conf.User, conf.Password, conf.DBname))
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// IfUserExists ...
func IfUserExists(userID int64) bool {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	var count int
	err := db.QueryRow(fmt.Sprintf("SHOW TABLES LIKE '%v'", userID)).Scan(&count)
	
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
