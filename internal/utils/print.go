package utils

import "fmt"

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
	colorBlue  = "\033[34m"
)

func PrintWarning(message string) {
	fmt.Println(colorRed + message + colorReset)
}

func PrintSuccess(message string) {
	fmt.Println(colorGreen + message + colorReset)
}

func PrintTitle(message string) {
	fmt.Println(colorBlue + message + colorReset)
}
