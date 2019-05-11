package main

import (
	"fmt"
	"reflect"
)

// api function
type ApiFunc func() string

// function adpter
func ApiOption(fn ApiFunc, isActive bool) ApiFunc {
	if isActive {
		return func() string {
			return fmt.Sprintf("wrap [%s]", fn())
		}
	}
	return fn
}

func Get() string { return "get" }
func Put() string { return "put" }
var OptGet = ApiOption(Get, false)
var OptPut = ApiOption(Put, true)

type Api struct {
	Get ApiFunc `active:"true"`
	Put ApiFunc `active:"false"`
}

var emptyArgs = []reflect.Value{}

func call(api Api, method string) string {
	t := reflect.ValueOf(&api).Elem()
	value := t.FieldByName(method)
	typeField, _ := t.Type().FieldByName(method)
	fval := value.Call(emptyArgs)
	if typeField.Tag.Get("active") == "true" {
		return fmt.Sprintf("wrap[%v]", fval[0])
	}
	return fval[0].String()
}

func main() {
	api := Api{Get: Get, Put: Put}
	println("api get>", call(api, "Get"))
	println("api put>", call(api, "Put"))

	println("opt get>", OptGet())
	println("opt put>", OptPut())
}
