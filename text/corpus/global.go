package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-14

import (
	"sync"
)

var mutex sync.RWMutex = sync.RWMutex{}
