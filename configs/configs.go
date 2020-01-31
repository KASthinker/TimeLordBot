package configs

import (
	"log"

	"github.com/BurntSushi/toml"

	"github.com/KASthinker/TimeLordBot/internal/methods"
)

const path = "/configs/helpconf.toml"

// Token ...
type Token struct {
	Token string `toml:"token"`
}

// DataBase ...
type DataBase struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	DBname   string `toml:"dbname"`
}

// GetToken will return the value of the token
func GetToken() string {
	var tkn Token

	if _, err := toml.DecodeFile(methods.GetPath(path), &tkn); err != nil {
		log.Fatalf("Token not received: %v", err)
	} else {
		log.Println("config.toml decoded!")
	}
	return tkn.Token
}

// Configs return config list
func Configs() *DataBase {
	var db DataBase

	if _, err := toml.DecodeFile(methods.GetPath(path), &db); err != nil {
		log.Fatalf("Configs not received: %v\n%v", err, db)
	} else {
		log.Println("config.toml decoded!")
	}
	return &db
}
