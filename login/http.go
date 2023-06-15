package login

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/**
 *  http GET 请求数据 仅支持json格式的
 * 	resp_data := make(map[string]interface{}) // resp_data 也可以是结构体
 *	err := Get("http://127.0.0.1:8080/info", &resp_data)
 */
func Get(address string, dest interface{}) error {
	resp, err := request("GET", address, nil)
	if err != nil {
		log.Print("[GET][REQ]" + err.Error())
	}
	err = json.Unmarshal(resp, dest)
	if err != nil {
		log.Print("[GET][Unmarshal]" + err.Error())
	}
	return err
}

func Post(address string, data map[string]interface{}, dest interface{}) error {
	resp, err := request("POST", address, data)
	if err != nil {
		log.Print("[POST][REQ]" + err.Error())
	}
	err = json.Unmarshal(resp, dest)
	if err != nil {
		log.Print("[POST][Unmarshal]" + err.Error())
	}
	return err
}

func PostForm(address string, data map[string]interface{}, dest interface{}) error {
	resp, err := request("FORM", address, data)
	if err != nil {
		log.Print("[POSTFORM][REQ]" + err.Error())
	}
	err = json.Unmarshal(resp, dest)
	if err != nil {
		log.Print("[POSTFORM][Unmarshal]" + err.Error())
	}
	return err
}

func request(method string, address string, data map[string]interface{}) ([]byte, error) {

	method = strings.ToUpper(method)
	var resp *http.Response
	var err error
	switch method {
	case "GET":
		resp, err = http.Get(address)
	case "POST":
		data_code, _ := json.Marshal(data)
		resp, err = http.Post(address, "application/json", bytes.NewBuffer(data_code))
	case "FORM":
		form_data := make(url.Values)
		for key, value := range data {
			form_data[string(key)] = []string{value.(string)}
		}
		resp, err = http.PostForm(address, form_data)
	}

	if err != nil {
		log.Print("[HTTP][method:%s][url:%s][param:%s]error:", method, address, data, err.Error())
		return []byte{}, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("[HTTP][BODY][READ]%s", err.Error())
	}
	return body, err
}
