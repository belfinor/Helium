package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-07

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text/corpus/forms"
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

	frms := forms.New(65535, false)

	result := make(map[string]*Record)

	br := bufio.NewReader(rh)

	last := 0
	cur := 0

	for {
		str, e := br.ReadString('\n')
		if e != nil {
			break
		}

		w := words.Parse(str, frms)
		if w == nil {
			continue
		}

		cur = frms.Total()

		fn := func(f string) {

			rec, has := result[f]
			if has {
				rec.Words = append(rec.Words, w)
			} else {
				result[f] = &Record{Name: f, Words: []*Word{w}}
			}
		}

		frms.RangeFunc(last, cur, fn)

		last = cur

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
			Name:  f,
			Words: []*Word{words.MakeNum()},
		}

	}

	return nil
}
