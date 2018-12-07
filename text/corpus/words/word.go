package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-07

import (
	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/token"
)

type Word struct {
	opt   opts.Opt
	start int
	end   int
	types int64
	tags  int64
	forms *forms.Forms
}

type FOR_EACH_TAG_FUNC func(uint16)
type FOR_EACH_TYPE_FUNC func(uint16)

// parse word from string
func Parse(str string, f *forms.Forms) *Word {

	toks := token.Make(str)

	if len(toks) < 1 {
		return nil
	}

	if toks[0] == ";" {
		return nil
	}

	w := &Word{
		opt:   opts.Parse(toks[0]),
		start: f.Total(),
		forms: f,
	}

	for _, v := range toks[1:] {

		if len(v) == 0 {
			continue
		}

		if len(v) > 1 {

			if v[0] == '%' {
				w.types = addCode(w.types, types.ToCode(v[1:]))
				continue
			}

			if v[0] == '@' {
				w.tags = addCode(w.tags, tags.ToCode(v[1:]))
				continue
			}
		}

		f.Add(v)
	}

	w.end = f.Total()

	return w
}

// make word for number
func MakeNum() *Word {
	return &Word{
		opt:   opts.Opt(opts.OPT_NUM),
		start: 0,
		end:   0,
	}
}

// get from by number
func (w *Word) Form(num int) string {

	if num >= 0 && w.start+num < w.end {
		return w.forms.Get(w.start + num)
	}

	return ""
}

// get forms slice
func (w *Word) Forms() []string {
	return w.forms.Range(w.start, w.end)
}

// check word has opt
func (w *Word) HasOpt(opt opts.Opt) bool {

	return w.opt.Include(opt)
}

// check type usage (max type = 4)
func (w *Word) HasType(code uint16) bool {

	return hasCode(w.types, code)
}

// check tag usage (max tags = 4)
func (w *Word) HasTag(code uint16) bool {

	return hasCode(w.tags, code)
}

func hasCode(data int64, code uint16) bool {

	for i := uint(0); i < 4; i++ {

		cur := uint16(data>>(i*16)) & 0xffff

		if cur == code {
			return true
		}
	}

	return false
}

func addCode(src int64, code uint16) int64 {

	if code != 0 {
		return src<<16 | int64(code)
	}

	return src
}

func (w *Word) ForEachTags(fn FOR_EACH_TAG_FUNC) {
	for i := uint(0); i < 4; i++ {
		v := uint16((w.tags >> (i * 16)) & 0xffff)
		if v == 0 {
			break
		}
		fn(v)
	}
}
