package parse

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	STRUCT_MUST_TAG = "must"
)

func Must(v interface{}) error {
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)

	switch vType.Kind() {
		case reflect.Ptr:
			if vType.Elem().Kind() != reflect.Struct {
				return fmt.Errorf("input pointer must point to a struct. Input is %v", vType.Kind())
			}

			vType = vType.Elem()
			vValue = reflect.Indirect(vValue)
		case reflect.Struct:
			break
		default: 
			return fmt.Errorf("input must be struct. Input is %v", vType.Kind())
	}
	
	if err := structMust(vType, vValue); err != nil {
		return err
	}

	return nil
}

func structMust(vType reflect.Type, vValue reflect.Value) error {
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		tag := field.Tag.Get(STRUCT_MUST_TAG)
		if tag != "" {
			validValues := strings.Split(tag, ",")
			strField := vValue.Field(i).String()

			if !strInSlice(strField, validValues) {
				return fmt.Errorf("field %s does not match any of the must values", field.Name)
			}
		}
	}

	return nil
}