//go:generate mockgen -source=main.go -destination=checker_mock.go -package=main

package main

import (
	"os"
	"strconv"
)

type Checker interface {
	isDividedBy3(nbr int) bool
	isDividedBy5(nbr int) bool
	isDividedBy3And5(nbr int) bool
}

type fizzBuzz struct{}

func (fizzBuzz) isDividedBy3(nbr int) bool {
	return nbr%3 == 0
}

func (fizzBuzz) isDividedBy5(nbr int) bool {
	return nbr%5 == 0
}

func (fizzBuzz) isDividedBy3And5(nbr int) bool {
	return nbr%3 == 0 && nbr%5 == 0
}

func main() {
	checker := fizzBuzz{}
	nbr, _ := strconv.Atoi(os.Args[1])
	fizzBuzzFunc(checker, nbr)
}

func fizzBuzzFunc(checker Checker, nbr int) string {
	if checker.isDividedBy3And5(nbr) {
		return "FizzBuzz"
	}
	if checker.isDividedBy3(nbr) {
		return "Fizz"
	}
	if checker.isDividedBy5(nbr) {
		return "Buzz"
	}
	return strconv.Itoa(nbr)
}
