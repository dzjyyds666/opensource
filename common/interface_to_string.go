package common

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ToStringWithoutError(data interface{}) string {
	val := reflect.ValueOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		// 如果不是指针也不是结构体，则直接输出为字符串
		return val.String()
	} else {
		typ := val.Type()
		str := make(map[string]interface{}, typ.NumField())
		for i := range val.NumField() {
			field := typ.Field(i)
			value := val.Field(i)

			if value.Kind() == reflect.Ptr {
				value = value.Elem()
			}

			if value.Kind() == reflect.Struct {
				// 如果是结构体，则递归调用
				str[field.Name] = ToStringWithoutError(value.Interface())
			} else {
				str[field.Name] = fmt.Sprintf("%v", value.Interface())
			}
		}
		marshal, _ := json.Marshal(str)
		resultStr := string(marshal)
		// 去除 \ todo 比较抽象的操作，后续可以优化
		resultStr = strings.ReplaceAll(resultStr, "\\", "")
		return resultStr
	}
}
