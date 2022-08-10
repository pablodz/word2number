package word2number

import "errors"

func Word2Num(text string, langCode string) (string, error) {
	var shortLangCode string
	if len(langCode) > 2 {
		shortLangCode = langCode[:2]
	}
	switch shortLangCode {
	case "es":
		return Text2NumES(text)
	case "en":
		return text, errors.New("language not supported")
	default:
		return text, errors.New("language not supported")
	}
}
