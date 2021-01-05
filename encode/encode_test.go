package encode

import "testing"

type People struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func TestEncode(t *testing.T) {
	jsonStr := JsonEncoded(People{Name: "Vincent", Age: 29})
	t.Log(jsonStr)

	//var p = People{}
	//p = JsonDecoded(jsonStr, p).(People)
	//t.Log(p)
}
