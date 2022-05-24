package goangelapi

import (
	"errors"
	"fmt"
	"testing"
)

func TestExponentialRetryFail(t *testing.T) {
	SetExpRetrySleepDuration("1")
	SetExpRetryMaxSleepDuration("10")
	valInterface, err := ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
	if err == nil {
		t.Error("Failure test returning no error")
	}

	if valInterface != nil {
		val := valInterface.(int)
		t.Errorf("Failure test ended up succesful returned %d", val)
	}
}

func TestExponentialRetrySucess(t *testing.T) {
	testInteger := 10
	valInterface, err := ExponentialRetry(func() (interface{}, error) {
		return testInteger, nil
	})
	if err != nil {
		t.Errorf("Success test returning erro")
	}
	val := valInterface.(int)
	if val != testInteger {
		t.Errorf("Succes test failing. want %d got %d", 10, val)
	}
}

func TestExponentialRetryPanicSleepDuration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Test: ", r)
		}
	}()
	SetExpRetrySleepDuration("A")
	ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
}

func TestExponentialRetryMaxSleepDuration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Test: ", r)
		}
	}()
	SetExpRetrySleepDuration("1")
	SetExpRetryMaxSleepDuration("A")
	ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
}
