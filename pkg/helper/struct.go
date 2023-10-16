package helper

import (
	"errors"
	"reflect"
)

// MergeStruct merge 2 struct, with param 1 as target and param 2 as source, return interface{}
// and it compare by field name.
// !!! noted: param 1 must be pointer of struct
func MergeStruct(obj1 interface{}, obj2 interface{}) error {
	m1 := make(map[string]int)
	m2 := make(map[string]int)

	val1 := reflect.ValueOf(obj1)
	if val1.Kind() == reflect.Ptr {
		val1 = val1.Elem()
	}

	if val1.Kind() != reflect.Struct {
		return errors.New("param 1 not a struct")
	}

	val2 := reflect.ValueOf(obj2)
	if val2.Kind() == reflect.Ptr {
		val2 = val2.Elem()
	}

	if val2.Kind() != reflect.Struct {
		return errors.New("param 2 not a struct")
	}

	for i := 0; i < val1.NumField(); i++ {
		m1[val1.Type().Field(i).Name] = i
	}

	for i := 0; i < val2.NumField(); i++ {
		m2[val2.Type().Field(i).Name] = i
	}

	for i := range m1 {
		// skip if FieldName on Struct 1 not exist on Struct 2
		if _, ok := m2[i]; !ok {
			continue
		}

		f1 := val1.Field(m1[i])
		f2 := val2.Field(m2[i])

		// skip if field 2 is nil
		if f2.Kind() == reflect.Ptr && f2.IsNil() {
			continue
		}

		if f1.Kind() == reflect.Ptr {
			f1 = f1.Elem()
		}

		if f1.Kind() == f2.Kind() && f1.Kind() != reflect.Struct {
			val1.Field(m1[i]).Set(val2.Field(m2[i]))
		}
	}

	return nil
}
