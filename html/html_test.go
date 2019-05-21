package html

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-04-15

import (
	"testing"
)

func TestTextHtmlProcessString(t *testing.T) {
	h := NewHtmlParser()

	if h == nil {
		t.Fatal("NewHtmlParser() erro")
	}

	h.AltExport = true

	src := `<html><head><title>Привет мир</title><style>a{color:#CCC;}</style></head><body><h1>Hello</h1><script><!-- alert('hello')
--></script><iframe src="xxx"></iframe><a href='/url1'>url1</a><a href='/url2'>url2</a><img src="" title="img title"/></body></html>`

	res := h.ProcessString(src)

	if res != " Привет мир Hello url1 url2 .\nimg title .\n" {
		t.Fatal("Wrong plain text")
	}

	if len(h.Iframes) != 1 || !h.Iframes["xxx"] {
		t.Fatal("Iframe src not found")
	}
}
