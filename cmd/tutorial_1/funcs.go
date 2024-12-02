package main

import (
	"errors"
	"fmt"
)

func main2() {
	printMe("hey")
	var res, mod, err = intDivision(5, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	print(res, " ", mod)
}

func printMe(val string) {
	fmt.Println(val)
}

func intDivision(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	return num / den, num % den, err
}
