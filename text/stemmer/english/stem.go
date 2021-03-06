package english

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-07-10

import (
	"strings"

	snowballword "github.com/belfinor/Helium/text/stemmer/word"
)

// Stem an English word.  This is the only exported
// function in this package.
//
func Stem(word string) string {

	word = strings.ToLower(strings.TrimSpace(word))

	// Return small words and stop words
	if len(word) <= 2 || isStopWord(word) {
		return word
	}

	// Return special words immediately
	if specialVersion := stemSpecialWord(word); specialVersion != "" {
		word = specialVersion
		return word
	}

	w := snowballword.New(word)

	// Stem the word.  Note, each of these
	// steps will alter `w` in place.
	//
	preprocess(w)
	step0(w)
	step1a(w)
	step1b(w)
	step1c(w)
	step2(w)
	step3(w)
	step4(w)
	step5(w)
	postprocess(w)

	return w.String()

}
