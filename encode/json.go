package encode

import (
	"encoding/json"
	"io/ioutil"
)

/*
 * Json Encoded
 */
func JsonEncoded(data interface{}) string {

	dataByte, jsonMarshalErr := json.Marshal(data)

	if jsonMarshalErr != nil {
		panic(jsonMarshalErr)
	}

	return string(dataByte)
}

/*
 * Json Decoded
 */
func JsonDecoded(jsonStr string, data interface{}) interface{} {

	var dataByte = []byte(jsonStr)
	//var data interface{}

	jsonUnmarshalErr := json.Unmarshal(dataByte, &data)

	if jsonUnmarshalErr != nil {
		panic(jsonUnmarshalErr)
	}

	return data
}

/*
 * Json Loaded
 */
func JsonLoaded(filePath string) interface{} {
	var jsonByte []byte
	jsonByte, ioErr := ioutil.ReadFile(filePath)
	if ioErr != nil {
		panic(ioErr)
	}

	var data interface{}
	jsonUnmarshalErr := json.Unmarshal(jsonByte, &data)
	if jsonUnmarshalErr != nil {
		panic(jsonUnmarshalErr)
	}

	return data
}
