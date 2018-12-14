package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-14

import (
	"github.com/belfinor/Helium/text/corpus/words"
)

// aliases for external types
type Word = words.Word

// corpus index value
type Record struct {
	Name     string
	Words    []*Word
	TypeMask int64
}

// check has word with oopts
func (r *Record) HasOpt(o int64) bool {

	for _, w := range r.Words {
		if w.HasOpt(o) {
			return true
		}
	}

	return false
}

func (r *Record) WordByOpt(o int64) *Word {

	for _, w := range r.Words {
		if w.HasOpt(o) {
			return w
		}
	}

	return nil
}

func (r *Record) WordByMask(m int64) *Word {
	for _, w := range r.Words {
		if w.TypeMask()&m == m {
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

func (r *Record) ForEachTags(fn words.FOREACH_UINT16_FUNC) {
	for _, w := range r.Words {
		w.ForEachTags(fn)
	}
}

func (r *Record) ForEachType(fn words.FOREACH_UINT16_FUNC) {
	for _, w := range r.Words {
		w.ForEachTypes(fn)
	}
}

func (r *Record) TT2Word(w *Word) {
	w.CopyTT(r.Words[0])
}
