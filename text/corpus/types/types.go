package types

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-09

import (
	"bufio"
	"strings"
)

var fromCode map[uint16]string
var toCode map[string]uint16

var TP_CITY uint16
var TP_ANIMAL uint16
var TP_NAME uint16
var TP_SLANG uint16
var TP_PATRONYMIC uint16
var TP_FLY uint16
var TP_POLITIC uint16
var TP_BIRD uint16
var TP_FISH uint16
var TP_COUNTRY uint16
var TP_LASTNAME uint16
var TP_MAN uint16

func init() {

	fromCode = make(map[uint16]string, 128)
	toCode = make(map[string]uint16, 128)

	txt := `
	город
	животное
	имя
	мат
	отчество
	полет
	политик
	птица
	рыба
	страна
	фамилия
	человек
	`
	br := bufio.NewReader(strings.NewReader(txt))

	i := 0

	for {

		str, err := br.ReadString('\n')
		if err != nil {
			break
		}

		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}

		fromCode[uint16(i+1)] = str
		toCode[str] = uint16(i + 1)

		i++
	}

	TP_CITY = ToCode("город")
	TP_ANIMAL = ToCode("животное")
	TP_NAME = ToCode("имя")
	TP_SLANG = ToCode("мат")
	TP_PATRONYMIC = ToCode("отчество")
	TP_FLY = ToCode("полет")
	TP_POLITIC = ToCode("политик")
	TP_BIRD = ToCode("птица")
	TP_FISH = ToCode("рыба")
	TP_COUNTRY = ToCode("страна")
	TP_LASTNAME = ToCode("фамилия")
	TP_MAN = ToCode("человек")
}

func ToCode(str string) uint16 {
	if v, h := toCode[str]; h {
		return v
	}

	return 0
}

func FromCode(code uint16) string {
	if v, h := fromCode[code]; h {
		return v
	}

	return ""
}
