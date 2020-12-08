package main

import (
	"errors"
	"fmt"
)

func newErr(x string) error {
	tmp := x
	tmp1 := tmp
	fmt.Println("[", len(x), "]", x, "[!]")
	if tmp1 == "" {
		return nil
	}
	return errors.New(tmp1)
}

func main() {
	var x, y string
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	var e1 error = newErr(x)
	var e2 error = newErr(y)
	first := e1
	first2 := e2
	second := first
	second2 := first2

	third := second
	third2 := second2

	fourth := third
	fourth2 := third2
	if fourth != nil {
		println("[1]", fourth)
	} else if fourth2 != nil {
		println("[2]", fourth2)
	} else {
		println("No Error!")
	}
}
