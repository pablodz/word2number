package main

import (
	"fmt"

	"github.com/pablodz/word2number/word2number"
)

func main() {
	text := "My code is one thousand"
	textFixed, err := word2number.Text2NumEN(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("text:\t\t", text)
	fmt.Println("textFixed:\t", textFixed)
}
