package categorizer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-10

import (
	"container/list"

	"github.com/belfinor/Helium/text/corpus/index"
)

// get *index.Record from *list.Element
func wsFromList(e *list.Element) *index.Record {

	if e == nil {
		return nil
	}

	return index.Get(e.Value.(string))
}
