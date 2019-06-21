package stemmer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.007
// @date    2019-06-21

import (
	"bufio"
	"os"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/stemmer/english"
	"github.com/belfinor/Helium/text/stemmer/russian"
	"github.com/belfinor/Helium/text/token"
)

// Stop words
//
var stopWords map[string]string = map[string]string{}

// Stem a word
//
func Word(word string) string {

	if wf, has := stopWords[word]; has {
		return wf
	}

	if word == "." || word == "," {
		return word
	}

	for _, run := range word {
		if run == '-' {

			lst := strings.Split(word, "-")

			res := make([]string, 0, len(lst))

			for _, w := range lst {

				if w == "." {
					res = append(res, w)
					continue
				}

				if text.IsRussian(word) {
					res = append(res, Word(w))
				} else {
					res = append(res, Word(w))
				}

			}

			return strings.Join(res, "-")
		}
	}

	if text.IsRussian(word) {
		return russian.Stem(word)
	}

	return english.Stem(word)
}

// TextToCode
//
func TextToCode(str string) string {
	res := make([]string, 0, 4)

	for wrd := range text.WordStream(strings.NewReader(str)) {
		rec := Word(wrd)
		res = append(res, rec)
	}

	return strings.Join(res, " ")
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

// Load global stop words with fixed forms list
//
func LoadStopWords(filename string) {

	fh, err := os.Open(filename)
	if err != nil {
		log.Error(err)
		return
	}
	defer fh.Close()

	log.Info("load stop words from file " + filename)

	reader := bufio.NewReader(fh)
	result := make(map[string]string)

	for {
		str, err := reader.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		list := token.Make(str)
		if len(list) == 0 {
			continue
		}

		if len(list) > 2 {
			log.Error("error string: " + str)
			continue
		}

		if len(list) == 1 {
			list = append(list, list[0])
		}

		result[list[0]] = list[1]
	}

	stopWords = result
}
