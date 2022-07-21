package wordwrap

import (
	"strings"
)

func Wrap(input string, maxWidth int) string {
	if len(input) == 0 || maxWidth == 0 {
		return ""
	}

	var outputParas []string
	paragraphs := strings.Split(input, "\n\n")

	for _, paragraph := range paragraphs {
		wrapped := makeWrappedPara(strings.Fields(paragraph), maxWidth)
		outputParas = append(outputParas, wrapped)
	}

	return strings.Join(outputParas, "\n\n")
}

func makeWrappedPara(words []string, maxWidth int) string {
	wrapped := words[0]
	charsLeft := maxWidth - len(wrapped)

	for _, word := range words[1:] {
		if len(word)+1 > charsLeft {
			wrapped += "\n" + word
			charsLeft = maxWidth - len(word)
		} else {
			wrapped += " " + word
			charsLeft -= 1 + len(word)
		}
	}
	return wrapped
}
