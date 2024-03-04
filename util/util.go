package util

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// import "reflect"

// func ContainsRequiredFields(obj *interface{}) bool {
// 	// va
// 	// val := reflect.ValueOf(obj).NumField()
// 	// for i := 0; i<val.NumField(); i++ {
// 	// 	field := val.Field(i)
// 	// 	field.Name
// 	// }
// 	return false
// }

// func GetField(obj *reflect.StructField, field string) (reflect.StructField, bool) {
// 	return reflect.TypeOf(obj).FieldByName(field)
// }

// func GetTag(field *reflect.StructField, tag string) (string, bool) {
// 	result, ok := field.Tag.Lookup(tag)
// 	return string(result), ok
// }
