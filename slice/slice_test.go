package slice

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-29

import (
	"testing"
)

func TestSlice(t *testing.T) {

	orig := []int{1, 2, 3, 4, 5}

	Reverse(orig)

	if len(orig) != 5 {
		t.Fatal("slice.Reverse corrupt data")
	}

	for i, v := range orig {
		if v != 5-i {
			t.Fatal("slice.Reverse not work")
		}
	}

	Shuffle(orig)

	if len(orig) != 5 {
		t.Fatal("slice.Shuffle corrupt data")
	}
}
