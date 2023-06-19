package main

import (
	"github.com/go-resty/resty/v2"
)

func GetApi(api_origin string, token string, path string) (*resty.Response, error) {
	return resty.New().R().SetAuthToken(token).Get(api_origin + path)
}

func PostApi(api_origin string, token string, path string, body map[string]interface{}) (*resty.Response, error) {
	return resty.New().R().SetAuthToken(token).SetBody(body).Post(api_origin + path)
}

func PutApi(api_origin string, token string, path string, body map[string]interface{}) (*resty.Response, error) {
	return resty.New().R().SetAuthToken(token).SetBody(body).Put(api_origin + path)
}

func DeleteApi(api_origin string, token string, path string) (*resty.Response, error) {
	return resty.New().R().SetAuthToken(token).Delete(api_origin + path)
}
