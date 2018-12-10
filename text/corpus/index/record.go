package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-07

import (
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/words"
)

// aliases for external types
type Word = words.Word
type Opt = opts.Opt
type FOR_EACH_TAG_FUNC = words.FOR_EACH_TAG_FUNC

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

func (r *Record) WordByOpt(o Opt) *Word {

	for _, w := range r.Words {
		if w.HasOpt(o) {
			return w
		}
	}

	return nil
}

// check has type
func (r *Record) HasType(code uint16) bool {
	for _, w := range r.Words {
		if w.HasType(code) {
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

func (r *Record) ForEachTags(fn FOR_EACH_TAG_FUNC) {
	for _, w := range r.Words {
		w.ForEachTags(fn)
	}
}

func (r *Record) TT2Word(w *Word) {
	w.CloneTT(r.Words[0])
}
