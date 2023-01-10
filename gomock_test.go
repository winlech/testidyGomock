package main

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFizzBuzz_fizzBuzz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	checker := NewMockChecker(ctrl)
	checker.EXPECT().isDividedBy3And5(gomock.Any()).Return(true)

	result := fizzBuzzFunc(checker, 15)
	if result != "FizzBuzz" {
		t.Error(`fizzBuzzFunc(checker, 15) = ` + result)
	}
}
