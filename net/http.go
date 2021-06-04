package net

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

const (
	ReqModePost = "POST"
	ReqModeGet  = "GET"
	ReqModePut  = "PUT"
)

/*
 * HttpGet
 */
func HttpGet(api string, header map[string]string) ([]byte,error) {

	// create client
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,// support https
				// 	MaxIdleConnsPerHost: Str2Int(GetConfigValue("http", "max_conn")),
			},
		},
		// Timeout: time.Duration(Str2Int(GetConfigValue("http", "timeout"))) * time.Millisecond,
	}

	// new http 'GET' request
	req, newClientErr := http.NewRequest(ReqModeGet, api, nil)
	if newClientErr != nil {
		return nil,newClientErr
	}

	// set header
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// send http request && receive response
	res, httpErr := client.Do(req)

	if httpErr != nil {
		return nil,httpErr
	}

	// recycle resource
	defer res.Body.Close()

	body,readErr:=ioutil.ReadAll(res.Body)

	if res.StatusCode != http.StatusOK {
		return body,fmt.Errorf("request fail: http status code is %d",res.StatusCode)
	}

	return body,readErr
}

/*
 * HttpPost
 */
func HttpPost(
	api string,
	data interface{},
	header map[string]string,
	isUseClientCert bool,
	certFile,
	keyFile string) ([]byte,error) {

	var tr *http.Transport

	// use tls client cert
	if isUseClientCert {
		cert, certErr := tls.LoadX509KeyPair(certFile, keyFile)
		if certErr != nil {
			return nil,certErr
		}

		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,// support https
				Certificates:       []tls.Certificate{cert},
			},
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,// support https
			},
		}
	}

	// create client
	client := &http.Client{Transport: tr}

	// set body
	var bodyStr string
	if data == nil {
		bodyStr = ""
	} else {
		if reflect.TypeOf(data).Kind() == reflect.String {
			bodyStr = data.(string)
		} else if reflect.TypeOf(data).Kind() == reflect.Map && reflect.TypeOf(data) == reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf("")) {
			payload := url.Values{}
			for k, v := range data.(map[string]string) {
				payload.Set(k, v)
			}
			bodyStr = payload.Encode()
		} else if reflect.ValueOf(data).Kind() == reflect.Struct {
			bodyByte, _ := json.Marshal(data)
			bodyStr = string(bodyByte)
		} else {
			return nil,fmt.Errorf("HttpPost: Error data type")
		}
	}

	// new http 'POST' request
	req, newClientErr := http.NewRequest(ReqModePost, api, strings.NewReader(bodyStr))
	if newClientErr != nil {
		return nil,newClientErr
	}

	// set header
	req.Header.Set("Content-Type", "application/x-www-validation-urlencoded")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	// send http request && receive response
	res, httpErr := client.Do(req)
	if httpErr != nil {
		return nil,httpErr
	}

	// recycle resource
	defer res.Body.Close()

	body,readErr:=ioutil.ReadAll(res.Body)

	// http status code
	if res.StatusCode != http.StatusOK {
		return body,fmt.Errorf("request fail: http status code is %d",res.StatusCode)
	}

	return body,readErr
}
