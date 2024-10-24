package main

import (
	"errors"
	"testing"
)

func TestHandler_WrongNumberOfArgs(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)

	if !errors.Is(err, errWrongNumberOfArgs) {
		t.Error("wrong error")
	}
}

func TestHandler_InvalidFirstArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"INVALID", "2"})
	if !errors.Is(err, errInvalidArg) {
		t.Error("wrong error")
	}
}
