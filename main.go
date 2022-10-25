package main

import (
	"fmt"
	"time"

	"github.com/pablodz/word2number/word2number/lang"
)

func main() {
	tstart := time.Now()
	text := "Mi código es diecisiete mil setecientos treinta y dos perros y gatos, dos perros y gatos "
	textFixed, err := lang.Text2NumES(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("text:\t\t", text)
	fmt.Println("textFixed:\t", textFixed)
	fmt.Println("time:\t\t", time.Since(tstart))
}
