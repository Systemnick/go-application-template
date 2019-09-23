package main

import (
	"fmt"
)

func main() {
	a, err := NewApplication()
	if err != nil {
		fmt.Printf("Creating application error: %s\n", err.Error())
		return
	}

	fmt.Printf("Application: %+v\n", a)

	a.Run()
}
