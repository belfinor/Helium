package code62

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-04-26

import (
	"testing"
)

func TestCalc(t *testing.T) {

	if Calc(205400) != "K0X" {
		t.Fatal("calc error")
	}
}
