package main

func main() {
	if err := returnCustomError(0); err != nil {
		println(err.Error())
	}

	if err := returnCustomError(-1); err != nil {
		println(err.Error())
	}

	if err := returnCustomError(1); err != nil {
		println(err.Error())
	}
}

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func returnCustomError(result int) error {
	if result == 0 {
		return &CustomError{Code: 404, Message: "Not Found"}
	}

	if result == -1 {
		return &CustomError{Code: 500, Message: "Internal Server Error"}
	}

	return &CustomError{Code: 400, Message: "Bad Request"}
}
