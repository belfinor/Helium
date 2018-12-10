package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-12-08

import (
	"strings"

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

// get opt code
func (w *Word) GetOpt() opts.Opt {
	return w.opt
}

// check type usage (max type = 4)
func (w *Word) HasType(code uint16) bool {

	return hasCode(w.types, code)
}

// set type
func (w *Word) AddType(code uint16) {

	w.types = addCode(w.types, code)
}

// check tag usage (max tags = 4)
func (w *Word) HasTag(code uint16) bool {

	return hasCode(w.tags, code)
}

func (w *Word) CloneTT(src *Word) {
	w.types = src.types
	w.tags = src.tags
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

func (w *Word) ForEachTypes(fn FOR_EACH_TYPE_FUNC) {
	for i := uint(0); i < 4; i++ {
		v := uint16((w.types >> (i * 16)) & 0xffff)
		if v == 0 {
			break
		}
		fn(v)
	}
}

func (w *Word) IsAlive() bool {
	return w.HasOpt(opts.Opt(opts.OPT_ALIVE))
}

func (w *Word) IsMR() bool {
	return w.HasOpt(opts.Opt(opts.OPT_MR))
}

func (w *Word) IsGR() bool {
	return w.HasOpt(opts.Opt(opts.OPT_GR))
}

func (w *Word) IsSR() bool {
	return w.HasOpt(opts.Opt(opts.OPT_SR))
}

func (w *Word) IsML() bool {
	return w.HasOpt(opts.Opt(opts.OPT_SR))
}

func NounNoun(frms *forms.Forms, w1 *Word, w2 *Word, o opts.Opt) *Word {
	start := frms.Total()

	builder := strings.Builder{}

	for i, f1 := range w1.Forms() {

		builder.Reset()

		builder.WriteString(f1)
		builder.WriteRune(' ')
		builder.WriteString(w2.Form(i))

		frms.Add(builder.String())
	}

	end := frms.Total()

	w := &Word{
		opt:   o,
		start: start,
		end:   end,
		forms: frms,
	}

	return w
}

func NounStr(frms *forms.Forms, w1 *Word, str string, o opts.Opt) *Word {
	start := frms.Total()

	builder := strings.Builder{}

	for _, f1 := range w1.Forms() {

		builder.Reset()

		builder.WriteString(f1)
		builder.WriteRune(' ')
		builder.WriteString(str)

		frms.Add(builder.String())
	}

	end := frms.Total()

	w := &Word{
		opt:   o,
		start: start,
		end:   end,
		forms: frms,
	}

	return w
}

func NounNounNoun(frms *forms.Forms, w1 *Word, w2 *Word, w3 *Word, o opts.Opt) *Word {
	start := frms.Total()

	builder := strings.Builder{}

	for i, f1 := range w1.Forms() {

		builder.Reset()

		builder.WriteString(f1)
		builder.WriteRune(' ')
		builder.WriteString(w2.Form(i))
		builder.WriteRune(' ')
		builder.WriteString(w3.Form(i))

		frms.Add(builder.String())
	}

	end := frms.Total()

	w := &Word{
		opt:   o,
		start: start,
		end:   end,
		forms: frms,
	}

	return w
}

func AdjNoun(frms *forms.Forms, w1 *Word, w2 *Word, o opts.Opt) *Word {

	start := frms.Total()
	builder := strings.Builder{}

	for pl, op := range []opts.Opt{opts.Opt(opts.OPT_MR), opts.Opt(opts.OPT_GR), opts.Opt(opts.OPT_SR), opts.Opt(opts.OPT_ML)} {

		if !w1.HasOpt(op) || !w2.HasOpt(op) {
			continue
		}

		for i := 0; i < 6; i++ {

			j := i

			switch i {
			case 2:
				if pl == 1 {
					j = 1
				}
			case 3:
				if pl == 1 {
					j = 2
				} else if pl == 0 || pl == 3 {
					if w2.HasOpt(opts.Opt(opts.OPT_ALIVE)) {
						j = 1
					} else {
						j = 0
					}
				} else {
					j = 0
				}
			case 4:
				if pl == 1 {
					j = 1
				} else {
					j = i - 1
				}
			case 5:
				if pl == 1 {
					j = 1
				} else {
					j = i - 1
				}
			}

			if i == 3 {
				if w2.HasOpt(opts.Opt(opts.OPT_ALIVE)) {
					j = 1
				} else {
					j = 0
				}
			} else if i > 3 {
				j = i - 1
			}

			builder.WriteString(w1.Form(j))
			builder.WriteRune(' ')
			builder.WriteString(w2.Form(i))

			frms.Add(builder.String())
			builder.Reset()

			frms.Add(builder.String())
			builder.Reset()
		}

		w := &Word{
			opt:   o,
			start: start,
			end:   frms.Total(),
			forms: frms,
		}

		return w
	}

	return nil
}
