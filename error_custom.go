package main

import "fmt"

type validationError struct {
	Message string
}

func (e *validationError) Error() string {
	return e.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation error"}
	}

	if id != "Febry" {
		return &notFoundError{"Username tidak di temukan"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		//// maka terjadi error
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validationErr.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("User tidak di temukan:", notFoundError.Error())
		} else {
			fmt.Println("Unknow error:", err.Error())
		}

		//switch finalError := err.(type) {
		//case *validationError:
		//	fmt.Println("Validation error use switch", finalError.Error())
		//case *notFoundError:
		//	fmt.Println("User tidak di temukan menggunakan switch:", finalError.Error())
		//default:
		//	fmt.Println("Unknow error:", finalError.Error())
		//}
	} else {
		// jika sukses
		fmt.Println("Successfully save data")
	}
}
