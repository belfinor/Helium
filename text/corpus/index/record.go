package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-12-11

import (
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/words"
)

// aliases for external types
type Word = words.Word
type Opt = opts.Opt
type FOR_EACH_TAG_FUNC = words.FOR_EACH_TAG_FUNC
type FOR_EACH_TYPE_FUNC = words.FOR_EACH_TYPE_FUNC

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

func (r *Record) WordByType(code uint16) *Word {
	for _, w := range r.Words {
		if w.HasType(code) {
			return w
		}
	}

	return nil
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

func (r *Record) ForEachType(fn FOR_EACH_TYPE_FUNC) {
	for _, w := range r.Words {
		w.ForEachTypes(fn)
	}
}

func (r *Record) TT2Word(w *Word) {
	w.CloneTT(r.Words[0])
}
