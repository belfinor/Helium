package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-01-22

var ruMap map[rune]bool
var enMap map[rune]bool

func init() {

	enMap = make(map[rune]bool)
	ruMap = make(map[rune]bool)

	for _, r := range "йцукенгшщзхъфывапролджэячсмитьбюё" {
		ruMap[r] = true
	}

	for _, r := range "qwertyuiopasdfghjklzxcvbnm" {
		enMap[r] = true
	}
}

func IsRussian(str string) bool {

	cntRu := 0
	cntEn := 0
	cntByUA := 0

	for _, s := range str {
		if _, h := ruMap[s]; h {
			cntRu++
		} else if _, h := enMap[s]; h {
			cntEn++
		}

		if s == 'і' || s == 'ў' {
			cntByUA++
		}
	}

	if cntByUA > 3 {
		return false
	}

	return cntRu >= cntEn
}
