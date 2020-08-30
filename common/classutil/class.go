package classutil

import "reflect"

func ClassName(value interface{}) string {
	return reflect.ValueOf(value).Elem().Type().Name()
}
