package convert_typescript

import (
	"os"
	"time"

	tf "github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

var registeredTypes map[string][]interface{}

func Generate(dir string) error {
	os.MkdirAll(dir, os.ModePerm)
	for namespace, models := range registeredTypes {
		converter := tf.New()
		converter.ManageType(time.Time{}, tf.TypeOptions{TSType: "Date", TSTransform: "new Date(__VALUE__)"})
		converter.BackupDir = "" // don't backup
		for _, model := range models {
			converter.Add(model)
		}
		err := converter.WithInterface(true).ConvertToFile(dir + "/" + namespace + ".ts")
		if err != nil {
			return err
		}
	}
	return nil
}

func Register(namespace string, models ...interface{}) {
	_, ok := registeredTypes[namespace]
	if !ok {
		registeredTypes[namespace] = make([]interface{}, 0)
	}
	registeredTypes[namespace] = append(registeredTypes[namespace], models...)
}
