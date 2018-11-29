package slice

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-29

import (
	"fmt"
	"testing"
)

func TestFromMapCnt(t *testing.T) {

	src := map[string]float64{"1": 11.1, "2": 12.2, "3": 13.3, "4": 14.4}

	//res := FromMapCnt(src, &MapCntOpts{MinVal: 12, Reverse: true, Limit: 10})
	res := FromMapCnt(src, nil)

	fmt.Println(res.([]string))
}
