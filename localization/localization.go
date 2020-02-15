package localization

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var language = map[string]string{
	"en_EN": "en_EN.json",
	"ru_RU": "ru_RU.json",
}

// Translate srting
func Translate(lang string, typeText string, text string) string {

	type trText map[string]map[string]string
	var data trText
	path := "./localization/lang/"

	if lang, ok := language[lang]; ok {
		path = path + lang
	} else {
		log.Fatal("Invalid language!")
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%v: %v", err, path)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	val, ok := data[typeText][text]
	if !ok {
		log.Printf("Key map \"[%v][%v]\" not fiund!", typeText, text)
		return ""
	}

	return val
}
