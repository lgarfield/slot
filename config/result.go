package config

import(
	"reflect"
)

type Result struct {
	Status int
	Code int
	Data []reflect.Value
}
