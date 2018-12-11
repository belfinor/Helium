package tools

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-11

import (
	"container/list"

	"github.com/belfinor/Helium/text/corpus/index"
)

// get *index.Record from *list.Element
func WsFromList(e *list.Element) *index.Record {

	if e == nil {
		return nil
	}

	return index.Get(e.Value.(string))
}
