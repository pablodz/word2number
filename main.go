package main

import (
	"fmt"

	"github.com/pablodz/word2number/word2number/lang"
)

func main() {
	text := "cuatro guion diecisiete "
	textFixed, err := lang.Text2NumES(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("text:\t\t", text)
	fmt.Println("textFixed:\t", textFixed)
}
