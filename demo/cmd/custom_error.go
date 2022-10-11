package main

import (
	"fmt"
)

type BusinessError struct {
	Code    int
	Message string
}

func (b BusinessError) Error() string {
	return fmt.Sprintf("Business Error with Code %d", b.Code)
}

func RunCustomError() {
	r, err := doWithError(false)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}
}

func doWithError(boolean bool) (int, error) {
	if boolean {
		return 1, nil
	} else {
		return 0, BusinessError{Code: 400}
	}
}
