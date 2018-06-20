package fibonacci

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-29

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	seq := New()

	if seq.GetNumber() != 0 || seq.GetCurrent() != 1 {
		t.Fatal("MakeSeq error")
	}

	if seq.Next() != 1 || seq.GetNumber() != 1 || seq.GetCurrent() != 1 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 2 || seq.GetNumber() != 2 || seq.GetCurrent() != 2 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 3 || seq.GetNumber() != 3 || seq.GetCurrent() != 3 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 5 || seq.GetNumber() != 4 || seq.GetCurrent() != 5 {
		t.Fatal("Next not work")
	}

	if seq.Next() != 8 || seq.GetNumber() != 5 || seq.GetCurrent() != 8 {
		t.Fatal("Next not work")
	}

	seq.Reset()

	if seq.GetNumber() != 0 || seq.GetCurrent() != 1 {
		t.Fatal("Next not work")
	}
}
