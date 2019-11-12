package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/configs"
	_ "github.com/go-sql-driver/mysql"
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
			date VARCHAR(10),
			time TIME NOT NULL,
			weekday VARCHAR(70),
			priority VARCHAR(20) NOT NULL,
			PRIMARY KEY (id)
		);`, strUserID))
	if err != nil {
		log.Printf("\n\nError in add User table\n%v\n\n\n", err)
		return err
	}
	_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO Users (user_id, language, timezone, time_format) 
		VALUES (%v, '%v', '%v', '%v');`, userID, user.Language, user.Timezone, user.TimeFormat))
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
		"SELECT language, timezone, time_format FROM Users WHERE user_id=%v", strUserID))
	err = row.Scan(&user.Language, &user.Timezone, &user.TimeFormat)
	if err == sql.ErrNoRows {
		log.Printf("\n\n\n%v\n\n\n", err)
	}
	log.Printf("\n\n\nOK-> %v:%v:%v\n\n\n", user.Language, user.Timezone, user.TimeFormat)
}

// DeleteUserAccount ...
func DeleteUserAccount(userID int64) error {
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
	_, err = db.Exec(
		fmt.Sprintf("UPDATE Users SET language='%v' WHERE user_id=%v", lang, strUserID))
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
	_, err = db.Exec(
		fmt.Sprintf("UPDATE Users SET timezone='%v' WHERE user_id=%v", tz, strUserID))
	if err != nil {
		log.Printf("\n\nError in update line\n%v\n\n\n", err)
		return err
	}

	return nil
}

// ChangeTimeFormat ...
func ChangeTimeFormat(userID int64, tf int) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	strUserID := fmt.Sprintf("'%v'", userID)
	_, err = db.Exec(
		fmt.Sprintf("UPDATE Users SET time_format='%v' WHERE user_id=%v", tf, strUserID))
	if err != nil {
		log.Printf("\n\nError in update line\n%v\n\n\n", err)
		return err
	}

	return nil
}

//AddNewTask ...
func AddNewTask(userID int64, task *data.Task) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	strUserID := fmt.Sprintf("`%v`", userID)
	_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO %v (type_task, text, time, date, weekday, priority) 
		VALUES ('%v','%v','%v','%v','%v','%v');`, strUserID, task.TypeTask, task.Text,
		task.Time, task.Date, task.WeekDay, task.Priority))
	if err != nil {
		log.Printf("\n\nError in insert line\n%v\n\n\n", err)
		return err
	}
	return nil
}

// GetTasks ...
func GetTasks(userID int64, typeTask string) ([]data.Task, error) {
	db, err = Connect()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		fmt.Sprintf("SELECT * FROM `%v` WHERE type_task='%v'", userID, typeTask))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	tasks := make([]data.Task, 0)
	for rows.Next() {
		task := new(data.Task)
		err := rows.Scan(&task.ID, &task.TypeTask, &task.Text,
			&task.Date, &task.Time, &task.WeekDay, &task.Priority)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		tasks = append(tasks, *task)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return tasks, nil
}

// DeleteTask ...
func DeleteTask(userID int64, ID int) error {
	db, err = Connect()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	strUserID := fmt.Sprintf("`%v`", userID)
	_, err = db.Exec(fmt.Sprintf("DELETE FROM %v WHERE id=%v;", strUserID, ID))
	if err != nil {
		log.Printf("\n\nError in delete line\n%v\n\n\n", err)
		return err
	}
	return nil
}
