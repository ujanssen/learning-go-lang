package db

import (
	"encoding/json"
	"github.com/devbliss/acd/model"
	"io/ioutil"
	"log"
)

type DatabaseModel interface {
	Filename() string
}

func filename(name string) string {
	return  name + ".json"
}

func WriteModelToFile(data DatabaseModel) error {
	name := data.Filename()

	log.Println("WriteDataToFile: ", data, filename(name))

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename(name), b, 0666)
}

func ReadModelFromFile(data DatabaseModel) error {
	name := data.Filename()
	log.Println("ReadDataFromFile: ", filename(name))

	b, err := ioutil.ReadFile(filename(name))
	if err != nil {
		return err
	}
	log.Println("Got: ", string(b))

	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	return nil
}
