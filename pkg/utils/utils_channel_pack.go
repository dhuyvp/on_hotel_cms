package utils

import (
	"encoding/json"
	"fmt"
	"hotel_cms/app/models"
	"reflect"
)

func GetColumnsAndValuesChannelPackManage(p models.ChannelPackManage) (string, string) {
	var myMap map[string]interface{}
	data, _ := json.Marshal(p)
	json.Unmarshal(data, &myMap)

	v := reflect.ValueOf(p)
	typeMap := make(map[string]string)
	for i := 0; i < v.NumField(); i += 1 {
		fieldType := v.Field(i).Type()
		fieldName := v.Type().Field(i).Name

		typeMap[fieldName] = fmt.Sprintf("%v", fieldType)
	}

	queryColumns := ""
	queryValues := ""
	for keyMap, valMap := range myMap {
		if valMap == nil {
			continue
		}
		queryColumns += "," + keyMap
		if typeMap[keyMap] == "*string" {
			queryValues += ",'" + fmt.Sprintf("%v", valMap) + "'"
		} else {
			queryValues += "," + fmt.Sprintf("%v", valMap)
		}
	}

	if queryColumns == "" {
		queryColumns = ","
	}
	if queryValues == "" {
		queryValues = ","
	}
	return queryColumns[1:], queryValues[1:]
}

func GetQueryUpdateChannelPackManage(p models.ChannelPackManage) string {
	var myMap map[string]interface{}
	data, _ := json.Marshal(p)
	json.Unmarshal(data, &myMap)

	v := reflect.ValueOf(p)
	typeMap := make(map[string]string)
	for i := 0; i < v.NumField(); i += 1 {
		fieldType := v.Field(i).Type()
		fieldName := v.Type().Field(i).Name

		typeMap[fieldName] = fmt.Sprintf("%v", fieldType)
	}

	queryUpdate := ""
	for keyMap, valMap := range myMap {
		if valMap == nil {
			continue
		}
		queryUpdate += "," + keyMap + "="
		if typeMap[keyMap] == "*string" {
			queryUpdate += "'" + fmt.Sprintf("%v", valMap) + "'"
		} else {
			queryUpdate += fmt.Sprintf("%v", valMap)
		}
	}

	if queryUpdate == "" {
		queryUpdate = ","
	}
	return queryUpdate[1:]
}
