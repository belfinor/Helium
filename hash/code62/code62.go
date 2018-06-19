package code62

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-04-26

var tab []string = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l",
	"z", "x", "c", "v", "b", "n", "m",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L",
	"Z", "X", "C", "V", "B", "N", "M",
}

var size int64 = int64(len(tab))

func Calc(val int64) string {

	res := ""
	sign := ""

	if val < 0 {
		sign = "-"
		val *= -1
	}

	for val > 0 {
		i := int(val % size)
		res = tab[i] + res
		val = val / size
	}

	if res == "" {
		res = "q"
	}

	res = sign + res

	return res
}
