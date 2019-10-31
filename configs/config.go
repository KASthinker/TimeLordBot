package configs

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/KASthinker/TimeLordBot/"
)

// Token ...
type Token struct {
	Token string `tom;:"Token"`
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
	//path := "/media/data/Projects/GO/src/github.com/KASthinker/TimeLordBot/configs/"
	help := TimeLordBot.configs.helpconf
	if _, err := toml.DecodeFile(help, &tkn); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return tkn.Token
}

// Config ...
func Config() *DataBase {
	var db DataBase
	path := "/media/data/Projects/GO/src/github.com/KASthinker/TimeLordBot/configs/"
	if _, err := toml.DecodeFile(path+"helpconf.toml", &db); err != nil {
		log.Fatalf("Token not received: %v", err)
	}
	return &db
}
