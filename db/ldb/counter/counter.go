package counter

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.005
// @date    2018-04-26

import (
	"sync"

	"github.com/belfinor/Helium/db/ldb"
	"github.com/belfinor/Helium/pack"
)

var mutex sync.Mutex

func Inc(key []byte, val int64) int64 {

	mutex.Lock()
	defer mutex.Unlock()

	raw := ldb.Get(key)
	cur := pack.Bytes2Int(raw)

	cur += val

	if cur <= int64(0) {
		ldb.Del(key)
		cur = 0
	} else {
		ldb.Set(key, pack.Int2Bytes(cur))
	}

	return cur
}

func Get(key []byte) int64 {
	raw := ldb.Get(key)
	cur := pack.Bytes2Int(raw)
	return cur
}

func Reset(key []byte) {

	mutex.Lock()
	defer mutex.Unlock()

	ldb.Del(key)
}

func Set(key []byte, val int64) {

	mutex.Lock()
	defer mutex.Unlock()

	if val <= int64(0) {
		ldb.Del(key)
	} else {
		ldb.Set(key, pack.Int2Bytes(val))
	}
}
