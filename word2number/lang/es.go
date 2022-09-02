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
	unitsES_LVL0 = []string{
		"cero",
		"uno",
		"dos",
		"tres",
		"cuatro",
		"cinco",
		"seis",
		"siete",
		"ocho",
		"nueve",
		"diez",
		"once",
		"doce",
		"trece",
		"catorce",
		"quince",
		"dieciseis",
		"diecisiete",
		"dieciocho",
		"diecinueve",
		"veinte",
		"veintiuno",
		"veintidos",
		"veintitres",
		"veinticuatro",
		"veinticinco",
		"veintiseis",
		"veintisiete",
		"veintiocho",
		"veintinueve",
	}
	unitsES_LVL1 = []string{
		"",
		"",
		"",
		"treinta",
		"cuarenta",
		"cincuenta",
		"sesenta",
		"setenta",
		"ochenta",
		"noventa",
	}

	unitsES_LVL2_1 = []string{
		"",
		"cien",
		"doscientos",
		"trescientos",
		"cuatrocientos",
		"quinientos",
		"seiscientos",
		"setecientos",
		"ochocientos",
		"novecientos",
	}

	unitsES_LVL2_2 = []string{
		"",
		"ciento",
		"doscientos",
		"trescientos",
		"cuatrocientos",
		"quinientos",
		"seiscientos",
		"setecientos",
		"ochocientos",
		"novecientos",
	}

	mapESLVL3 = map[string]string{
		"dos mil":    "dosmil",
		"tres mil":   "tresmil",
		"cuatro mil": "cuatromil",
		"cinco mil":  "cincomil",
		"seis mil":   "seismil",
		"siete mil":  "setemil",
		"ocho mil":   "ochomil",
		"nueve mil":  "nuevemil",
	}

	unitsES_LVL3_1 = []string{
		"",
		"",
		"dosmil",
		"tresmil",
		"cuatromil",
		"cincomil",
		"seismil",
		"setemil",
		"ochomil",
		"nuevemil",
	}

	unitsES_LVL4_0 = []string{
		"cien",
		"mil",
		"millon",
		"billon",
		"trillon",
	}

	conectorsES      = "y"
	hyphenES         = "guion"
	hyphenSymbolES   = "-"
	phoneES          = "mas"
	phoneSymbolES    = "+"
	negativeES       = "menos"
	negativeSymbolES = "-"
)

func Text2NumES(text string) (string, error) {

	// LVL 0
	Dictionary := map[string]int{}
	for i, v := range unitsES_LVL0 {
		Dictionary[v] = i
	}
	// LVL 1
	for i, v := range unitsES_LVL1 {
		Dictionary[v] = i * 10
	}
	// LVL 2_1
	for i, v := range unitsES_LVL2_1 {
		Dictionary[v] = i * 100
	}
	// LVL 2_2
	for i, v := range unitsES_LVL2_2 {
		Dictionary[v] = i * 100
	}
	// LVL 3_1
	for i, v := range unitsES_LVL3_1 {
		Dictionary[v] = i * 1000
	}
	// LVL 4_1
	for i, v := range unitsES_LVL4_0 {
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
		case utils.IsItemInSlice(wordLower, unitsES_LVL0) || utils.IsItemInSlice(wordLower, unitsES_LVL1):
			// fmt.Println("case 1")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
		case utils.IsItemInSlice(wordLower, unitsES_LVL2_1) || utils.IsItemInSlice(wordLower, unitsES_LVL2_2) ||
			utils.IsItemInSlice(wordLower, unitsES_LVL3_1) ||
			utils.IsItemInSlice(wordLower, unitsES_LVL4_0):
			// fmt.Println("case 2")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
			newText = append(newText, conectorsES)
		default:
			// fmt.Println("case default")
			newText = append(newText, word)
		}

	}
	// fmt.Println("newText: ", newText)

	for i := 0; i < len(newText)-1; i++ {
		// fmt.Println("newText[", i, "]: ", v)
		v := newText[i]

		if v == conectorsES && i == 0 {
			continue
		}

		if v == conectorsES {
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
		if v == conectorsES {
			i = i - 1
		}
	}
	// remove y from the end of the text if it exists
	if newText[len(newText)-1] == conectorsES {
		newText = newText[:len(newText)-1]
	}

	// Hyphens between numbers and phone symbol
	for i := 0; i < len(newText)-1; i++ {
		v := newText[i]
		if i != 0 {
			if v == hyphenES && utils.DigitCheck.MatchString(newText[i-1]) && utils.DigitCheck.MatchString(newText[i+1]) {
				newText = append(newText[:i], newText[i+1:]...)
				newText = utils.InsertValueByIndexInSlice(newText, i, hyphenSymbolES)
				i = i - 1
			}
		}
		if utils.RemoveAccentMarks(v) == negativeES && utils.DigitCheck.MatchString(newText[i+1]) {
			fmt.Println(v)
			newText = append(newText[:i], newText[i+1:]...)
			newText = utils.InsertValueByIndexInSlice(newText, i, negativeSymbolES)
		}
		if utils.RemoveAccentMarks(v) == phoneES && utils.DigitCheck.MatchString(newText[i+1]) {
			fmt.Println(v)
			newText = append(newText[:i], newText[i+1:]...)
			newText = utils.InsertValueByIndexInSlice(newText, i, phoneSymbolES)
		}
	}
	// Join that contains numbers without space
	finalText := ""
	for i := 0; i < len(newText); i++ {
		if i == len(newText)-1 {
			finalText = finalText + newText[i]
			continue
		}
		switch {
		case utils.DigitCheck.MatchString(newText[i]) && utils.DigitCheck.MatchString(newText[i+1]):
			finalText = finalText + newText[i]
		case newText[i] == negativeSymbolES && utils.DigitCheck.MatchString(newText[i+1]):
			finalText = finalText + newText[i]
		case newText[i] == phoneSymbolES && utils.DigitCheck.MatchString(newText[i+1]):
			finalText = finalText + newText[i]
		default:
			finalText = finalText + newText[i] + " "
		}
	}

	// return strings.Join(newText, " "), nil
	return finalText, nil
}
