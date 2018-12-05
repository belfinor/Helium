package index

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-05

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/words"
	"github.com/belfinor/Helium/time/timer"
)

var data map[string]*Record = map[string]*Record{}

// relaod corpus index from file
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

func load(rh io.Reader) {
	tm := timer.New()

	result := make(map[string]*Record)

	br := bufio.NewReader(rh)

	for {
		str, e := br.ReadString('\n')
		if e != nil {
			break
		}

		w := words.Parse(str)
		if w == nil {
			continue
		}

		for _, f := range w.Forms {
			rec, has := result[f.Name]
			if has {
				rec.Words = append(rec.Words, w)
			} else {
				result[f.Name] = &Record{Name: f.Name, Words: []*Word{w}}
			}
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

var optNum opts.Opt = opts.Parse("num")

// get corpus data for form
func Get(f string) *Record {

	if r, h := data[f]; h {
		return r
	}

	_, err := strconv.Atoi(f)
	if err == nil {

		return &Record{
			Name: f,
			Words: []*Word{&Word{
				Forms: []*forms.Form{&forms.Form{
					Name: f,
					Opt:  optNum,
				}}}},
		}

	}

	return nil
}
