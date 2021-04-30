package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestGetFieldName(t *testing.T) {
	n := struct {
		Name string
		Age  string
	}{}
	strArr, _ := ReflectGetFieldName(n)
	t.Log(strArr)
}

func TestGetFieldTagName(t *testing.T) {
	n := struct {
		Name string `json:"name" db:"localhost"`
		Age  string `json:"age" db:"localhost"`
	}{}
	strArr, _ := ReflectGetTagName(n)
	t.Log(strArr)
}

// Gets the name of the field in the struct
func ReflectGetFieldName(structName interface{}) ([]string, error) {
	// 'TypeOf' is core func
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	} else if t.Kind() != reflect.Struct {
		return nil, errors.New("not struct error, please check type")
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result, nil
}

// Gets the value of the Tag in the struct, or returns the field value if there is no Tag
func ReflectGetTagName(structName interface{}) ([]string, error) {
	// 'TypeOf' is core func
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	} else if t.Kind() != reflect.Struct {
		return nil, errors.New("not struct error, please check type")
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		tagName := t.Field(i).Name
		tags := strings.Split(string(t.Field(i).Tag), "\"")
		if len(tags) > 1 {
			// get first tag value
			tagName = tags[1]
		}
		result = append(result, tagName)
	}
	return result, nil
}
