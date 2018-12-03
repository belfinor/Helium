package notions

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-03

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/belfinor/Helium/chars"
	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/slice"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/buffer"
	"github.com/belfinor/Helium/text/stemmer"
)

var skip map[string]string = map[string]string{}
var known map[string]string = map[string]string{}

// Fetch unknown notions with counter > 1
//
func FindNew(src io.RuneReader) ([]string, []string) {

	wordStream := text.WordStream(src, text.WSO_NO_URLS, text.WSO_HASHTAG, text.WSO_ENDS)
	tokStream := stemmer.Stream(wordStream)

	unknown := make(map[string]int)
	list := make([]string, 0, 100)
	used := make(map[string]bool)
	asso := make(map[string]int)

	bfr := buffer.New(3)

	strBuilder := strings.Builder{}

	procAsso := func() {

		prev := make([]string, 0, len(used))

		for k := range used {

			for _, v := range prev {
				strBuilder.WriteString(v)
				strBuilder.WriteRune(chars.ARROW_LEFT)
				strBuilder.WriteString(k)

				asso[strBuilder.String()]++

				strBuilder.Reset()

				strBuilder.WriteString(k)
				strBuilder.WriteRune(chars.ARROW_LEFT)
				strBuilder.WriteString(v)

				asso[strBuilder.String()]++

				strBuilder.Reset()
			}

			prev = append(prev, k)
		}

		used = map[string]bool{}
	}

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
				used[str] = true
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
				used[str] = true
				bfr.Shift(2)
				return
			}

			unknown[str]++
		}

		str := bfr.Get(0)

		if _, h := skip[str]; !h {
			if _, h = known[str]; !h {
				unknown[str]++
			} else {
				used[str] = true
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

			procBfr()
		}

		list = list[:0]
		procAsso()
	}

	for val := range tokStream {

		if val == "." {
			procList()
		} else {
			list = append(list, val)
		}
	}

	procList()

	return slice.FromMapCnt(unknown, &slice.MapCntOpts{MinVal: 2, Limit: 1000}).([]string), slice.FromMapCnt(asso, &slice.MapCntOpts{MinVal: 2, Limit: 1000}).([]string)
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
