package main

import (
	"errors"
	"fmt"

	"github.com/digitalhouse-content/go-fundamentals-bases/errors/mypackage"
)

type MyCustomErr struct {
	ID string
}

func (m MyCustomErr) Error() string {
	return fmt.Sprintf("error with id: %s", m.ID)
}

func main() {

	var err error
	err = errors.New("my new error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("my format err, string: %s, number: %d", "my string", 231)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	err3 := TestError(1)
	fmt.Println(errors.As(err3, &MyCustomErr{}))

	err4 := TestError(4)
	fmt.Println(errors.As(err4, &MyCustomErr{}))

	err5 := errors.Join(err, err2, err3)
	fmt.Println(err5)

	fmt.Println(errors.Is(err5, err2))
	fmt.Println(errors.Is(err5, err4))

	err6 := fmt.Errorf("error1: [%w]", err)
	err7 := fmt.Errorf("error2: [%w]", err6)
	fmt.Println(err7)
	fmt.Println(errors.Unwrap(err7))
	fmt.Println(errors.Unwrap(errors.Unwrap(err7)))
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err7))))

	defer func() {
		fmt.Println("end main")
		r := recover()
		if r != nil {
			fmt.Println("Recover in ", r)
		}
	}()
	mypackage.Run()
	panic("my panic")
}

func TestError(v int) error {
	if v == 1 {
		return MyCustomErr{"4"}
	} else {
		return errors.New("my error")
	}
}
