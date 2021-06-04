package encode

import (
	"encoding/base64"
	"errors"
	"reflect"
)

func Base64Encoded(msg string) string {
	headerStrByte := []byte(msg)
	return base64.StdEncoding.EncodeToString(headerStrByte)
}

func Base64Decoded(msg string) (string, error) {
	msgByte, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", errors.New("base64 decode fail")
	}

	return string(msgByte), nil
}

// example:
//m := map[string]interface{}{"a": 1, "b": "a"}
//type ta struct {
//	A int    `map:"a"`
//	B string `map:"b"`
//}
//var tt ta
//map2struct(m, &tt)

func Map2struct(m interface{}, s interface{}) {
	//fmt.Println(reflect.ValueOf(m).Kind())
	//fmt.Println(reflect.ValueOf(s).Kind())
	//fmt.Println(reflect.ValueOf(s).Elem().Kind())
	if reflect.ValueOf(m).Kind() == reflect.Map && reflect.ValueOf(s).Elem().Kind() == reflect.Struct {
		var structTag2Name = make(map[string]string)
		refVal := reflect.ValueOf(s).Elem()
		for i := 0; i < refVal.NumField(); i++ {
			structTag2Name[refVal.Type().Field(i).Tag.Get("map")] = refVal.Type().Field(i).Name
		}

		for k, v := range m.(map[string]interface{}) {
			//fmt.Println(k, "==>", v)
			//fmt.Println(structTag2Name[k], "==>", v)
			//
			//fmt.Println(refVal.FieldByName(structTag2Name[k]).Type())
			//fmt.Println(reflect.ValueOf(v).Type())
			_, ok := structTag2Name[k]
			if ok {
				refVal.FieldByName(structTag2Name[k]).Set(reflect.ValueOf(v))
			}
		}
	}
}