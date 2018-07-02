package fibo

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-02

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	seq := New()

	if seq.Get() != 1 {
		t.Fatal("MakeSeq error")
	}

	if seq.Get() != 1 {
		t.Fatal("Next not work")
	}

	if seq.Get() != 2 {
		t.Fatal("Next not work")
	}

	if seq.Get() != 3 {
		t.Fatal("Next not work")
	}

	if seq.Get() != 5 {
		t.Fatal("Next not work")
	}

	if seq.Get() != 8 {
		t.Fatal("Next not work")
	}

	seq.Close()
}
