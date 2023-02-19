package swagger

import (
	"git.sample.ru/sample/internal/logger"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type Swagger struct {
	Path string
}

func GetSwaggerHandler(path string) *Swagger {
	return &Swagger{
		path,
	}
}

func (s Swagger) GetSwaggerJsonHandler(w http.ResponseWriter, r *http.Request, path string) {
	path, err := filepath.Abs(path)
	if err != nil {
		logger.Error.Print(err)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error.Print(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,UPDATE,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")

	_, err = w.Write(file)
	if err != nil {
		logger.Error.Print(err)
	}

	return
}
