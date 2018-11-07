# go-jsonObject
json字符串校验、反序列化JSON字符串成Golang对象、反序列化JSON字符串成Golang代码

## 安装
1. 使用go get
```bash
go get github.com/ChengjinWu/gojson
```
2. 导入包
```go
import (
	"github.com/ChengjinWu/gojson"
)
```

## 主要特性
* json字符串校验
* 反序列化JSON字符串成Golang对象
* 反序列化JSON字符串成Golang结构体、get、set方法

## 示例
#### json字符串有效性校验
```go
func Test_Validator(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	err := gojson.CheckValid([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
}
```
#### json字符串转Golang对象
```go
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
```
#### json字符串内嵌数组转Golang对象
```go
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
```
