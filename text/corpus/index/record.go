package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-05

import (
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/words"
)

// word object
type Word = words.Word
type Opt = opts.Opt

// corpus index value
type Record struct {
	Name  string
	Words []*Word
}
