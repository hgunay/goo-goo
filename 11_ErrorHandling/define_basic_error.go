package main

import "errors"

func main() {
	err := returnError()
	if err != nil {
		println(err.Error())
	}
}

func returnError() error {
	return errors.New("error test message")
}
