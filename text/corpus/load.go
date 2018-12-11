package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-11

import (
	"github.com/belfinor/Helium/text/corpus/categorizer/schemas"
	"github.com/belfinor/Helium/text/corpus/index"
)

func Load(corpus string, schema string) {
	index.Load(corpus)
	schemas.Load(schema)
}
