package schemas

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-11

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/token"
	"github.com/belfinor/Helium/time/timer"
)

// reload corpus index from file
func Load(filename string) {

	fh, err := os.Open(filename)
	if err != nil {
		log.Error("error load corpus schema from " + filename)
		return
	}
	defer fh.Close()

	log.Info("reload corpus schema from " + filename)

	load(fh)

}

// reload corpus from txt
func LoadFromString(txt string) {

	log.Info("reload corpus schema from text")

	load(strings.NewReader(txt))
}

func load(rh io.Reader) {
	tm := timer.New()

	result := make(map[int64][]uint16)

	br := bufio.NewReader(rh)

	for {
		str, e := br.ReadString('\n')
		if e != nil {
			break
		}

		toks := token.Make(str)
		if len(toks) == 0 {
			continue
		}

		code := int64(0)
		tgs := []uint16{}

		for _, v := range toks {

			if len(v) > 2 && v[0] == '@' {
				tc := tags.ToCode(v[1:])
				if tc != 0 {
					tgs = append(tgs, tc)
				}
				continue
			}

			code = types.Append(code, types.ToCode(v))
		}

		if code == 0 {
			continue
		}

		result[code] = tgs
	}

	tab = result

	log.Info(fmt.Sprintf("corpus schema reloaded %.4fs", tm.DeltaFloat()))
	log.Info(fmt.Sprintf("corpus schema size = %d", Total()))
}

func Total() int {
	return len(tab)
}
