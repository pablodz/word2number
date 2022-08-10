package main

import (
	"fmt"

	"github.com/pablodz/word2number/word2number"
)

func main() {
	text := "Mi código es cuatro mil novecientos noventa y nueve , y llueve cada cuatro mil novecientos noventa y ocho"
	textFixed, err := word2number.Text2NumES(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("text:\t\t", text)
	fmt.Println("textFixed:\t", textFixed)
}
