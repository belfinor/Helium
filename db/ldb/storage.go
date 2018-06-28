package ldb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-06-28

type Storage interface {
	Close()
	Set([]byte, []byte)
	Get([]byte) []byte
	Has([]byte) bool
	Del([]byte)
	Total(prefix []byte) int64
	List(prefix []byte, limit int, offset int, RemovePrefix bool) [][]byte
	ForEach(prefix []byte, RemovePrefix bool, fn FOR_EACH_FUNC)
	ForEachKey(prefix []byte, limit int, offset int, RemovePrefix bool, fn FOR_EACH_KEY_FUNC)
}
