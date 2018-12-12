package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-12

import (
	"github.com/belfinor/Helium/text/corpus/categorizer/schemas"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
)

func Load(dir string) {
	types.Load(dir + "/types.txt")
	tags.Load(dir + "/tags.txt")
	index.Load(dir + "/corpus.txt")
	schemas.Load(dir + "/schema.txt")
}
