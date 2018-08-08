package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-08-08

import (
	"testing"
)

func TestTextGetWords(t *testing.T) {
	text := "hello World! Сегодня хороший день для 1-ого теста. Ёлка в лесу"

	res := GetWords(text)

	wait := []string{"hello", "world", "сегодня", "хороший", "день", "для", "1-ого", "теста", "елка", "в", "лесу"}

	if len(res) != len(wait) {
		t.Fatal("bad result list size")
	}

	for i, val := range res {
		if val != wait[i] {
			t.Fatal("invalid item in slice")
		}
	}
}

func TestGetPrefixes(t *testing.T) {
	wait := []string{"т", "те", "тес", "тест"}
	res := GetPrefixes("тест")
	if len(wait) != len(res) {
		t.Fatal("invalid result length")
	}
	for i, v := range res {
		if wait[i] != v {
			t.Fatal("invalid value: " + wait[i])
		}
	}
}

func TestTextMakePrefSuf(t *testing.T) {
	wait := [][]string{
		[]string{"", "тест"},
		[]string{"т", "ест"},
		[]string{"те", "ст"},
		[]string{"тес", "т"},
		[]string{"тест", ""},
	}

	res := MakePrefSuf("тест")

	if len(res) != len(wait) {
		t.Fatal("Invalid result length")
	}

	for i, rec := range res {
		if wait[i][0] != rec[0] || wait[i][1] != rec[1] {
			t.Fatal("Invalid result")
		}
	}
}

func TestTruncate(t *testing.T) {

	if Truncate("привет, мир!", 9) != "привет, м…" {
		t.Fatal("Truncate not working")
	}

	if Truncate("привет, мир!", 11) != "привет, мир…" {
		t.Fatal("Truncate not working")
	}

}
