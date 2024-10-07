package exceptions

import "fmt"

type MyException struct {
	Data   int    `json:"data"`
	Object string `json:"object"`
}

func (error *MyException) Error() string {
	return fmt.Sprintf("Please input a value %v %v", error.Data, error.Object)
}
