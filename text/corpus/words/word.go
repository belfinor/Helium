package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-12

import (
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/token"
)

type Word struct {
	base  string
	opt   int64
	types int64
	tags  int64
}

type FOREACH_UINT16_FUNC func(uint16)
type FORM_CALLBACK func(string, *Word)

// parse word from string
func Parse(str string, fn FORM_CALLBACK) *Word {

	toks := token.Make(str)

	if len(toks) < 1 || toks[0] == ";" || len(toks) < 2 {
		return nil
	}

	w := &Word{
		base: toks[1],
		opt:  opts.Parse(toks[0]),
	}

	for _, v := range toks[1:] {

		if len(v) == 0 {
			continue
		}

		if len(v) > 1 {

			if v[0] == '%' {
				w.types = types.Append(w.types, types.ToCode(v[1:]))
				continue
			}

			if v[0] == '@' {
				w.tags = tags.Append(w.tags, tags.ToCode(v[1:]))
				continue
			}
		}

		fn(v, w)
	}

	return w
}

// make word for number
func MakeNum(str string) *Word {
	return &Word{
		base:  str,
		opt:   opts.OPT_NUM,
		types: int64(types.TP_NUMBER),
	}
}

// copy tags&types from other
func (w *Word) CopyTT(src *Word) {
	w.types = src.types
	w.tags = src.tags
}

// work with opts

func (w *Word) HasOpt(opt int64) bool {

	return opts.Include(w.opt, opt)
}

func (w *Word) GetOpt() int64 {
	return w.opt
}

func (w *Word) SetOpt(v int64) {
	w.opt = v
}

// work wth types

func (w *Word) HasType(code uint16) bool {

	return types.Has(w.types, code)
}

func (w *Word) AddType(code uint16) {

	w.types = types.Append(w.types, code)
}

func (w *Word) ForEachTypes(fn FOREACH_UINT16_FUNC) {
	types.ForEach(w.types, types.FOREACH_FUNC(fn))
}

// work with tags

func (w *Word) HasTag(code uint16) bool {

	return tags.Has(w.tags, code)
}

func (w *Word) AddTag(code uint16) {

	w.tags = tags.Append(w.tags, code)
}

func (w *Word) ForEachTags(fn FOREACH_UINT16_FUNC) {
	tags.ForEach(w.tags, tags.FOREACH_FUNC(fn))
}
