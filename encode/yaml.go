package encode

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func YamlLoaded(filePath string) interface{} {
	var yamlByte []byte
	yamlByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		panic(ioErr)
	}

	var data interface{}
	yamlUnmarshalErr := yaml.Unmarshal(yamlByte, &data)
	if yamlUnmarshalErr != nil {
		panic(yamlUnmarshalErr)
	}

	return data
}
