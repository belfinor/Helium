package prefix

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-19

import (
	"testing"
)

func TestPrefix(t *testing.T) {
	txt := `
  ai @технологии
  python @it
  vr @технологии @техника
  `

	LoadFromString(txt)

	wait := map[string][]string{
		"ai":     []string{"технологии"},
		"python": []string{"it"},
		"vr":     []string{"технологии", "техника"},
	}

	for k, v := range wait {

		res := Get(k)

		if len(res) != len(v) {
			t.Fatal("fail " + k + " length")
		}

		for i, item := range res {
			if item != v[i] {
				t.Fatal("wrong value for " + k)
			}
		}
	}
}
