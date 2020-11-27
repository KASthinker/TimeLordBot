package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

const path = "./configs/configs.toml"

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

	if _, err := toml.DecodeFile(path, &tkn); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return tkn.Token
}

// Configs return config struct
func Configs() DataBase {
	var db DataBase

	if _, err := toml.DecodeFile(path, &db); err != nil {
		log.Fatalf("Configs not received: %v\n%v", err, db)
	}
	return db
}
