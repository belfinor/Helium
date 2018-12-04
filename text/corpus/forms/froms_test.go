package forms

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-04

import (
	"testing"

	"github.com/belfinor/Helium/text/corpus/opts"
)

func TestForms(t *testing.T) {

	f1 := Parse("привет:ru.noun.mr.ip.vp")

	if f1.Name != "привет" {
		t.Fatal("f1.Name failed")
	}

	if f1.Opt != opts.Parse("ru.noun.mr.ip.vp") {
		t.Fatal("f1.Opt failed")
	}

	if f1.String() != "привет:ru.mr.ip.vp.noun" {
		t.Fatal("f1.String failed")
	}
}
