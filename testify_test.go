package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MyMockedObject is a mocked object that implements an interface
// that describes an object that the code I am testing relies on.
type MyMockedObject struct {
	mock.Mock
}

// DoSomething is a method on MyMockedObject that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.

func (m *MyMockedObject) isDividedBy3(nbr int) bool {
	args := m.Called(nbr)
	return args.Bool(0)
}

func (m *MyMockedObject) isDividedBy5(nbr int) bool {
	args := m.Called(nbr)
	return args.Bool(0)
}

func (m *MyMockedObject) isDividedBy3And5(nbr int) bool {
	args := m.Called(nbr)
	return args.Bool(0)
}

func TestFizzBuzz_testify(t *testing.T) {
	a := assert.New(t)
	testObj := new(MyMockedObject)

	testObj.On("isDividedBy3And5", 15).Return(true)
	defer testObj.AssertExpectations(t)

	result := fizzBuzzFunc(testObj, 15)
	a.Equal("FizzBuzz", result, `fizzBuzzFunc(checker, 15) = `+result)
}
