package word2number

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pablodz/word2number/utils"
)

var (
	// SlicEN without accents
	unitsEN_LVL0 = []string{
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

	// unitsEN_LVL4_1 = []string{
	// 	"cien",
	// 	"mil",
	// 	"millon",
	// 	"billon",
	// 	"trillon",
	// }

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
		Dictionary[v] = i * 1000
	}
	// LVL 4_1
	// for i, v := range unitsEN_LVL4_1 {
	// 	if i == 0 {
	// 		Dictionary[v] = int(math.Pow(10, float64(2)))
	// 	} else {
	// 		Dictionary[v] = int(math.Pow(10, float64(i*3)))
	// 	}
	// }
	// fmt.Printf("Dictionary: %+v\n", Dictionary)

	/* Algorithm */
	newText := []string{}
	text = strings.Replace(text, ".", " . ", -1)
	text = strings.Replace(text, ",", " , ", -1)
	for k, v := range mapENLVL3 {
		text = strings.Replace(text, k, v, -1)
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
		case utils.IsItemInSlice(wordLower, unitsEN_LVL0) || utils.IsItemInSlice(wordLower, unitsEN_LVL1):
			// fmt.Println("case 1")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
		case utils.IsItemInSlice(wordLower, unitsEN_LVL2_2) || // utils.IsItemInSlice(wordLower, unitsEN_LVL2_1) ||
			utils.IsItemInSlice(wordLower, unitsEN_LVL3_1): // utils.IsItemInSlice(wordLower, unitsEN_LVL4_1)
			// fmt.Println("case 2")
			value := Dictionary[wordLower]
			newText = append(newText, fmt.Sprint(value))
			newText = append(newText, "y")
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
	if newText[len(newText)-1] == "y" {
		newText = newText[:len(newText)-1]
	}
	// fmt.Println("newText: ", newText)
	return strings.Join(newText, " "), nil
}
