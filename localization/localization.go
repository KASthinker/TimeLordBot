package localization

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/KASthinker/TimeLordBot/internal/methods"
)

// Translate srting
func Translate(lang string, typeText string, text string) string {

	type trText map[string]map[string]string
	var data trText

	path := methods.GetPath("/localization/lang/")
	path = path + lang + ".json"

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(path + " not found!")
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