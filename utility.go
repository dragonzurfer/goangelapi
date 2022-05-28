package goangelapi

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

var (
	//really need to fix the environment name
	FUNCTION_RETRY_SLEEP_DURATION     = "5"
	FUNCTION_RETRY_MAX_SLEEP_DURATION = "300"
)

func SetExpRetrySleepDuration(val string) {
	FUNCTION_RETRY_SLEEP_DURATION = val
}

func SetExpRetryMaxSleepDuration(val string) {
	FUNCTION_RETRY_MAX_SLEEP_DURATION = val
}

func ExponentialRetry(f func() (interface{}, error)) (interface{}, error) {
	sleepTime, err := strconv.Atoi(FUNCTION_RETRY_SLEEP_DURATION)
	if err != nil {
		panic("env variable FUNCTION_RETRY_SLEEP_DURATION is invalid: got " + FUNCTION_RETRY_SLEEP_DURATION)
	}

	maxSleepDuration, err := strconv.Atoi(FUNCTION_RETRY_MAX_SLEEP_DURATION)
	if err != nil {
		panic("env variable FUNCTION_RETRY_MAX_SLEEP_DURATION is invalid: got " + FUNCTION_RETRY_MAX_SLEEP_DURATION)
	}

	for {
		ret, err := f()
		if err == nil {
			return ret, err
		}

		if sleepTime > maxSleepDuration {
			return ret, err
		}

		time.Sleep(time.Second * time.Duration(sleepTime))
		fmt.Println("retrying ", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		sleepTime = sleepTime * 2
	}
}
