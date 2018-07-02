package uniq

import (
	"testing"
)

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-02

func TestUniq(t *testing.T) {

	u := New()
	defer u.Close()

	for i := 0; i < 20; i++ {
		if u.Next() == "" {
			t.Fatal("not worl")
		}
	}
}
