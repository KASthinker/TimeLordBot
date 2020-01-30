package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Token ...
type Token struct {
	Token string `toml:"Token"`
}

// DataBase ...
type DataBase struct {
	User     string `toml:"User"`
	Password string `toml:"Password"`
	Host     string `toml:"Host"`
	DBname   string `toml:"DBname"`
}

// GetToken ...
func GetToken() string {
	var tkn Token
	path := ""
	if _, err := toml.DecodeFile("", &tkn); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return tkn.Token
}

// Configs return config list
func Configs() *DataBase {
	var db DataBase
	path := ""
	if _, err := toml.DecodeFile("", &db); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return &db
}