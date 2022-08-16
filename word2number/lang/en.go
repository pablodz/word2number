package lang

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pablodz/word2number/utils"
)

var (
	// SlicEN without accents
	unitsEN_LVL0 = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
	unitsEN_LVL1 = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	// unitsEN_LVL2_1 = []string{
	// 	"",
	// 	"hundred",
	// 	"doscientos",
	// 	"trENcientos",
	// 	"cuatrocientos",
	// 	"quinientos",
	// 	"seiscientos",
	// 	"setecientos",
	// 	"ochocientos",
	// 	"novecientos",
	// }

	unitsEN_LVL2_2 = []string{
		"",
		"onehundred",
		"twohundred",
		"threehundred",
		"fourhundred",
		"fivehundred",
		"sixhundred",
		"sevenhundred",
		"eighhundred",
		"ninehundred",
	}

	mapENLVL3 = map[string]string{
		"one hundred":   "onehundred",
		"two hundred":   "twohundred",
		"three hundred": "threehundred",
		"four hundred":  "fourhundred",
		"five hundred":  "fivehundred",
		"six hundred":   "sixhundred",
		"seven hundred": "sevenhundred",
		"eight hundred": "eighhundred",
		"nine hundred":  "ninehundred",
	}

	unitsEN_LVL3_1 = []string{
		"",
		"onehundred",
		"twohundred",
		"threehundred",
		"fourhundred",
		"fivehundred",
		"sixhundred",
		"sevenhundred",
		"eighhundred",
		"ninehundred",
	}

	mapENLVL4 = map[string]string{
		"one thousand":   "onethousand",
		"two thousand":   "twothousand",
		"three thousand": "threethousand",
		"four thousand":  "fourthousand",
		"five thousand":  "fivethousand",
		"six thousand":   "sixthousand",
		"seven thousand": "seventhousand",
		"eight thousand": "eighthousand",
		"nine thousand":  "ninethousand",
	}

	unitsEN_LVL4_1 = []string{
		"",
		"onethousand",
		"twothousand",
		"threethousand",
		"fourthousand",
		"fivethousand",
		"sixthousand",
		"seventhousand",
		"eighthousand",
		"ninethousand",
	}

	unitsEN_LVL4_0 = []string{
		"hundred",
		"thousand",
		"million",
		"billion",
		"trillon",
	}

	conectorsEN = "and"
)

func Text2NumEN(text string) (string, error) {

	// LVL 0
	Dictionary := map[string]int{}
	for i, v := range unitsEN_LVL0 {
		Dictionary[v] = i
	}
	// LVL 1
	for i, v := range unitsEN_LVL1 {
		Dictionary[v] = i * 10
	}
	// // LVL 2_1
	// for i, v := range unitsEN_LVL2_1 {
	// 	Dictionary[v] = i * 100
	// }
	// LVL 2_2
	for i, v := range unitsEN_LVL2_2 {
		Dictionary[v] = i * 100
	}
	// LVL 3_1
	for i, v := range unitsEN_LVL3_1 {
		Dictionary[v] = i * 100
	}
	// LVL 4_0
	for i, v := range unitsES_LVL4_0 {
		if i == 0 {
			Dictionary[v] = int(math.Pow(10, float64(2)))
		} else {
			Dictionary[v] = int(math.Pow(10, float64(i*3)))
		}
	}
	// LVL 4_1
	for i, v := range unitsEN_LVL4_1 {
		Dictionary[v] = i * 1000
	}
	// fmt.Printf("Dictionary: %+v\n", Dictionary)

	/* Algorithm */
	newText := []string{}
	text = strings.ReplaceAll(text, ".", " . ")
	text = strings.ReplaceAll(text, ",", " , ")
	for k, v := range mapENLVL3 {
		text = strings.ReplaceAll(text, k, v)
	}
	for k, v := range mapENLVL4 {
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
		case strings.Contains(wordLower, "-"):
			// fmt.Println("case 0")
			listnumbers := strings.Split(wordLower, "-")
			val1 := Dictionary[listnumbers[0]]
			val2 := Dictionary[listnumbers[1]]
			// fmt.Println("val1: <" + strconv.Itoa(val1) + ">")
			// fmt.Println("val2: <" + strconv.Itoa(val2) + ">")
			value := val1 + val2
			newText = append(newText, strconv.Itoa(value))
		case utils.IsItemInSlice(wordLower, unitsEN_LVL0) || utils.IsItemInSlice(wordLower, unitsEN_LVL1):
			// fmt.Println("case 1")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
		case utils.IsItemInSlice(wordLower, unitsEN_LVL2_2) || // utils.IsItemInSlice(wordLower, unitsEN_LVL2_1) ||
			utils.IsItemInSlice(wordLower, unitsEN_LVL3_1) ||
			utils.IsItemInSlice(wordLower, unitsEN_LVL4_0) ||
			utils.IsItemInSlice(wordLower, unitsEN_LVL4_1): //
			// fmt.Println("case 2", wordLower)
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
			newText = append(newText, conectorsEN)
		default:
			// fmt.Println("case default")
			newText = append(newText, word)
		}

	}
	// fmt.Println("newText: ", newText)

	for i := 0; i < len(newText)-1; i++ {
		// fmt.Println("newText[", i, "]: ", v)
		v := newText[i]

		if v == conectorsEN && i == 0 {
			continue
		}

		if v == conectorsEN {
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
		if v == conectorsEN {
			i = i - 1
		}
	}
	// remove y from the end of the text if it exists
	if newText[len(newText)-1] == conectorsEN {
		newText = newText[:len(newText)-1]
	}
	// fmt.Println("newText: ", newText)
	return strings.Join(newText, " "), nil
}
