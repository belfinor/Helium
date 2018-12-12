package tags

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-12

import (
	"testing"
)

func TestTags(t *testing.T) {

	LoadFromString(`
    политика
    общество
    спорт
    музыка
    `)

	if ToCode("политика") != 1 {
		t.Fatal("invalid ToCode for политика")
	}

	if FromCode(2) != "общество" {
		t.Fatal("invalid From code for 2")
	}

	if FromCode(0) != "" {
		t.Fatal("invalid FromCode for 0")
	}

	val := Append(0, 1)
	if val != int64(1) {
		t.Fatal("Append not work")
	}

	val = Append(val, 2)
	if val != (int64(1)<<16)|int64(2) {
		t.Fatal("Append not work")
	}

	val = Append(val, 2)
	if val != (int64(1)<<16)|int64(2) {
		t.Fatal("Append not work")
	}

	if Total() != 4 {
		t.Fatal("Total not work")
	}

	val = Join(1, 2, 3, 2)

	if val != int64(1)<<32|int64(2)<<16|int64(3) {
		t.Fatal("Join not work")
	}
}
