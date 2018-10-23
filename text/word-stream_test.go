package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-10-23

import (
	"strings"
	"testing"
)

func TestWordStream(t *testing.T) {

	tests := [][]string{
		[]string{"привет мир", "привет", "мир"},
		[]string{"привет, мир", "привет", "мир"},
		[]string{"Корпорация Mail.ru предложила амнистировать россиян", "корпорация", "mail.ru", "предложила", "амнистировать", "россиян"},
		[]string{"группа Ёрш.", "группа", "ерш"},
		[]string{"ставь #хэштег#хештег тег #мояцлица друг", "ставь", "тег", "друг"},
		[]string{"2*2=4 это знают в целом мире, да?", "2", "2", "4", "это", "знают", "в", "целом", "мире", "да"},
		[]string{"Сайт https://yandex.ru - отличный поиск", "сайт", "https://yandex.ru", "отличный", "поиск"},
		[]string{"Зайди на http://mail.ru/?", "зайди", "на", "http://mail.ru/"},
	}

	tests_ends := [][]string{
		[]string{"привет мир 12", "привет", "мир", "12", "."},
		[]string{"привет, мир", "привет", "мир", "."},
		[]string{"Корпорация Mail.ru https://mail.ru/ предложила амнистировать россиян", "корпорация", "mail.ru", "предложила", "амнистировать", "россиян", "."},
		[]string{"группа Ёрш.", "группа", "ерш", "."},
		[]string{"ставь #хэштег#хештег тег? #мояцлица друг https://yandex.ru", "ставь", "тег", ".", "друг", "."},
		[]string{"2*2=4 это знают в целом мире, да https://yandex.ru ?", "2", "2", "4", "это", "знают", "в", "целом", "мире", "да", "."},
	}

	test_nonums := [][]string{
		[]string{"привет 12 мир", "привет", "мир", "."},
		[]string{"это случилось 28.05.1985", "это", "случилось", "."},
		[]string{"стук в дверь в 12:00 заставил", "стук", "в", "дверь", "в", "заставил", "."},
		[]string{"132*2.1478=4.0 это знают в целом мире, да https://yandex.ru ?", "это", "знают", "в", "целом", "мире", "да", "."},
	}

	test_nonums_hashtags := [][]string{
		[]string{"привет 12 мир", "привет", "мир", "."},
		[]string{"это случилось 28.05.1985 #test#1#тест #0 xxx", "это", "случилось", "test", "тест", "xxx", "."},
		[]string{"стук в дверь в 12:00 #привет заставил xxx:", "стук", "в", "дверь", "в", "привет", "заставил", "."},
		[]string{"132*2.1478=4.0 это знают в целом мире, ххх: да https://yandex.ru ? #тест", "это", "знают", "в", "целом", "мире", "да", "тест", "."},
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

	for _, v := range tests_ends {

		src := v[0]
		v = v[1:]
		stream := WordStream(strings.NewReader(src), WSO_ENDS, WSO_NO_URLS)

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

	for _, v := range test_nonums {

		src := v[0]
		v = v[1:]
		stream := WordStream(strings.NewReader(src), WSO_ENDS, WSO_NO_URLS, WSO_NO_NUMS)

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

	for _, v := range test_nonums_hashtags {

		src := v[0]
		v = v[1:]
		stream := WordStream(strings.NewReader(src), WSO_ENDS, WSO_NO_URLS, WSO_NO_NUMS, WSO_HASHTAG, WSO_NO_XXX_COLON)

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
