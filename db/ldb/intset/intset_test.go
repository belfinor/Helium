package intset

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-06-08

import (
	"github.com/belfinor/Helium/db/ldb"
	"github.com/belfinor/Helium/test"
	"testing"
)

func TestDBIntSet(t *testing.T) {
	test.Init()

	key := []byte("test.key")
	list := []int64{1, 2, 3}

	Create(key, list...)

	res := Get(key)

	if res == nil || len(res) != len(list) {
		t.Fatal("Create error")
	}

	for i, val := range res {
		if list[i] != val {
			t.Fatal("Create add wrong data")
		}
	}

	Pop(key, 2)
	list = []int64{1, 3}
	res = Get(key)

	if res == nil || len(res) != len(list) {
		t.Fatal("Pop error")
	}

	for i, val := range res {
		if list[i] != val {
			t.Fatal("Pop result wrong data")
		}
	}

	Push(key, 13, 10, 9)
	list = []int64{1, 3, 13, 10, 9}
	res = Get(key)

	if res == nil || len(res) != len(list) {
		t.Fatal("Push error")
	}

	for i, val := range res {
		if list[i] != val {
			t.Fatal("Push result wrong data")
		}
	}

	Remove(key)

	if ldb.Get(key) != nil {
		t.Fatal("Remove not work")
	}
}
