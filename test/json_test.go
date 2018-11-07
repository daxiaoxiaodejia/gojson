package test

import (
	"encoding/json"
	"fmt"
	"gojson"
	"testing"
)


func Test_Validator(t *testing.T) {
	data := `{"id":5240423333333333333333,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	err := gojson.CheckValid([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
}

func Test_json(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json22(t *testing.T) {
	data := `{
  "id": 524042,
  "name": "酷旅-mob-otv-2",
  "male": true,
  "other": null
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_object(t *testing.T) {
	data := `{
  "id": -524042.5,
  "name": "酷旅-mob-otv-2",
  "male": true,
  "children": {
    "id": -524042.5,
    "name": "酷旅-mob-otv-2",
    "male": true,
    "other": null
  },
  "other": null
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_array(t *testing.T) {
	data := `{
  "id": [
    -524042.5,
    231231.2
  ],
  "name": "酷旅-mob-otv-2",
  "male": true,
  "other": null
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}

func Test_json_array111(t *testing.T) {
	data := `{
  "id": [
    -524042.5,
    2312314444455555555555555
  ],
  "name": "酷旅-mob-otv-2",
  "male": true,
  "other": null
}`
	object, err := gojson.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}