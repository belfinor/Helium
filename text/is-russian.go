package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2019-06-17

import "unicode"

var ruMap map[rune]bool
var enMap map[rune]bool
var otherMap map[rune]bool

func init() {

	enMap = make(map[rune]bool)
	ruMap = make(map[rune]bool)
	otherMap = make(map[rune]bool)

	for _, r := range "йцукенгшщзхъфывапролджэячсмитьбюё" {
		ruMap[r] = true
	}

	for _, r := range "qwertyuiopasdfghjklzxcvbnm" {
		enMap[r] = true
	}

	for _, r := range "іўაბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰәҙөғҫҡңһѣ" {
		otherMap[unicode.ToLower(r)] = true
	}
}

func OtherCyr(str string) bool {
	cntByOther := 0

	for _, s := range str {
		if _, h := otherMap[s]; h {
			cntByOther++
		}
	}

	return cntByOther > 5
}

func IsRussian(str string) bool {

	cntRu := 0
	cntEn := 0
	cntByOther := 0
	cntSim := 0

	for _, s := range str {
		if _, h := ruMap[s]; h {
			cntRu++
			continue
		}

		if _, h := enMap[s]; h {
			cntEn++
			continue
		}

		if _, h := otherMap[s]; h {
			cntByOther++
		}

		if unicode.IsLetter(s) {
			cntSim++
		}
	}

	if cntByOther > 6 {
		return false
	}

	return cntRu >= cntEn && cntRu > cntSim*2
}
