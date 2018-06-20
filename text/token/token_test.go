package token

// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-20

import "testing"

func TestTextToken(t *testing.T) {

	data := map[string][]string{
		"123":                            []string{"123"},
		"123\\ 123":                      []string{"123 123"},
		"535 t456 876l":                  []string{"535", "t456", "876l"},
		"welcome to \"4th test\"":        []string{"welcome", "to", "4th test"},
		"welcome to \"5th test\" 098765": []string{"welcome", "to", "5th test", "098765"},
		"welcome to \"6th test":          nil,
		"welco\\me to \"7th test\"":      nil,
		"только тест":                    []string{"только", "тест"},
	}

	for k, r := range data {
		res := Make(k)

		if res == nil && r == nil {
			continue
		}

		if res != nil && r == nil || res == nil && r != nil {
			t.Fatal(k)
		}

		if res != nil && r != nil {

			if len(res) != len(r) {
				t.Fatal(k)
			}

			for i, _ := range res {
				if res[i] != r[i] {
					t.Fatal(k)
				}
			}

		}
	}
}
