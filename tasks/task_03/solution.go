package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}
	word := ""
	if n%3 == 0 {
		word += "Fizz"
	}
	if n%5 == 0 {
		word += "Buzz"
	}
	if word == "" {
		word = strconv.Itoa(n)
	}
	return word, nil
}
