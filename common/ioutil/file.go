package ioutil

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ReadFileContent(path string) ([]byte, error) {

	filePtr, err := os.Open(path)
	if err != nil {
		log.Error("could not open file: "+path, err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(filePtr)
	if err != nil {
		log.Error("could not read file: "+path, err)
		return nil, err
	}

	return bytes, nil
}

func ReadFileAsJson(path string, target interface{}) error {

	content, err := ReadFileContent(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, target)
	if err != nil {
		log.Error("could not unmarshal json file: "+path, err)
		return err
	}

	return nil
}

func ReadFileAsYml(path string, target interface{}) error {

	content, err := ReadFileContent(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(os.ExpandEnv(string(content))), target)
	if err != nil {
		log.Error("could not unmarshal yml file: "+path, err)
		return err
	}

	return nil
}
