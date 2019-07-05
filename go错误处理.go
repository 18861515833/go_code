package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	res, err := divice(1, 2)
	//var err error

	if err != nil {
		fmt.Println(err.Error()) //err.Error())
	} else {

		fmt.Println(res)
	}
}

func divice(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数为0")
	}
	return float64(a) / float64(b), nil
}
