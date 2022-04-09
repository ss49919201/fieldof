package fieldof

import (
	"errors"
	"reflect"
)

func FieldOf(v any) ([]string, error) {
	val := reflect.ValueOf(v)
	switch val.Type().Kind() {
	case reflect.Struct:
		// valid value
	case reflect.Ptr:
		return FieldOf(val.Elem().Interface())
	default:
		return nil, errors.New("invalid type arg")
	}

	if !val.IsValid() {
		return nil, errors.New("invalid value")
	}

	fields := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Type().Field(i).Name
	}
	return fields, nil
}
