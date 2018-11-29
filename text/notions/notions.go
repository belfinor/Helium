package notions

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-28

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/buffer"
	"github.com/belfinor/Helium/text/stemmer"
)

var skip map[string]string = map[string]string{}
var known map[string]string = map[string]string{}

// Fetch unknown notions with counter > 1
//
func FindNew(src io.RuneReader) map[string]int {

	wordStream := text.WordStream(src, text.WSO_NO_URLS, text.WSO_HASHTAG, text.WSO_ENDS)
	tokStream := stemmer.Stream(wordStream)

	unknown := make(map[string]int)
	list := []string{}

	bfr := buffer.New(3)

	procBfr := func() {

		if bfr.Empty() {
			return
		}

		if bfr.Size() == 3 {

			str := bfr.Join(" ", 3)

			if _, h := skip[str]; h {
				bfr.Shift(3)
				return
			}

			if _, h := known[str]; h {
				bfr.Shift(3)
				return
			}

			unknown[str]++
		}

		if bfr.Size() == 2 {
			str := bfr.Join(" ", 2)

			if _, h := skip[str]; h {
				bfr.Shift(2)
				return
			}

			if _, h := known[str]; h {
				bfr.Shift(2)
				return
			}

			unknown[str]++
		}

		str := bfr.Get(0)

		if _, h := skip[str]; !h {
			if _, h = known[str]; !h {
				unknown[str]++
			}
		}

		bfr.Shift(1)
	}

	procList := func() {

		for _, v := range list {
			if bfr.Full() {
				procBfr()
			}

			bfr.Add(v)
		}

		for {

			if bfr.Empty() {
				break
			}
		}

	}

	for val := range tokStream {

		if val == "." {
			procList()
		} else {
			list = append(list, val)
		}
	}

	procList()

	res := make(map[string]int)

	for k, v := range unknown {
		if v > 1 {
			res[k] = v
		}
	}

	return res
}

func load(filename string) map[string]string {

	rf, err := os.Open(filename)
	if err != nil {
		log.Error("open " + filename + " error: " + err.Error())
		return nil
	}
	defer rf.Close()

	rb := bufio.NewReader(rf)

	orig := strings.Builder{}
	short := strings.Builder{}

	res := make(map[string]string)

	for {

		str, err := rb.ReadString('\n')
		if err != nil {
			break
		}

		i := 0

		orig.Reset()
		short.Reset()

		for w := range text.WordStream(strings.NewReader(str)) {

			if i > 0 {
				orig.WriteRune(' ')
				short.WriteRune(' ')
			}

			orig.WriteString(w)
			short.WriteString(stemmer.Word(w))

			i++
		}

		if i > 0 {
			res[short.String()] = orig.String()
		}

	}

	return res
}

func LoadIgnore(filename string) {
	r := load(filename)
	if r != nil {
		skip = r
		log.Info("notions.LoadIgnore from " + filename + " success")
	}
}

func LoadNotions(filename string) {
	r := load(filename)
	if r != nil {
		known = r
		log.Info("notions.LoadNotions from " + filename + " success")
	}
}
