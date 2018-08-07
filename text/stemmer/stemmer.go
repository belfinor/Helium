package stemmer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-08-07

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

	if word == "." {
		return word
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
		if err != nil {
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
