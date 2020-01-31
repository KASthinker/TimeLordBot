package methods

import (
	"log"

	"github.com/markbates/pkger"
)

//GetPath returns the path to the main working directory (if path = "")
//or the path to the file inside the application, if it's path is specified.
func GetPath(path string) string {
	info, err := pkger.Info("")
	if err != nil {
		log.Fatalf("Error GetPath(): %v", err)
	}

	return info.Dir + path
}
