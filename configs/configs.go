package configs

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/markbates/pkger"
)

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
	info, err := pkger.Info("")
	if err != nil {
		log.Fatalf("Pkger error: %v", err)
	}

	path := info.Dir + "/configs/helpconf.toml"
	log.Printf("PATH: %v", path)

	if _, err := toml.DecodeFile(path, &tkn); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return tkn.Token
}

// Configs return config list
func Configs() *DataBase {
	var db DataBase
	info, err := pkger.Info("")
	if err != nil {
		log.Fatalf("Pkger error: %v\n", err)
	}

	path := info.Dir + "/configs/helpconf.toml"
	log.Printf("PATH: %v", path)

	if _, err = toml.DecodeFile(path, &db); err != nil {
		log.Fatalf("Configs not received: %v\n%v", err, db)
	} else {
		log.Println("config.toml decoded!")
	}
	return &db
}
