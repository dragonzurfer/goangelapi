package goangelapi_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dragonzurfer/goangelapi"
)

func TestExponentialRetryFail(t *testing.T) {
	goangelapi.SetExpRetrySleepDuration("1")
	goangelapi.SetExpRetryMaxSleepDuration("10")
	valInterface, err := goangelapi.ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
	if err == nil {
		t.Error("Failure test returning no error")
	}

	if valInterface != nil {
		val := valInterface.(int)
		t.Errorf("Failure test ended up succesful returned %d", val)
	}
	goangelapi.SetExpRetrySleepDuration("5")
	goangelapi.SetExpRetryMaxSleepDuration("30")
}

func TestExponentialRetrySucess(t *testing.T) {
	testInteger := 10
	valInterface, err := goangelapi.ExponentialRetry(func() (interface{}, error) {
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
	goangelapi.SetExpRetrySleepDuration("A")
	goangelapi.ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
	goangelapi.SetExpRetrySleepDuration("5")
	goangelapi.SetExpRetryMaxSleepDuration("30")
}

func TestExponentialRetryMaxSleepDuration(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Test: ", r)
		}
	}()
	goangelapi.SetExpRetrySleepDuration("1")
	goangelapi.SetExpRetryMaxSleepDuration("5")
	goangelapi.ExponentialRetry(func() (interface{}, error) {
		return nil, errors.New("This is a test generated error")
	})
	goangelapi.SetExpRetrySleepDuration("5")
	goangelapi.SetExpRetryMaxSleepDuration("30")
}
