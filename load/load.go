package load

import (
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadedYaml(filePath string, data interface{}) error {
	var yamlByte []byte
	yamlByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		panic(ioErr)
	}

	err := yaml.Unmarshal(yamlByte, data)
	if err != nil {
		return err
	}

	return nil
}

func LoadedJson(filePath string, data interface{}) error {
	var jsonByte []byte
	jsonByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		panic(ioErr)
	}

	err := json.Unmarshal(jsonByte, data)
	if err != nil {
		return err
	}

	return nil
}

func LoadedXml(filePath string, data interface{}) error {
	var xmlByte []byte
	xmlByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		panic(ioErr)
	}

	err := xml.Unmarshal(xmlByte, data)
	if err != nil {
		return err
	}

	return nil
}

func LoadedToml(filePath string, data interface{}) error {
	//var tomlByte []byte
	//tomlByte, ioErr := ioutil.ReadFile(filePath)
	//if ioErr != nil {
	//	panic(ioErr)
	//}

	_, err := toml.DecodeFile(filePath, data)
	if err != nil {
		return err
	}

	return nil
}
