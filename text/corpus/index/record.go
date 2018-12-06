package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-05

import (
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/words"
)

// word object
type Word = words.Word
type Opt = opts.Opt

// corpus index value
type Record struct {
	Name  string
	Words []*Word
}

// get word by filter
func (r *Record) Filter(o Opt) *Record {

	var res *Record

	for _, w := range r.Words {
		if w.Is(r.Name, o) {
			if res == nil {
				res = &Record{
					Name:  res.Name,
					Words: []*Word{w},
				}
			} else {
				res.Words = append(res.Words, w)
			}
		}
	}

	return res
}

// get form by opt
func (r *Record) Form(o Opt) (string, bool) {

	for _, w := range r.Words {
		if str, has := w.Form(o); has {
			return str, true
		}
	}

	return "", false
}
