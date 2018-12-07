package types

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-07

var fromCode map[uint16]string
var toCode map[string]uint16

func init() {

	fromCode := make(map[uint16]string, 128)
	toCode := make(map[string]uint16, 128)

	for i, v := range []string{"город", "имя", "лицо", "отчество", "мат", "страна", "фамилия", "человек"} {
		fromCode[uint16(i+1)] = v
		toCode[v] = uint16(i + 1)
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
