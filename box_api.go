package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type CreateBoxResponse struct {
	BoxId int `json:"box_id"`
}

type ListBoxesResponse []struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetBoxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateBoxResponse struct {
	Ok bool `json:"ok"`
}

type DelteBoxResponse struct {
	Ok bool `json:"ok"`
}

func CreateBoxApi(api_origin string, token string, name string) (*CreateBoxResponse, error) {
	raw_response, err := PostApi(api_origin, token, "/v1/box/", map[string]interface{}{"name": name})

	if err != nil {
		return nil, err
	}

	var response CreateBoxResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func ListBoxesApi(api_origin string, token string) (*ListBoxesResponse, error) {
	raw_response, err := GetApi(api_origin, token, "/v1/box/")

	if err != nil {
		return nil, err
	}

	var response ListBoxesResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func GetBoxApi(api_origin string, token string, box_id int) (*GetBoxResponse, error) {
	raw_response, err := GetApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/")

	if err != nil {
		return nil, err
	}

	var response GetBoxResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func UpdateBoxApi(api_origin string, token string, box_id int, name string) (*UpdateBoxResponse, error) {
	raw_response, err := PutApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/", map[string]interface{}{"name": name})
	fmt.Println(raw_response)
	if err != nil {
		return nil, err
	}

	var response UpdateBoxResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func DeleteBoxApi(api_origin string, token string, box_id int) (*DelteBoxResponse, error) {
	raw_response, err := DeleteApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/")

	if err != nil {
		return nil, err
	}

	var response DelteBoxResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}
