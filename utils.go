package main

import "reflect"

func ToInt(value reflect.Value) int {
	return int(value.Int())
}

func ToStr(value reflect.Value) string {
	return value.String()
}
