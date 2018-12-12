package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-12

import (
	"testing"

	"github.com/belfinor/Helium/text/corpus/opts"
)

func TestIndex(t *testing.T) {

	txt := `
ru.noun.mr президент президента президенту президентом президенте %человек @политика
ru.noun.gr наука науки науке науку наукой %наука @наука
ru.adj.mr российский российского российскому российским российском
`

	LoadFromString(txt)

	if len(data) < 15 {
		t.Fatal("data load failed")
	}

	if Get("абракадабра") != nil {
		t.Fatal("index.Get return unknown value")
	}

	ws := Get("науке")
	if ws == nil {
		t.Fatal("index.Get not found object")
	}

	if !ws.HasOpt(opts.OPT_RU | opts.OPT_NOUN | opts.OPT_GR) {
		t.Fatal("word наука has invalid options")
	}

	if ws.HasOpt(opts.OPT_RU | opts.OPT_ADJ | opts.OPT_GR) {
		t.Fatal("word наука has invalid options")
	}
}
