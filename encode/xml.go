package encode

import (
	"encoding/xml"
)

func XMLDecoded(xmlStr string) interface{} {

	var dataByte = []byte(xmlStr)
	var data interface{}

	xmlUnmarshalErr := xml.Unmarshal(dataByte, &data) //将文件转化成对象

	if xmlUnmarshalErr != nil {
		panic(xmlUnmarshalErr)
	}

	return data
}

func XMLEncoded() {

}
