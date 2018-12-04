package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-04

import (
	"testing"
)

func TestOpts(t *testing.T) {

	// test Parse/String
	data1 := map[string]int64{
		"ru.mr.ip.noun": OPT_RU | OPT_NOUN | OPT_MR | OPT_IP,
		"ru.verb.undef": OPT_RU | OPT_VERB | OPT_UNDEF,
	}

	for k, v := range data1 {
		res := Parse(k)

		if int64(res) != v {
			t.Fatal("opts.Parse " + k + " failed")
		}

		if res.String() != k {
			t.Fatal("opts.String " + k + " failed")
		}
	}

	// test Include
	if Parse("ru.noun.mr.ip").Include(Parse("ru.verb.undef")) {
		t.Fatal("opts.Include test 1 failed")
	}

	if !Parse("ru.noun.mr.ip").Include(Parse("noun.ip")) {
		t.Fatal("opts.Include 2 test failed")
	}
}
