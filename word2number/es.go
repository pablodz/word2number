package word2number

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pablodz/word2number/utils"
)

var (
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
		"dieciséis",
		"diecisiete",
		"dieciocho",
		"diecinueve",
	}
	unitsES_LVL1 = []string{
		"",
		"",
		"veinte",
		"treinta",
		"cuarenta",
		"cincuenta",
		"sesenta",
		"setenta",
		"ochenta",
		"noventa",
	}

	unitsES_LVL2 = []string{
		"", // cien
		"", // doscientos
		"", // trescientos
		"", // cuatrocientos
		"quinientos",
		"", // seiscientos
		"", // setecientos
		"", // ochocientos
		"", // novecientos
	}

	unitsES_LVL3 = []string{
		"ciento",
		"mil",
		"millón",
		"billón",
		"trillón",
	}
	conectorsES = "y"
)

func Text2NumESNoOrder(text string) (string, error) {

	// LVL 0
	Dictionary := map[string]int{}
	for i, v := range unitsES_LVL0 {
		Dictionary[v] = i
	}
	// LVL 1
	for i, v := range unitsES_LVL1 {
		Dictionary[v] = i * 10
	}
	// LVL 2
	for i, v := range unitsES_LVL3 {
		if i == 0 {
			Dictionary[v] = int(math.Pow(10, float64(2)))
		} else {
			Dictionary[v] = int(math.Pow(10, float64(i*3)))
		}
	}
	fmt.Printf("Dictionary: %+v\n", Dictionary)

	/* Algorithm */
	newText := []string{}
	text = strings.Replace(text, ".", " . ", -1)
	text = strings.Replace(text, ",", " , ", -1)
	textSplitted := strings.Split(text, " ")
	for _, word := range textSplitted {
		if word == "" {
			continue
		}
		wordLower := strings.ToLower(word)
		// fmt.Println("word: <" + word + ">")
		// First iteration
		switch {
		case wordLower == conectorsES:
			newText = append(newText, wordLower)
		case utils.IsItemInSlice(wordLower, unitsES_LVL0) || utils.IsItemInSlice(wordLower, unitsES_LVL1) || utils.IsItemInSlice(wordLower, unitsES_LVL3):
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
		case strings.Contains(wordLower, "y"):

		default:
			newText = append(newText, word)
		}

	}
	fmt.Println("newText: ", newText)

	for i, v := range newText {
		fmt.Println("newText[", i, "]: ", v)
		if v == conectorsES {
			x := newText[i-1]
			y := newText[i+1]
			xNum, _ := strconv.Atoi(x)
			yNum, _ := strconv.Atoi(y)
			zNum := xNum + yNum
			// fmt.Println("x: ", x, "y: ", y, "z: ", zNum)
			newText = append(newText[:i-1], newText[i+2:]...)
			newText = utils.InsertValueByIndexInSlice(newText, i-1, fmt.Sprint(zNum))
		}
	}
	fmt.Println("newText: ", newText)
	return strings.Join(newText, " "), nil
}

// func Text2NumESNoOrder(text string) (int, error) {

// 	// LVL 0
// 	DictionaryLVL0 := map[string]int{}
// 	DictionaryLVL1 := map[string]int{}
// 	DictionaryLVL2 := map[string]int{}
// 	for i, v := range unitsES_LVL0 {
// 		DictionaryLVL0[v] = i
// 	}
// 	// LVL 1
// 	for i, v := range unitsES_LVL1 {
// 		DictionaryLVL1[v] = (i + 2) * 10 // started from 20
// 	}
// 	// LVL 2
// 	DictionaryLVL2["ciento"] = 100
// 	DictionaryLVL2["mil"] = 1000
// 	DictionaryLVL2["millón"] = 1000000
// 	DictionaryLVL2["millones"] = 1000000

// 	// print dicts with fields and values
// 	fmt.Printf("DictionaryLVL0: %+v\n", DictionaryLVL0)
// 	fmt.Printf("DictionaryLVL1: %+v\n", DictionaryLVL1)
// 	fmt.Printf("DictionaryLVL2: %+v\n", DictionaryLVL2)

// 	/* Algorithm */
// 	tempNumber := 0
// 	finalNumber := 0
// 	decimalFlag := false
// 	numberFlag := false
// 	finalString := ""

// 	textSplitted := strings.Split(text, " ")
// 	fmt.Println("textSplitted:", textSplitted)
// 	for i, word := range textSplitted {
// 		fmt.Println("word", word)

// 		if utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL0) ||
// 			utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL1) ||
// 			utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL2) &&
// 				!decimalFlag {
// 			fmt.Println("1º condition")
// 			numberFlag = true
// 			if utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL0) {
// 				fmt.Println("1.1º condition", word)
// 				tempNumber += DictionaryLVL0[strings.ToLower(word)]
// 			} else if utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL1) {
// 				fmt.Println("1.2º condition")
// 				tempNumber += DictionaryLVL1[strings.ToLower(word)]
// 			} else if strings.ToLower(word) == unitsES_LVL2[0] { // "ciento"
// 				fmt.Println("1.3º condition")
// 				tempNumber *= 100
// 			} else if utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL1[1:]) { // "mil", "millón"
// 				fmt.Println("1.4º condition")
// 				finalNumber += tempNumber + tempNumber*DictionaryLVL2[strings.ToLower(word)]
// 				tempNumber = 0
// 			}

// 		} else if utils.IsItemInSlice(strings.ToLower(word), unitsES_LVL0) &&
// 			decimalFlag {
// 			fmt.Println("2º condition")
// 			finalString += fmt.Sprint(DictionaryLVL0[strings.ToLower(word)])
// 		} else if word == conectorsES &&
// 			(utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL0) ||
// 				utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL1) ||
// 				utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL2)) {
// 			fmt.Println("3º condition")
// 			continue
// 		} else if word == "punto" &&
// 			(utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL0) ||
// 				utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL1) ||
// 				utils.IsItemInSlice(strings.ToLower(textSplitted[i+1]), unitsES_LVL2)) {
// 			fmt.Println("4º condition")
// 			finalNumber += tempNumber
// 			decimalFlag = true
// 			numberFlag = false
// 			finalString += " " + fmt.Sprint(finalNumber) + "."
// 			continue
// 		} else {
// 			fmt.Println("5º condition")
// 			decimalFlag = false
// 			if numberFlag {
// 				finalNumber += tempNumber
// 				if word != "." && word != "?" && word != "!" {
// 					finalString += " " + fmt.Sprint(finalNumber) + " " + word
// 				} else {
// 					finalString += " " + fmt.Sprint(finalNumber) + word
// 				}
// 				numberFlag = false
// 			} else {
// 				if word != "." && word != "?" && word != "!" {
// 					finalString += " " + word
// 				} else {
// 					finalString += word
// 				}
// 			}
// 			finalNumber = 0
// 			tempNumber = 0
// 		}
// 		fmt.Println("finalString", finalString)
// 	}
// 	fmt.Println("finalString", finalString)
// 	return finalNumber, nil
// }
