package main

import "reflect"
import "os"

func ToInt(value reflect.Value) int {
	return int(value.Int())
}

func ToStr(value reflect.Value) string {
	return value.String()
}

func ValidDir(path string) bool {
	info, err := os.Stat(path)

	return err == nil && info.IsDir()
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
