package main

import "fmt"

func random() any {
	return 200
}

func main() {
	var result any = random()
	//var resultString string = result.(string)
	//fmt.Println(resultString)
	//
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknow Type", value)
	}

}
