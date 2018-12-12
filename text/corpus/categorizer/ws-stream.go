package categorizer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-12

import (
	"container/list"

	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/tools"
	"github.com/belfinor/Helium/text/corpus/words"
)

// original word stream with agg words to phrases
func wsStream(input <-chan string, slang *int) <-chan *index.Record {

	output := make(chan *index.Record, 2048)

	go func() {

		dot := &index.Record{Words: []*index.Word{words.Parse("eos . %.", func(f string, w *index.Word) {})}, Name: "."}

		bufSize := 3
		buf := list.New()

		tryAgg := func() {

			first := buf.Front()

			ws1 := tools.WsFromList(first)
			if ws1 == nil {
				output <- nil
				buf.Remove(first)
				return
			}

			second := first.Next()
			ws2 := tools.WsFromList(second)

			if ws2 == nil {
				output <- ws1
				buf.Remove(first)
				return
			}

			output <- ws1
			buf.Remove(first)
		}

		proc := func() {

			first := buf.Front().Value.(string)
			if first == "." {
				buf.Remove(buf.Front())
				output <- dot
				return
			}

			tryAgg()

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
