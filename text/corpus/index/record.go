package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-06

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

// check has word with oopts
func (r *Record) HasOpt(o Opt) bool {

	for _, w := range r.Words {
		if w.HasOpt(o) {
			return true
		}
	}

	return false
}

// get form by num for word include opts
func (r *Record) OptForm(o Opt, num int) string {

	for _, w := range r.Words {
		if w.HasOpt(o) {
			return w.Form(num)
		}
	}

	return ""
}
