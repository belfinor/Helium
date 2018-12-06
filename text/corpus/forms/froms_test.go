package forms

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-06

import (
	"testing"
)

func TestForms(t *testing.T) {

	f := New(10, false)
	if f == nil {
		t.Fatal("forms.New failed")
	}

	for _, v := range []string{"1", "2", "3", "4", "5"} {
		f.Add(v)
	}

	if f.Total() != 5 {
		t.Fatal("forms.Add/Total not work")
	}

	if f.Get(3) != "4" {
		t.Fatal("forms.Get not work")
	}

	res := f.Range(1, 4)
	if len(res) != 3 {
		t.Fatal("forms.Range not work")
	}

	for i, v := range []string{"2", "3", "4"} {
		if res[i] != v {
			t.Fatal("forms.Range not work")
		}
	}
}
