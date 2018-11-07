package gojson

import "reflect"

type JsonObjectType int
const (
	JsonObjectTypeBool JsonObjectType = iota
	JsonObjectTypeString
	JsonObjectTypeInt32
	JsonObjectTypeInt64
	JsonObjectTypeObject
	JsonObjectTypeArray

)

type JsonObject struct {
	Attributes map[string]*JsonObject
	Value interface{}
	VType reflect.Kind
	//dataType

}

func(j *JsonObject) GetString(key string) string {
	return ""
}
