package main

import (
	"fmt"
	sec "packApp/secondary"
	third "packApp/third"
)
func main() {
	var mainVar = "this is man var"

	fmt.Println(mainVar)
	fmt.Println(sec.ExtVar)
	sec.ExternalFunction()
	fmt.Println(sec.AnotherExtVar)
	fmt.Println(sec.DiffExtVar)
	fmt.Println(third.ThirdExtVar)
}