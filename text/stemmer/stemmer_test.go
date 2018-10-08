package stemmer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-10-08

import (
	"testing"
)

func TestTextToCode(t *testing.T) {

	data := map[string]string{
		"":                  "",
		"игра":              "игр",
		"купить ботинок":    "куп ботинок",
		"убить время":       "уб врем",
		"майорка":           "майорк",
		"пилота-разведчика": "пилот-разведчик",
		"science":           "scienc",
	}

	for k, v := range data {
		if TextToCode(k) != v {
			t.Fatal("TextToCode wait " + k)
		}
	}

}
