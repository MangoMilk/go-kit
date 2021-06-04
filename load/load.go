package load

import (
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadedYaml(filePath string, data interface{}) error {
	yamlByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		return ioErr
	}

	if err := yaml.Unmarshal(yamlByte, data);err != nil {
		return err
	}

	return nil
}

func LoadedJson(filePath string, data interface{}) error {
	jsonByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		return ioErr
	}

	if err := json.Unmarshal(jsonByte, data);err != nil {
		return err
	}

	return nil
}

func LoadedXml(filePath string, data interface{}) error {
	xmlByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		return ioErr
	}

	if err := xml.Unmarshal(xmlByte, data);err != nil {
		return err
	}

	return nil
}

func LoadedToml(filePath string, data interface{}) error {
	if _, err := toml.DecodeFile(filePath, data);err != nil {
		return err
	}

	return nil
}
