package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-06

import (
	"github.com/belfinor/Helium/text/corpus/index"
)

func Load(filename string) {
	index.Load(filename)
}
