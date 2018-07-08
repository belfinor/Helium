package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.005
// @date    2018-07-08

import (
	"strings"
	"unicode"
)

func GetWords(text string) []string {
	text = strings.Replace(strings.ToLower(text), "ё", "е", -1)

	has := false
	str := ""
	list := make([]string, 0, 10000)

	allow := map[string]bool{"-": true, ".": true, "+": true}
	clean := map[string]bool{"-": true, ".": true}

	for _, run := range text {

		c := string(run)
		_, h := allow[c]

		if h || unicode.IsLetter(run) || unicode.IsDigit(run) {
			has = true
			str += c
		} else {
			has = false

			str = strings.TrimFunc(str, func(r rune) bool {
				_, h := clean[string(r)]
				return h
			})

			if str != "" {
				list = append(list, str)
			}
			str = ""
		}

	}

	if has {
		str = strings.TrimFunc(str, func(r rune) bool {
			_, h := clean[string(r)]
			return h
		})

		if str != "" {
			list = append(list, str)
		}
	}

	return list
}

func Truncate(text string, limit int) string {
	result := ""
	i := 0
	for _, rune := range text {
		i++
		if i >= limit {
			result += "..."
			break
		}
		result += string(rune)
	}
	return result
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
