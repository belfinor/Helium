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
	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/statements"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/corpus/words"
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
	slang := int(0)

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

		if ws == nil {
			return
		}

		if ws.HasOpt(opts.Opt(opts.OPT_EOS)) {
			st.Tact()
			return
		}

		ws.ForEachTags(addTag)
	}

	// read forms stream
	for ws := range makeGroupStream(wordStream, &slang) {

		if buf.Len() >= bufSize {
			procBuf()
		}

		buf.PushBack(ws)
	}

	for buf.Len() > 0 {
		procBuf()
	}

	result := st.Finish()
	res := []string{}

	for _, v := range result {
		val := tags.FromCode(v)
		if val != "" {
			res = append(res, val)
		}
	}

	return res, slang > 0
}

// get *index.Record from *list.Element
func wsFromList(e *list.Element) *index.Record {

	if e == nil {
		return nil
	}

	return index.Get(e.Value.(string))
}

// original word stream with agg words to phrases
func makeGroupStream(input <-chan string, slang *int) <-chan *index.Record {

	output := make(chan *index.Record, 2048)

	go func() {

		frms := forms.New(1024)
		dot := &index.Record{Words: []*index.Word{words.Parse("eos .", frms)}, Name: "."}

		typeName := types.ToCode("имя")
		typeSlang := types.ToCode("мат")
		//typePatr := types.ToCode("отчество")
		typeLName := types.ToCode("фамилия")
		typeHuman := types.ToCode("человек")

		bufSize := 3
		buf := list.New()

		builder := strings.Builder{}

		tryAgg := func() {

			first := buf.Front()

			ws1 := wsFromList(first)
			if ws1 == nil {
				output <- nil
				buf.Remove(first)
				return
			}

			second := first.Next()
			ws2 := wsFromList(second)

			if ws2 == nil {
				output <- ws1
				buf.Remove(first)
				return
			}

			if ws1.HasType(typeName) && ws2.HasType(typeLName) {

				for _, nt := range []int32{opts.OPT_MR | opts.OPT_NOUN | opts.OPT_RU, opts.OPT_GR | opts.OPT_NOUN | opts.OPT_RU} {

					w1 := ws1.WordByOpt(opts.Opt(nt))

					if w1 != nil {

						w2 := ws2.WordByOpt(opts.Opt(nt))
						if w2 != nil {

							w3 := words.NounNoun(frms, w1, w2, opts.Opt(nt|opts.OPT_ALIVE))
							w3.AddType(typeHuman)

							ws := &index.Record{
								Name:  first.Value.(string) + " " + second.Value.(string),
								Words: []*words.Word{w3},
							}

							output <- ws
							buf.Remove(buf.Front())
							buf.Remove(buf.Front())
							return
						}

					}
				}

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

			for i := buf.Len(); i > 1; i-- {

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

					if rec.HasType(typeSlang) {
						*slang = *slang + 1
					}

					for j := 0; j < i; j++ {
						buf.Remove(buf.Front())
					}

					output <- rec
					return
				}
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
