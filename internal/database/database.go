package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
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
	strUserID := fmt.Sprintf("'%v'", userID)
	row := db.QueryRow(fmt.Sprintf("SHOW TABLES LIKE %v;", strUserID))
	err = row.Scan()
	if err == sql.ErrNoRows {
		log.Printf("\n\n\n%v", row.Scan(err))
		return false
	}
	return true
}

// NewUser ...
func NewUser(user *data.UserData, userID int64) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	strUserID := fmt.Sprintf("`%v`", userID)
	_, err = db.Exec(fmt.Sprintf(`
		CREATE TABLE %v (
			id INT NOT NULL AUTO_INCREMENT,
			type_task VARCHAR(15) NOT NULL,
			text VARCHAR(255) NOT NULL,
			date DATE,
			time TIME NOT NULL,
			weekday VARCHAR(70),
			priority VARCHAR(20) NOT NULL,
			PRIMARY KEY (id)
		);`,strUserID))
	if err != nil {
		log.Printf("\n\nError in add User table\n%v\n\n\n", err)
		return err
	}
	_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO Users (user_id, language, timezone) 
		VALUES (%v, '%v', '%v');`, userID, user.Language, user.Timezone))
	if err != nil {
		log.Printf("\n\nError in insert line\n%v\n\n\n", err)
		return err
	}

	return nil
}

// GetUserData ...
func GetUserData(userID int64, user *data.UserData) {
	strUserID := fmt.Sprintf("'%v'", userID)
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	row := db.QueryRow(fmt.Sprintf(
		"SELECT language, timezone FROM Users WHERE user_id=%v", strUserID))
	err = row.Scan(&user.Language, &user.Timezone)
	if err == sql.ErrNoRows {
		log.Printf("\n\n\n%v\n\n\n", err)
	}
	log.Printf("\n\n\nOK-> %v:%v\n\n\n", user.Language, user.Timezone)
}

// DeleteUserAccount ...
func DeleteUserAccount(userID int64) error{
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	
	strUserID := fmt.Sprintf("'%v'", userID)
	_, err = db.Exec(fmt.Sprintf(`DELETE FROM Users WHERE user_id=%v;`, strUserID))
	if err != nil {
		log.Printf("\n\nError in delete line\n%v\n\n\n", err)
		return err
	}
	strUserID = fmt.Sprintf("`%v`", userID)
	_, err = db.Exec(fmt.Sprintf(`DROP TABLE %v;`, strUserID))
	if err != nil {
		log.Printf("\n\nError DROP TABLE\n%v\n\n\n", err)
		return err
	}

	return nil
}

// ChangeLanguage ... 
func ChangeLanguage(userID int64, lang string) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	
	strUserID := fmt.Sprintf("'%v'", userID)
	_, err = db.Exec(fmt.Sprintf("UPDATE Users SET language='%v' WHERE user_id=%v", lang, strUserID))
	if err != nil {
		log.Printf("\n\nError in update line\n%v\n\n\n", err)
		return err
	}

	return nil
}

// ChangeTimeZone ...
func ChangeTimeZone(userID int64, tz string) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	
	strUserID := fmt.Sprintf("'%v'", userID)
	_, err = db.Exec(fmt.Sprintf("UPDATE Users SET timezone='%v' WHERE user_id=%v", tz, strUserID))
	if err != nil {
		log.Printf("\n\nError in update line\n%v\n\n\n", err)
		return err
	}

	return nil
}