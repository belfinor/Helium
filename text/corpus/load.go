package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-19

import (
	"github.com/belfinor/Helium/text/corpus/categorizer/schemas"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/prefix"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
)

func Load(dir string) {

	mutex.Lock()
	defer mutex.Unlock()

	types.Load(dir + "/types.txt")
	tags.Load(dir + "/tags.txt")
	index.Load(dir + "/corpus.txt")
	prefix.Load(dir + "/prefixes.txt")
	schemas.Load(dir + "/schema.txt")
}
