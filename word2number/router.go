package word2number

import (
	"errors"

	"github.com/pablodz/word2number/word2number/lang"
)

func Word2Num(text string, langCode string) (string, error) {
	var shortLangCode string
	if len(langCode) > 2 {
		shortLangCode = langCode[:2]
	}
	switch shortLangCode {
	case "es":
		return lang.Text2NumES(text)
	case "en":
		return lang.Text2NumEN(text)
	case "pr":
		return lang.Text2NumPR(text)
	default:
		return text, errors.New("language not supported <" + langCode + ">")
	}
}
