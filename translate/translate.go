package main

import "strings"

type letter struct {
	translations map[string]string
}

func main() {

}

func (l *letter) addWordTranslation(word string, translationWord string) {
	if l.isEmpty() {
		l.translations = make(map[string]string)
		l.translations[word] = translationWord
	} else {
		if l.existTranslation(word) {
			translationWord = (l.translationWord(word) + " " + translationWord)
		}
		l.translations[word] = translationWord
	}

}

func (l letter) isEmpty() bool {
	if l.translations == nil {
		return true
	}
	return false
}

func (l letter) existTranslation(word string) bool {
	if l.translations[word] != "" {
		return true
	}
	return false
}

func (l *letter) translationWord(word string) string {
	return l.translations[word]
}

func (l *letter) translationPhrase(phrase string) string {
	slices := strings.Split(phrase, " ")
	var translatedWord string = ""

	for _, slice := range slices {
		translatedWord += " " + l.translationWord(slice)
	}
	translatedWord = translatedWord[1:]
	return translatedWord
}
