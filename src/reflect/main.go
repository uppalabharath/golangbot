package main

import (
	"fmt"
	"reflect"
)

type order struct {
	id       string
	quantity int
	address
}

type address struct {
	city, state string
	zipCode     int
}

func main() {
	o := order{"test", 10, address{"a", "b", 123456}}
	processField(reflect.ValueOf(o))
}

func processField(fieldValue reflect.Value) {
	numFields := fieldValue.NumField()
	for i := 0; i < numFields; i++ {
		internalFieldValue := fieldValue.Field(i)
		kindField := internalFieldValue.Kind()
		fmt.Println("--------------------------------")
		var value interface{}
		switch kindField {
		case reflect.Int:
			value = internalFieldValue.Int()
			fmt.Println("int:", kindField, value)
		case reflect.String:
			value = internalFieldValue.String()
			fmt.Println("string:", kindField, value)
		case reflect.Bool:
			value = internalFieldValue.Bool()
			fmt.Println("boolean:", kindField, value)
		case reflect.Struct:
			fmt.Println("struct", kindField, internalFieldValue)
			processField(internalFieldValue)
		default:
			value = internalFieldValue
			fmt.Println("internal:", kindField, value)
		}
		fmt.Println("--------------------------------")
	}

}
