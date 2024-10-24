package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github/smartyjenny/calc-lib"
)

func main() {

	handler := NewHandler(os.Stdout, &calc_lib.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}

type Handler struct {
	stdout     *os.File
	calculator *calc_lib.Addition
}

func NewHandler(stdout *os.File, calculator *calc_lib.Addition) Handler {
	return Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) < 2 {
		return errWrongNumberOfArgs
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w:%w", errInvalidArg, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	result := this.calculator.Calculate(a, b)

	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return err
	}

	return nil
}

var (
	errWrongNumberOfArgs = errors.New("usages: calc [a] [b]")
	errInvalidArg        = errors.New("invalidArgument")
)
