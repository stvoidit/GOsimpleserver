package routers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

// Jsonify - отправка json response
func Jsonify(w http.ResponseWriter, i interface{}, code int) {
	switch i.(type) {
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case string:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	}
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.Encode(i)
}

// JSONLoad - преобразование json в структуру
func JSONLoad(req *http.Request, i interface{}) (interface{}, error) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&i)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// TemplatesMap - ...
type TemplatesMap map[string][]byte

var tmp = make(TemplatesMap)

// RegistrateTemplates - ...
func RegistrateTemplates(tmpPath string) {
	files, _ := ioutil.ReadDir(tmpPath)
	for _, file := range files {
		fullpath := path.Join(tmpPath, file.Name())
		filename := strings.TrimSuffix(file.Name(), ".html")
		bfile, err := ioutil.ReadFile(fullpath)
		if err != nil {
			panic(err)
		}
		tmp[filename] = bfile
	}
}

// RenderTemplate - ...
func RenderTemplate(w http.ResponseWriter, name string) {
	w.Write(tmp[name])
}
