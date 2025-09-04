package main

import (
	"fmt"
	"reflect"
)

func main() {
	data := []interface{}{
		"asdfa",
		3241,
		true,
		make(chan any),
		nil,
	}
	for _, v := range data {
		typeDeterminator(v)
	}
}

func typeDeterminator(data any) any {
	value := reflect.ValueOf(data)
	kind := value.Kind()

	switch kind {
	case reflect.Int:
		fmt.Println("data type is integer")
		determinedData := data.(int)
		return determinedData
	case reflect.String:
		fmt.Println("data type is string")
		determinedData := data.(string)
		return determinedData
	case reflect.Bool:
		fmt.Println("data type is bool")
		determinedData := data.(bool)
		return determinedData
	case reflect.Chan:
		cType := value.Type().String()
		fmt.Printf("data type is %s \n", cType)
		return data
	}
	fmt.Println("type can't be determinated")
	return nil
}
