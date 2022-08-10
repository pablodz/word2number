package main

import (
	"fmt"

	"github.com/pablodz/word2number/word2number"
)

func main() {
	number, err := word2number.Text2NumESNoOrder("Mi código es tres cuatro cinco seis, veinte. Treinta y uno. Noventa y cinco. Trescientos noventa y cuatro .")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("number", number)
}
