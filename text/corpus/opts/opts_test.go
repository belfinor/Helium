package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-12

import (
	"testing"
)

func TestOpts(t *testing.T) {

	// test Parse/String
	data1 := map[string]int64{
		"ru.noun.mr": OPT_RU | OPT_NOUN | OPT_MR,
		"ru.verb":    OPT_RU | OPT_VERB,
	}

	for k, v := range data1 {
		res := Parse(k)

		if res != v {
			t.Fatal("opts.Parse " + k + " failed")
		}

		if ToString(res) != k {
			t.Fatal("opts.String " + k + " failed")
		}
	}

	// test Include
	if Include(Parse("ru.noun.mr"), Parse("ru.verb")) {
		t.Fatal("opts.Include test 1 failed")
	}

	if !Include(Parse("ru.noun.mr"), Parse("noun")) {
		t.Fatal("opts.Include 2 test failed")
	}
}
