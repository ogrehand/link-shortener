package controller

import "fmt"

func PanicRecovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f", r)
	}
}
