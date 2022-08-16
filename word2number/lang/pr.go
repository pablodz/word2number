package lang

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pablodz/word2number/utils"
)

var (
	// Slices without accents
	unitsPR_LVL0 = []string{
		"zero",
		"um",
		"dois",
		"tres",
		"quatro",
		"cinco",
		"seis",
		"sete",
		"oito",
		"nove",
		"dez",
		"onze",
		"doze",
		"treze",
		"quatorze",
		"quinze",
		"dezesseis",
		"dezessete",
		"dezoito",
		"dezenove",
	}
	unitsPR_LVL1 = []string{
		"",
		"",
		"vinte",
		"trinta",
		"quarenta",
		"cinquenta",
		"sessenta",
		"setenta",
		"oitenta",
		"noventa",
	}

	unitsPR_LVL2_1 = []string{
		"",
		"cento",
		"duzentos",
		"trezentos",
		"quatrocentos",
		"quinhentos",
		"seiscentos",
		"setecentos",
		"oitocentos",
		"novecentos",
	}

	unitsPR_LVL2_2 = []string{
		"",
		"cem",
		"duzentos",
		"trezentos",
		"quatrocentos",
		"quinhentos",
		"seiscentos",
		"setecentos",
		"oitocentos",
		"novecentos",
	}

	unitsPR_LVL4_0 = []string{
		"cento",
		"mil",
		// "cien",
		// "mil",
		// "millon",
		// "billon",
		// "trillon",
	}

	conectorsPR = "e"
)

func Text2NumPR(text string) (string, error) {

	// LVL 0
	Dictionary := map[string]int{}
	for i, v := range unitsPR_LVL0 {
		Dictionary[v] = i
	}
	// LVL 1
	for i, v := range unitsPR_LVL1 {
		Dictionary[v] = i * 10
	}
	// LVL 2_1
	for i, v := range unitsPR_LVL2_1 {
		Dictionary[v] = i * 100
	}
	// LVL 2_2
	for i, v := range unitsPR_LVL2_2 {
		Dictionary[v] = i * 100
	}
	// // LVL 3_1
	// for i, v := range unitsPR_LVL3_1 {
	// 	Dictionary[v] = i * 1000
	// }
	// LVL 4_1
	for i, v := range unitsPR_LVL4_0 {
		if i == 0 {
			Dictionary[v] = int(math.Pow(10, float64(2)))
		} else {
			Dictionary[v] = int(math.Pow(10, float64(i*3)))
		}
	}
	// fmt.Printf("Dictionary: %+v\n", Dictionary)

	/* Algorithm */
	newText := []string{}
	text = strings.ReplaceAll(text, ".", " . ")
	text = strings.ReplaceAll(text, ",", " , ")
	for k, v := range mapESLVL3 {
		text = strings.ReplaceAll(text, k, v)
	}
	textSplitted := strings.Split(text, " ")
	for _, word := range textSplitted {
		if word == "" {
			continue
		}
		wordLower := strings.ToLower(word)
		// fmt.Println("wordLower: <" + wordLower + ">")
		// First iteration

		wordLower = utils.RemoveAccentMarks(wordLower)
		switch {
		case utils.IsItemInSlice(wordLower, unitsPR_LVL0) || utils.IsItemInSlice(wordLower, unitsPR_LVL1):
			// fmt.Println("case 1")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
		case utils.IsItemInSlice(wordLower, unitsPR_LVL2_1) ||
			utils.IsItemInSlice(wordLower, unitsPR_LVL2_2) ||
			utils.IsItemInSlice(wordLower, unitsPR_LVL4_0):
			// fmt.Println("case 2")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
			newText = append(newText, conectorsPR)
		default:
			// fmt.Println("case default")
			newText = append(newText, word)
		}

	}
	// fmt.Println("newText: ", newText)

	for i := 0; i < len(newText)-1; i++ {
		// fmt.Println("newText[", i, "]: ", v)
		v := newText[i]

		if v == conectorsPR && i == 0 {
			continue
		}

		if v == conectorsPR {
			x := newText[i-1]
			y := newText[i+1]
			xNum, err := strconv.Atoi(x)
			if err != nil {
				continue
			}
			yNum, err := strconv.Atoi(y)
			if err != nil {
				continue
			}
			zNum := xNum + yNum
			// fmt.Println("x: ", x, "y: ", y, "z: ", zNum)
			newText = append(newText[:i-1], newText[i+2:]...)
			newText = utils.InsertValueByIndexInSlice(newText, i-1, fmt.Sprint(zNum))
		}
		if v == conectorsPR {
			i = i - 1
		}
	}
	// remove y from the end of the text if it exists
	if newText[len(newText)-1] == conectorsPR {
		newText = newText[:len(newText)-1]
	}
	// fmt.Println("newText: ", newText)
	return strings.Join(newText, " "), nil
}
