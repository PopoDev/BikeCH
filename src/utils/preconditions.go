package utils

import "fmt"

func CheckArgument(shouldBeTrue bool, msg string) {
	if !shouldBeTrue {
		panic(fmt.Sprintf("IllegalArgumentException: %s", msg))
	}
}
