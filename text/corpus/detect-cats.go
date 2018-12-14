package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-14

import (
	"io"

	"github.com/belfinor/Helium/text/corpus/categorizer"
)

func DetectCats(rh io.RuneReader) ([]string, bool) {

	mutex.RLock()
	defer mutex.RUnlock()

	return categorizer.New().Proc(rh)
}
