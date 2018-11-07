package gojson

import "reflect"

type JsonObject struct {
	Attributes map[string]*JsonObject
	Value      interface{}
	VType      reflect.Kind
}

// 获取属性对象
func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	if j == nil {
		return nil
	}
	return j.Attributes[key]
}

// 获取string值
func (j *JsonObject) GetString() string {
	if j == nil && j.VType != reflect.String {
		return ""
	}
	return j.Value.(string)
}

// 获取数组
func (j *JsonObject) GetJsonArray() []*JsonObject {
	if j == nil && j.VType != reflect.Slice {
		return nil
	}
	return j.Value.([]*JsonObject)
}

// 获取bool值
func (j *JsonObject) GetBool() bool {
	if j == nil && j.VType != reflect.Bool {
		return false
	}
	return j.Value.(bool)
}

// 获取浮点值
func (j *JsonObject) GetFloat64() float64 {
	if j == nil {
		return 0
	}
	switch j.VType {
	case reflect.Int64:
		return float64(j.Value.(int64))
	case reflect.Float64:
		return j.Value.(float64)
	}
	return 0
}

// 获取整型值
func (j *JsonObject) GetInt64() int64 {
	if j == nil {
		return 0
	}
	switch j.VType {
	case reflect.Int64:
		return j.Value.(int64)
	case reflect.Float64:
		return int64(j.Value.(float64))
	}
	return 0
}

// 获取整型值
func (j *JsonObject) GetInterface() interface{} {
	if j == nil && j.VType != reflect.Struct {
		return nil
	}
	return j.Value
}
