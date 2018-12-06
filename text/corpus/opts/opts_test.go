package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-06

import (
	"testing"
)

func TestOpts(t *testing.T) {

	// test Parse/String
	data1 := map[string]Opt{
		"ru.noun.mr": Opt(OPT_RU | OPT_NOUN | OPT_MR),
		"ru.verb":    Opt(OPT_RU | OPT_VERB),
	}

	for k, v := range data1 {
		res := Parse(k)

		if res != v {
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
