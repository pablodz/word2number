package word2number

import "errors"

func Word2Num(text string, langCode string) (string, error) {

	if langCode == "en" {
		// return Word2NumEN(text)
	} else if langCode == "es" {
		return Text2NumESNoOrder(text)
	}
	return "", errors.New("language not supported")
}
