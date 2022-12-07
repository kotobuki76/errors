package errors

import (
	"reflect"
	"testing"
)

func TestErrors(t *testing.T) {

	err := New(-1, "test", nil)

	if reflect.TypeOf(err).String() == "error" {
		t.Log(reflect.TypeOf(err).String())
	}

	t.Log(reflect.TypeOf(err.Error()).String())

	t.Log(reflect.TypeOf(err).String())
}

func TestIsErrorItem(t *testing.T) {

	err := New(-1, "test", nil)

	if !IsErrorItem(err) {
		t.Error("error is not ErrorItem")
	} else {
		t.Log("error is ErrorItem")
	}
}

func TestToErrorItem(t *testing.T) {

	str := "test done"

	err := New(-1, "test", str)

	errItem := ToErrorItem(err)

	t.Logf("%v", errItem)
}
