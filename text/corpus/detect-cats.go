package corpus

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-06

import (
	"fmt"
	"io"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/buffer"
	"github.com/belfinor/Helium/text/corpus/index"
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

	buf := buffer.New(32)
	st := statements.New()

	addTag := func(t uint16) {
		st.Add(t)
	}

	procBuf := func() {
		f := buf.Get(0)
		buf.Shift(1)

		if f == "." {
			st.Tact()
			return
		}

		ws := index.Get(f)
		if ws == nil {
			return
		}

		ws.ForEachTags(addTag)
	}

	// read forms stream
	for f := range makeFormStream(wordStream) {

		if buf.Full() {
			procBuf()
		}

		buf.Add(f)
	}

	for buf.Full() {
		procBuf()
	}

	return nil, false
}

// orgnal word stream with agg words to phrases
func makeFormStream(input <-chan string) <-chan string {

	output := make(chan string, 2048)

	go func() {

		buf := buffer.New(3)

		proc := func() {

			for i := buf.Size(); i > 1; i-- {
				str := buf.Join(" ", i)

				if index.Get(str) != nil {
					buf.Shift(i)
					output <- str
					return
				}
			}

			output <- buf.Get(0)
			buf.Shift(1)
		}

		for w := range input {
			if buf.Full() {
				proc()
			}

			buf.Add(w)
		}

		for buf.Full() {
			proc()
		}

		close(output)

	}()

	return output
}
