package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-14

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/corpus/words"
	"github.com/belfinor/Helium/time/timer"
)

var data map[string]*Record = map[string]*Record{}

// reload corpus index from file
func Load(filename string) {

	fh, err := os.Open(filename)
	if err != nil {
		log.Error("error load corpus index from " + filename)
		return
	}
	defer fh.Close()

	log.Info("reload corpus index from " + filename)

	load(fh)

}

// reload corpus from txt
func LoadFromString(txt string) {

	log.Info("reload corpus from text")

	load(strings.NewReader(txt))
}

func load(rh io.Reader) {
	tm := timer.New()

	result := make(map[string]*Record)

	br := bufio.NewReader(rh)

	callback := func(f string, w *Word) {

		rec, has := result[f]
		if has {
			rec.Words = append(rec.Words, w)
			rec.TypeMask |= w.TypeMask()
		} else {
			result[f] = &Record{Name: f, Words: []*Word{w}, TypeMask: w.TypeMask()}
		}

	}

	for {
		str, e := br.ReadString('\n')
		if e != nil && str == "" {
			break
		}

		w := words.Parse(str, callback)
		if w == nil {
			continue
		}
	}

	data = result

	log.Info(fmt.Sprintf("corpus index reloaded %.4fs", tm.DeltaFloat()))
	log.Info(fmt.Sprintf("corpus index size = %d", Total()))
}

// get corpus index size
func Total() int {
	return len(data)
}

// get corpus data for form
func Get(f string) *Record {

	if r, h := data[f]; h {
		return r
	}

	_, err := strconv.Atoi(f)
	if err == nil {

		return &Record{
			Name:     f,
			Words:    []*Word{words.MakeNum(f)},
			TypeMask: types.MTP_NUMBER,
		}

	}

	return nil
}
