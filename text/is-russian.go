package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2019-01-29

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

	for _, r := range "іўაბგდევზთიკლმნოპჟრსტუფქღყშჩცძწჭხჯჰ" {
		otherMap[r] = true
	}
}

func IsRussian(str string) bool {

	cntRu := 0
	cntEn := 0
	cntByOther := 0

	for _, s := range str {
		if _, h := ruMap[s]; h {
			cntRu++
		} else if _, h := enMap[s]; h {
			cntEn++
		}

		if _, h := otherMap[s]; h {
			cntByOther++
		}
	}

	if cntByOther > 5 {
		return false
	}

	return cntRu >= cntEn
}
