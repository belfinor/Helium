package russian

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-10

import (
	"strings"

	snowballword "github.com/belfinor/Helium/text/stemmer/word"
)

// Stem an Russian word.  This is the only exported
// function in this package.
//
func Stem(word string) string {

	word = strings.ToLower(strings.TrimSpace(word))
	w := snowballword.New(word)

	// Return small words and stop words
	if len(w.RS) <= 2 || isStopWord(word) {
		return word
	}

	preprocess(w)
	step1(w)
	step2(w)
	step3(w)
	step4(w)
	return w.String()

}
