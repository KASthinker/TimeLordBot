package localization

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// TrMess translate message
func TrMess(lang string, text string) string {

	type trText map[string]string
	var data trText
    path := "/media/data/Projects/GO/src/github.com/KASthinker/TimeLordBot/internal/localization/"
	file, err := ioutil.ReadFile(path + lang + ".json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	val, ok := data[text]
	if !ok {
		log.Printf("Key map \"%v\" not fiund!", text)
		return ""
	}
	return val
}
