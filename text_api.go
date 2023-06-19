package main

import (
	"encoding/json"
	"strconv"
)

type CreateTextResponse struct {
	TextId int `json:"text_id"`
}

type ListTextsResponse []struct {
	Id    int    `json:"id"`
	BoxId int    `json:"box_id"`
	Text  string `json:"text"`
}

type GetTextResponse struct {
	Id    int    `json:"id"`
	BoxId int    `json:"box_id"`
	Text  string `json:"text"`
}

type UpdateTextResponse struct {
	Ok bool `json:"ok"`
}

type DelteTextResponse struct {
	Ok bool `json:"ok"`
}

func CreateTextApi(api_origin string, token string, box_id int, text string) (*CreateTextResponse, error) {
	raw_response, err := PostApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/text/", map[string]interface{}{"text": text})

	if err != nil {
		return nil, err
	}

	var response CreateTextResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func ListTextsApi(api_origin string, token string, box_id int) (*ListTextsResponse, error) {
	raw_response, err := GetApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/text/")

	if err != nil {
		return nil, err
	}

	var response ListTextsResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func GetTextApi(api_origin string, token string, box_id int, text_id int) (*GetTextResponse, error) {
	raw_response, err := GetApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/text/"+strconv.Itoa(text_id)+"/")

	if err != nil {
		return nil, err
	}

	var response GetTextResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func UpdateTextApi(api_origin string, token string, box_id int, text_id int, text string) (*UpdateTextResponse, error) {
	raw_response, err := PutApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/text/"+strconv.Itoa(text_id)+"/", map[string]interface{}{"text": text})

	if err != nil {
		return nil, err
	}

	var response UpdateTextResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}

func DeleteTextApi(api_origin string, token string, box_id, text_id int) (*DelteTextResponse, error) {
	raw_response, err := DeleteApi(api_origin, token, "/v1/box/"+strconv.Itoa(box_id)+"/text/"+strconv.Itoa(text_id)+"/")

	if err != nil {
		return nil, err
	}

	var response DelteTextResponse
	json.Unmarshal([]byte(raw_response.Body()), &response)

	return &response, nil
}
