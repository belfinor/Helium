package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.005
// @date    2018-12-14

import (
	"strings"

	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/token"
)

type Word struct {
	Base     string
	opt      int64
	types    int64
	tags     int64
	typemask int64
}

type FOREACH_UINT16_FUNC func(uint16)
type FORM_CALLBACK func(string, *Word)

// parse word from string
func Parse(str string, fn FORM_CALLBACK) *Word {

	toks := token.Make(str)

	if len(toks) < 1 || toks[0] == ";" || len(toks) < 2 {
		return nil
	}

	has := make(map[string]bool)

	w := &Word{
		Base: toks[1],
		opt:  opts.Parse(toks[0]),
	}

	for _, v := range toks[1:] {

		if len(v) == 0 {
			continue
		}

		if len(v) > 1 {

			if v[0] == '%' {
				name := v[1:]
				w.types = types.Append(w.types, types.ToCode(name))
				w.typemask = w.typemask | types.ToMask(name)
				continue
			}

			if v[0] == '@' {
				w.tags = tags.Append(w.tags, tags.ToCode(v[1:]))
				continue
			}
		}

		has[v] = true
	}

	for v := range has {
		fn(v, w)
	}

	return w
}

// get type mask
func (w *Word) TypeMask() int64 {
	return w.typemask
}

// make word for number
func MakeNum(str string) *Word {
	return &Word{
		Base:     str,
		opt:      opts.OPT_NUM,
		types:    int64(types.TP_NUMBER),
		typemask: types.MTP_NUMBER,
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
	w.typemask = w.typemask | types.ToMask(types.FromCode(code))
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

// Join words to single

func Join(o int64, wl ...*Word) *Word {

	builder := strings.Builder{}

	for i, w := range wl {

		if i > 0 {
			builder.WriteRune(' ')
		}

		builder.WriteString(w.Base)
	}

	return &Word{
		Base: builder.String(),
		opt:  o,
	}
}
