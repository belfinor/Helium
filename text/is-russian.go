package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-09

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

	for _, s := range str {
		if _, h := ruMap[s]; h {
			cntRu++
		} else if _, h := enMap[s]; h {
			cntEn++
		}
	}

	return cntRu >= cntEn
}
