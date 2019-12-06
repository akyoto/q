package build

import (
	"fmt"
	"strings"
)

// This flag controls whether parametric polymorphism is enabled or not.
const PolymorphismEnabled = false

// PolymorphName attaches parameter-specific information to the function name.
func PolymorphName(functionName string, parameterCount int) string {
	if !PolymorphismEnabled {
		return functionName
	}

	if parameterCount == 0 {
		return functionName
	}

	if functionName == "syscall" || functionName == "print" {
		return functionName
	}

	return fmt.Sprintf("%s|%d", functionName, parameterCount)
}

// UnpolymorphName removes parameter-specific information from the function name.
func UnpolymorphName(functionName string) string {
	if !PolymorphismEnabled {
		return functionName
	}

	index := strings.LastIndex(functionName, "|")

	if index == -1 {
		return functionName
	}

	return functionName[:index]
}
