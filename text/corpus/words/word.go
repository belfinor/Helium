package words

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-05

import (
	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/token"
)

type Word struct {
	Forms []*forms.Form
	Tags  []string
}

func Parse(str string) *Word {

	toks := token.Make(str)

	if len(toks) == 0 {
		return nil
	}

	w := &Word{
		Forms: []*forms.Form{},
		Tags:  []string{},
	}

	for _, v := range toks {

		if len(v) == 0 {
			continue
		}

		if v[0] == '%' && len(v) > 1 {
			w.Tags = append(w.Tags, v[1:])
			continue
		}

		f := forms.Parse(v)
		if f == nil {
			continue
		}

		w.Forms = append(w.Forms, f)
	}

	if len(w.Forms) == 0 {
		return nil
	}

	return w
}
