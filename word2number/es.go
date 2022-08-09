package word2number

func Text2NumES(value string) (int, error) {
	unitsLVL0 := []string{
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
	unitsLVL1 := []string{
		"veinte",
		"treinta",
		"cuarenta",
		"cincuenta",
		"sesenta",
		"setenta",
		"ochenta",
		"noventa",
	}
	// unitsLVL2 := []string{
	// 	"ciento",
	// 	"mil",
	// 	"millón",
	// }
	// LVL 0
	DictionaryLVL0 := map[string]int{}
	DictionaryLVL1 := map[string]int{}
	DictionaryLVL2 := map[string]int{}
	for i, v := range unitsLVL0 {
		DictionaryLVL0[v] = i
	}
	// LVL 1
	for i, v := range unitsLVL1 {
		DictionaryLVL1[v] = (i + 2) + 10 // started from 20
	}
	// LVL 2
	DictionaryLVL2["ciento"] = 100
	DictionaryLVL2["mil"] = 1000
	DictionaryLVL2["millón"] = 1000000

	/* Algorithm */

	return 0, nil
}
