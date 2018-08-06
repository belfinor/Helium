package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-08-06

import (
	"strings"
	"testing"
)

func TestWordStream(t *testing.T) {

	tests := [][]string{
		[]string{"привет мир", "привет", "мир"},
		[]string{"привет, мир", "привет", "мир"},
		[]string{"Корпорация Mail.ru предложила амнистировать россиян", "корпорация", "mail.ru", "предложила", "амнистировать", "россиян"},
		[]string{"группа Ёрш", "группа", "ерш"},
		[]string{"ставь #хэштег#хештег тег #мояцлица друг", "ставь", "тег", "друг"},
		[]string{"2*2=4 это знают в целом мире, да?", "2", "2", "4", "это", "знают", "в", "целом", "мире", "да"},
	}

	for _, v := range tests {

		src := v[0]
		v = v[1:]
		stream := WordStream(strings.NewReader(src))

		for wrd := range stream {
			if len(v) == 0 {
				t.Fatal(src)
			}

			if wrd != v[0] {
				t.Fatal(src)
			}

			if len(v) > 1 {
				v = v[1:]
			} else {
				v = []string{}
			}
		}

		if len(v) != 0 {
			t.Fatal(src)
		}
	}
}
