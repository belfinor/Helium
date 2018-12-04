package forms

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-04

import (
	"strings"

	"github.com/belfinor/Helium/text/corpus/opts"
)

type Form struct {
	Name string
	Opt  opts.Opt
}

func Parse(str string) *Form {
	list := strings.Split(str, ":")

	switch len(list) {
	case 1:
		return &Form{Name: list[0], Opt: 0}
	case 2:
		return &Form{Name: list[0], Opt: opts.Parse(list[1])}
	}

	return nil
}

func (f *Form) String() string {

	builder := strings.Builder{}

	builder.WriteString(f.Name)
	builder.WriteRune(':')
	builder.WriteString(f.Opt.String())

	return builder.String()
}
