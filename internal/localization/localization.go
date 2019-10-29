package localization

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// TrMess translate message
func TrMess(lang string, typeText string, text string) string {

	type trText map[string]map[string]string
	var data trText
	//Изменить path
	path := "/media/data/Projects/GO/src/github.com/KASthinker/TimeLordBot/internal/localization/"
	///////////////
	file, err := ioutil.ReadFile(path + lang + ".json")
	if err != nil {
		log.Fatal(err)
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
