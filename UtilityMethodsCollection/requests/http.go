package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

/*
*网络请求封装
 */

type ClientOption struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Body    map[string]interface{}
}

//Get http Get method
func (option ClientOption) Get() (string, error) {
	//new request
	req, err := http.NewRequest("GET", option.Url, nil)
	if err != nil {
		log.Println(err)
		return "0", err
	}
	//add params
	q := req.URL.Query()
	if option.Params != nil {
		for key, val := range option.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if option.Headers != nil {
		for key, val := range option.Headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	//log.Printf("Go GET URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return "0", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body) //一定要关闭res.Body
	//读取body
	resBody, err := ioutil.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return "0", err
	}

	return string(resBody), nil
}

//Post http post method
func (option *ClientOption) Post() (string, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if option.Body != nil {
		var err error
		bodyJson, err = json.Marshal(option.Body)
		if err != nil {
			log.Println(err)
			return "0", errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", option.Url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return "0", errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if option.Params != nil {
		for key, val := range option.Params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if option.Headers != nil {
		for key, val := range option.Headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return "0", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body) //一定要关闭res.Body
	//读取body
	resBody, err := ioutil.ReadAll(res.Body) //把  body 内容读入字符串 s
	if err != nil {
		return "0", err
	}

	return string(resBody), nil
}
