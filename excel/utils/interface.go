package utils

import (
	"errors"
	"reflect"
)

func ToInterfaceSlice(arr interface{}) ([]interface{}, error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
		return nil, errors.New("arr is not array or slice")
	}
	var sli []interface{}
	for i := 0; i < rv.Len(); i++ {
		sli = append(sli, rv.Index(i).Interface())
	}
	return sli, nil
}
