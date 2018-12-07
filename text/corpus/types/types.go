package types

import (
	"bufio"
	"strings"
)

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-07

var fromCode map[uint16]string
var toCode map[string]uint16

func init() {

	fromCode = make(map[uint16]string, 128)
	toCode = make(map[string]uint16, 128)

	txt := `
	город
	животное
	имя
	мат
	отчество
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
