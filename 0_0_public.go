package errMock

import (
	"errors"
	"fmt"
)

func PanicToErrMsg(callback func()) (errMsg string) {
	defer func() {
		errMsg = recoverPanic()
	}()
	callback()
	return
}

func PanicToErr(callback func()) (err error) {
	errMsg := PanicToErrMsg(callback)
	if errMsg == "" {
		return
	}
	err = errors.New(errMsg)
	return
}

func PanicToErrMsgAndStack(callback func()) (errMsg string, stack string) {
	defer func() {
		errMsg = recoverPanic()
		if errMsg != "" {
			stack = readStack()
		}
	}()
	callback()
	return
}

func PanicToErrAndLog(callback func()) (err error) {
	err = PanicToErr(callback)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func PanicToErrMsgAndLog(callback func()) (errMsg string) {
	errMsg = PanicToErrMsg(callback)
	if errMsg != "" {
		fmt.Println(errMsg)
	}
	return
}

func PanicToErrMsgStackAndLog(callback func()) (errMsg string, stack string) {
	errMsg, stack = PanicToErrMsgAndStack(callback)
	if errMsg != "" {
		fmt.Println("ErrMsg:", errMsg)
		fmt.Println("Stack:", stack)
	}
	return
}
