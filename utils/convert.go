package utils

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func JsonToStruct(data []byte, s interface{}) error {
	err := json.Unmarshal(data, s)
	if err != nil {
		return err
	}
	return nil
}

func StructToJson(v interface{}) string {
	data, _ := json.Marshal(v)

	return string(data)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func StringToInt(str string) int {
	variable, _ := strconv.Atoi(str)
	return variable
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
