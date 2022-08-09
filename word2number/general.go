package word2number

import "errors"

func Word2Num(value string, langCode string) (int, error) {

	if langCode == "en" {
		return Word2NumEN(value)
	} else if langCode == "es" {
		return Text2NumES(value)
	}
	return 0, errors.New("language not supported")
}
