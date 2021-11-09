package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(1, 1)
	result_type := reflect.TypeOf(result).String()
	if result_type != "int" {
		t.Error("result type is not int but got", result_type)
	}
}
