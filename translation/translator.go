package translation

import "strings"

func Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return ""
	}
	switch strings.ToLower(language) {
	case "english":
		return word
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
