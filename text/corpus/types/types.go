package types

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-12-11

import (
	"bufio"
	"strings"
)

var fromCode map[uint16]string
var toCode map[string]uint16

var TP_ABOUT uint16
var TP_ANIMAL uint16
var TP_BIRD uint16
var TP_CITY uint16
var TP_COMPANY uint16
var TP_COUNTRY uint16
var TP_DOT uint16
var TP_FISH uint16
var TP_FLY uint16
var TP_FOR uint16
var TP_FROM uint16
var TP_HPERSON uint16
var TP_ILLNESS uint16
var TP_LASTNAME uint16
var TP_MAN uint16
var TP_NAME uint16
var TP_ON uint16
var TP_PATRONYMIC uint16
var TP_PLANET uint16
var TP_POLITIC uint16
var TP_SLANG uint16
var TP_TO uint16

func init() {

	fromCode = make(map[uint16]string, 128)
	toCode = make(map[string]uint16, 128)

	txt := `
	.
	болезнь
	в
	город
	для
	животное
	из
	имя
	истпер
	компания
	мат
	на
	о
	отчество
	планета
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

	TP_ABOUT = ToCode("о")
	TP_ANIMAL = ToCode("животное")
	TP_BIRD = ToCode("птица")
	TP_CITY = ToCode("город")
	TP_COMPANY = ToCode("компания")
	TP_COUNTRY = ToCode("страна")
	TP_DOT = ToCode(".")
	TP_FISH = ToCode("рыба")
	TP_FLY = ToCode("полет")
	TP_FOR = ToCode("для")
	TP_FROM = ToCode("из")
	TP_HPERSON = ToCode("истпер")
	TP_ILLNESS = ToCode("болезнь")
	TP_LASTNAME = ToCode("фамилия")
	TP_MAN = ToCode("человек")
	TP_NAME = ToCode("имя")
	TP_ON = ToCode("на")
	TP_PATRONYMIC = ToCode("отчество")
	TP_PLANET = ToCode("планета")
	TP_POLITIC = ToCode("политик")
	TP_SLANG = ToCode("мат")
	TP_TO = ToCode("в")

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

func FromList(lst ...uint16) int64 {
	val := int64(0)

	for _, v := range lst {
		val = (val << 16) | int64(v)
	}

	return val
}

func AppendCode(val int64, code uint16) int64 {
	return (val << 16) | int64(code)
}
