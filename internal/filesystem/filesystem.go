package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func CreateFile(filename string, body string, dir string) {
	newfile := fmt.Sprintf("%s/%s", dir, filepath.Base(filename))
	var _, err = os.Stat(newfile)

	if !os.IsNotExist(err) {
		log.Info("File exist")

		return
	}

	os.Mkdir(dir, 0755)
	file, err := os.OpenFile(newfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Info("My error is:", err)

		return
	}

	if _, err := file.WriteString(body); err != nil {
		log.Info(err)
	}

	log.Info("File created ", newfile)
}
