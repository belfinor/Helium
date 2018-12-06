package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-06

import (
	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/token"
)

type Word struct {
	opt   opts.Opt
	start int
	end   int
	types int64
	tags  int64
}

func Parse(str string, f *forms.Forms) *Word {

	toks := token.Make(str)

	if len(toks) < 2 {
		return nil
	}

	w := &Word{
		opt:   opts.Parse(toks[0]),
		start: f.Total(),
	}

	for _, v := range toks[1:] {

		if len(v) == 0 {
			continue
		}

		f.Add(v)
	}

	w.end = f.Total()

	return w
}

func (w *Word) Form(num int) string {

	if num >= 0 && w.start+num < w.end {
		return forms.Get(w.start + num)
	}

	return ""
}

func (w *Word) IsOpt(opt opts.Opt) bool {

	return w.opt.Include(opt)
}
