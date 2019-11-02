package database

import (
	"database/sql"
	"fmt"
	"log"

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
	err error
	db  *sql.DB
)

//Connect ...
func Connect() (*sql.DB, error) {
	conf := configs.Config()
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", conf.User, conf.Password, conf.DBname))

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
	defer db.Close()
	err := db.QueryRow(fmt.Sprintf("SHOW TABLES LIKE '%v'", userID))

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// NewUser ...
func NewUser(lang, tz string, userID int64) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE %v%v%v (
			id INT NOT NULL AUTO_INCREMENT,
			type_task VARCHAR(15) NOT NULL,
			text VARCHAR(255) NOT NULL,
			date DATE,
			time TIME NOT NULL,
			weekday VARCHAR(70),
			priority VARCHAR(20) NOT NULL,
			PRIMARY KEY (id)
		);`,"`", userID,"`"))
	if err != nil {
		log.Printf("\n\nError in add in Users\n%v\n\n\n", err)
		return err
	}
	_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO Users (user_id, language, timezone) 
		VALUES (%v, %v, %v);`, userID, lang, tz))
	if err != nil {
		log.Printf("\n\nError in insert line\n%v\n\n\n", err)
		return err
	}

	return nil
}
