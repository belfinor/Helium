package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-07

import (
	"container/list"
	"fmt"
	"io"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/statements"
	"github.com/belfinor/Helium/time/timer"
)

// detect text categories and slang flag
func DetectCats(rh io.RuneReader) ([]string, bool) {

	// init calc process time
	tm := timer.New()
	defer func() {
		log.Debug(fmt.Sprintf("corpus.DetectCats time %.4fs", tm.DeltaFloat()))
	}()

	// make word stream
	wordStream := text.WordStream(rh, text.WSO_ENDS, text.WSO_NO_URLS, text.WSO_HASHTAG, text.WSO_NO_XXX_COLON)

	bufSize := 32
	buf := list.New()
	st := statements.New()

	addTag := func(t uint16) {
		st.Add(t)
	}

	procBuf := func() {

		if buf.Len() == 0 {
			return
		}

		v := buf.Front().Value
		buf.Remove(buf.Front())

		if v == nil {
			return
		}

		ws := v.(*index.Record)

		if ws.HasOpt(opts.Opt(opts.OPT_EOS)) {
			st.Tact()
			return
		}

		ws.ForEachTags(addTag)
	}

	// read forms stream
	for ws := range makeGroupStream(wordStream) {

		if buf.Len() >= bufSize {
			procBuf()
		}

		buf.PushBack(ws)
	}

	for buf.Len() > 0 {
		procBuf()
	}

	return nil, false
}

// original word stream with agg words to phrases
func makeGroupStream(input <-chan string) <-chan *index.Record {

	output := make(chan *index.Record, 2048)

	go func() {

		bufSize := 3
		buf := list.New()

		builder := strings.Builder{}

		proc := func() {

			for i := buf.Len(); i > 0; i-- {

				builder.Reset()
				rec := buf.Front()

				for j := 0; j < i; j++ {
					if j > 0 {
						builder.WriteRune(' ')
					}
					builder.WriteString(rec.Value.(string))
				}

				str := builder.String()

				if rec := index.Get(str); rec != nil {

					for j := 0; j < i; j++ {
						buf.Remove(buf.Front())
					}

					output <- rec
					return
				}
			}

			output <- nil
			buf.Remove(buf.Front())
		}

		for w := range input {
			if buf.Len() >= bufSize {
				proc()
			}

			buf.PushBack(w)
		}

		for buf.Len() > 0 {
			proc()
		}

		close(output)

	}()

	return output
}
