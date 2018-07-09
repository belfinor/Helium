package stemmer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-09

import (
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/stemmer/english"
	"github.com/belfinor/Helium/text/stemmer/russian"
)

// Stem a word
//
func Word(word string) string {

	if text.IsRussian(word) {
		return russian.Stem(word, true)
	}

	return english.Stem(word, true)
}

// Stem from chan to chan
//
func Stream(input <-chan string) <-chan string {
	output := make(chan string, 10000)

	go func() {
		for w := range input {
			output <- Word(w)
		}
		close(output)
	}()

	return output
}
