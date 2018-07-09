package text

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-09

import (
	"testing"
)

func TestIsRussian(t *testing.T) {

	data := map[string]bool{
		"привет":   true,
		"hello":    false,
		"test":     false,
		"игра":     true,
		"политика": true,
	}

	for k, v := range data {
		if IsRussian(k) != v {
			t.Fatal("IsRussian error on word: " + k)
		}
	}
}
