package prefix

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-19

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text/token"
	"github.com/belfinor/Helium/time/timer"
)

var data map[string][]string

func init() {

	data = make(map[string][]string)

}

func Load(filename string) {

	fh, err := os.Open(filename)
	if err != nil {
		log.Error("error load corpus word prefixes from " + filename)
		return
	}
	defer fh.Close()

	log.Info("reload corpus word prefixes from " + filename)

	load(fh)

}

func LoadFromString(txt string) {

	log.Info("reload corpus word prefixes from text")

	load(strings.NewReader(txt))
}

func load(rh io.Reader) {

	tm := timer.New()

	rb := bufio.NewReader(rh)

	rStr := make(map[string][]string, 128)

	appender := func(pref string, tg string) {
		if list, has := rStr[pref]; has {

			for _, t := range list {
				if t == tg {
					return
				}
			}

			list = append(list, tg)
			rStr[pref] = list

		} else {
			rStr[pref] = []string{tg}
		}
	}

	for {
		str, err := rb.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		toks := token.Make(str)
		if len(toks) > 1 {

			for _, tg := range toks[1:] {

				if len(tg) < 2 || tg[0] != '@' {
					continue
				}

				appender(toks[0], tg[1:])

			}
		}
	}

	data = rStr

	log.Info(fmt.Sprintf("corpus word prefixes reloaded %.4fs", tm.DeltaFloat()))
	log.Info(fmt.Sprintf("corpus word prefixes size = %d", Total()))
}

func Total() int {
	return len(data)
}

func Get(pref string) []string {
	if lst, h := data[pref]; h {
		return lst
	}

	return nil
}
