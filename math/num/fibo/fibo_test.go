package fibo

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-02

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	seq := New()
	defer seq.Close()

	if seq.Next() != 1 {
		t.Fatal("MakeSeq error")
	}

	if seq.Next() != 1 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 2 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 3 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 5 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 8 {
		t.Fatal("Next not work")
	}
}
