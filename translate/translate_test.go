package main

import (
	"testing"
)

func TestTranslateSpeechless(t *testing.T) {
	l := letter{}
	if !l.isEmpty() {
		t.Errorf("value of translate: notEmpty")
	}
}

func TestTranslateOneTraduction(t *testing.T) {
	l := letter{}
	l.addWordTranslation("bom", "good")

	if l.isEmpty() {
		t.Errorf("value of translate: Empty")
	}
	if "good" != l.translationWord("bom") {
		t.Errorf("wrong translation, its a word translation: %s", l.translations)
	}

}

func TestTranslateTwoTraduction(t *testing.T) {
	l := letter{}
	l.addWordTranslation("bom", "good")
	l.addWordTranslation("mau", "bad")

	if "good" != l.translationWord("bom") {
		t.Errorf("false")
	}
	if "bad" != l.translationWord("mau") {
		t.Errorf("false")
	}
}

func TestTranslateSameWordTraduction(t *testing.T) {
	l := letter{}
	l.addWordTranslation("bom", "good")
	l.addWordTranslation("bom", "nice")

	translation := l.translationWord("bom")
	if "good" != translation[:len("good")] || "nice" != translation[len("good")+1:] {
		t.Errorf("false")
	}
}

func TestTranslatePhraseTraduction(t *testing.T) {
	l := letter{}
	l.addWordTranslation("guerra", "war")
	l.addWordTranslation("é", "is")
	l.addWordTranslation("ruim", "bad")

	translation := l.translationPhrase("guerra é ruim")

	if "war is bad" != translation {
		t.Errorf("false, traductio is:%s", translation)
	}
}
