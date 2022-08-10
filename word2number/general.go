package word2number

import "errors"

func Word2Num(text string, langCode string) (string, error) {

	switch langCode {
	case "en":
		return "", errors.New("language not supported")
	case "es":
		return Text2NumES(text)
	}
	return "", errors.New("language not supported")
}
