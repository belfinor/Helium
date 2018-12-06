package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-04

import (
	"testing"
)

func TestOpts(t *testing.T) {

	// test Parse/String
	data1 := map[string]int32{
		"ru.noun.mr": OPT_RU | OPT_NOUN | OPT_MR,
		"ru.verb":    OPT_RU | OPT_VERB,
	}

	for k, v := range data1 {
		res := Parse(k)

		if int32(res) != v {
			t.Fatal("opts.Parse " + k + " failed")
		}

		if res.String() != k {
			t.Fatal("opts.String " + k + " failed")
		}
	}

	// test Include
	if Parse("ru.noun.mr").Include(Parse("ru.verb")) {
		t.Fatal("opts.Include test 1 failed")
	}

	if !Parse("ru.noun.mr").Include(Parse("noun")) {
		t.Fatal("opts.Include 2 test failed")
	}
}
