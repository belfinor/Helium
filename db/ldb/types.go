package ldb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-06-28

type FOR_EACH_KEY_FUNC func([]byte) bool
type FOR_EACH_FUNC func([]byte, []byte) bool
