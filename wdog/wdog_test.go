package wdog

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-02

import (
	"testing"
	"time"
)

func TestWatchDog(t *testing.T) {
	wd := New(time.Nanosecond * 500000000)
	defer wd.Close()

	for i := 0; i < 10; i++ {
		wd.Alive()
		<-time.After(time.Nanosecond * 100000000)
	}

}
