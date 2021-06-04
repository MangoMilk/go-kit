package net

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T)  {
	res,err:=HttpGet(
		"https://baidu.com",
		nil,
		)

	if err!= nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}

func TestPost(t *testing.T)  {
	res,err:=HttpPost(
		"http://127.0.0.1/v1/session",
		map[string]string{
			"shop_id":"1",
			"user_open_id":"1",
		},
		map[string]string{"Content-Type":"application/json"},
		false,
		"",
		"")

	if err!= nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}