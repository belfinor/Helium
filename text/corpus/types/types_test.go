package types

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-14

import (
	"testing"
)

func TestTypes(t *testing.T) {

	LoadFromString("игра")

	if ToCode("игра") != 1 {
		t.Fatal("invalid ToCode for игра")
	}

	if FromCode(2) != "имя" {
		t.Fatal("invalid From code for 2")
	}

	if FromCode(0) != "" {
		t.Fatal("invalid FromCode for 0")
	}

	val := Append(0, TP_MAN)
	if val != int64(TP_MAN) {
		t.Fatal("Append not work")
	}

	val = Append(val, TP_NAME)
	if val != (int64(TP_MAN)<<16)|int64(TP_NAME) {
		t.Fatal("Append not work")
	}

	val = Append(val, TP_NAME)
	if val != (int64(TP_MAN)<<16)|int64(TP_NAME) {
		t.Fatal("Append not work")
	}

	if Total() != 16 {
		t.Fatal("Total not work")
	}

	val = Join(TP_NAME, TP_PATRONYMIC, TP_LASTNAME, TP_LASTNAME)

	if val != int64(TP_NAME)<<32|int64(TP_PATRONYMIC)<<16|int64(TP_LASTNAME) {
		t.Fatal("Join not work")
	}
}
