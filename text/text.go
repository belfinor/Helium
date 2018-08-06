package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.007
// @date    2018-08-06

import (
	"strings"
)

func GetWords(text string) []string {

	list := make([]string, 0, 100)

	stream := WordStream(strings.NewReader(text))

	for v := range stream {
		list = append(list, v)
	}

	return list
}

func Truncate(text string, limit int) string {

	builder := strings.Builder{}
	i := 0

	for _, rune := range text {
		builder.WriteRune(rune)
		i++

		if i >= limit {
			builder.WriteString("...")
			break
		}
	}
	return builder.String()
}

func GetPrefixes(str string) []string {
	list := make([]string, 0, 16)
	prefix := ""

	for _, c := range str {
		prefix = prefix + string(c)
		list = append(list, prefix)
	}

	return list
}

func MakePrefSuf(str string) [][]string {
	list := [][]string{}
	prefix := ""

	for _, c := range str {
		rune := string(c)
		for _, val := range list {
			val[1] = val[1] + rune
		}
		list = append(list, []string{prefix, rune})
		prefix = prefix + rune
	}

	list = append(list, []string{prefix, ""})

	return list
}

func SubStrDistVal(str string, minlen int) map[string]int {

	runes := make([]string, 0, 20)

	for _, r := range str {
		runes = append(runes, string(r))
	}

	if minlen < 3 {
		minlen = 3
	}

	size := len(runes)
	res := make(map[string]int)

	if size < minlen {
		return res
	}

	for {
		clen := len(runes)
		data := runes

		if clen < minlen {
			break
		}

		for clen >= minlen {
			cstr := strings.Join(data, "")
			res[cstr] = size - clen
			data = data[:clen-1]
			clen--
		}

		runes = runes[1:]
	}

	return res
}
