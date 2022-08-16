package main

import (
	"fmt"

	"github.com/pablodz/word2number/word2number/lang"
)

func main() {
	text := "My code is cinquenta e nove"
	textFixed, err := lang.Text2NumPR(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("text:\t\t", text)
	fmt.Println("textFixed:\t", textFixed)
}
