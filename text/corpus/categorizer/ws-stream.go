package categorizer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-10

import (
	"container/list"

	"github.com/belfinor/Helium/text/corpus/forms"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/corpus/words"
)

// original word stream with agg words to phrases
func wsStream(input <-chan string, slang *int) <-chan *index.Record {

	output := make(chan *index.Record, 2048)

	go func() {

		frms := forms.New(1024)
		dot := &index.Record{Words: []*index.Word{words.Parse("eos .", frms)}, Name: "."}

		bufSize := 3
		buf := list.New()

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

			s1 := ws1
			s2 := ws2

			if ws1.HasType(types.TP_LASTNAME) && ws2.HasType(types.TP_NAME) {
				s1 = ws2
				s2 = ws1
			}

			if s1.HasType(types.TP_NAME) && s2.HasType(types.TP_LASTNAME) {

				if s1 == ws2 {
					ws3 := wsFromList(second.Next())

					if ws3 != nil && ws3.HasType(types.TP_PATRONYMIC) {
						for _, nt := range []int32{opts.OPT_MR | opts.OPT_NOUN | opts.OPT_RU, opts.OPT_GR | opts.OPT_NOUN | opts.OPT_RU} {

							w1 := ws1.WordByOpt(opts.Opt(nt))

							if w1 != nil {

								w2 := ws2.WordByOpt(opts.Opt(nt))
								if w2 != nil {

									w3 := ws3.WordByOpt(opts.Opt(nt))
									if w3 != nil {

										wr := words.NounNounNoun(frms, w2, w3, w1, w1.GetOpt())

										if ws := index.Get(wr.Form(0)); ws != nil {
											ws.TT2Word(wr)
										} else {
											wr.AddType(types.TP_MAN)
										}

										ws := &index.Record{
											Name:  ws2.Name + " " + ws3.Name + " " + ws1.Name,
											Words: []*words.Word{wr},
										}

										output <- ws
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										return
									}
								}
							}
						}
					}
				}

				for _, nt := range []int32{opts.OPT_MR | opts.OPT_NOUN | opts.OPT_RU, opts.OPT_GR | opts.OPT_NOUN | opts.OPT_RU} {

					w1 := s1.WordByOpt(opts.Opt(nt))

					if w1 != nil {

						w2 := s2.WordByOpt(opts.Opt(nt))
						if w2 != nil {

							w3 := words.NounNoun(frms, w1, w2, w2.GetOpt())

							if ws := index.Get(w3.Form(0)); ws != nil {
								ws.TT2Word(w3)
							} else {
								w3.AddType(types.TP_MAN)
							}

							ws := &index.Record{
								Name:  s1.Name + " " + s2.Name,
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

			if ws1.HasType(types.TP_NAME) && ws2.HasType(types.TP_PATRONYMIC) {

				ws3 := wsFromList(second.Next())

				if ws3 != nil && ws3.HasType(types.TP_LASTNAME) {
					for _, nt := range []int32{opts.OPT_MR | opts.OPT_NOUN | opts.OPT_RU, opts.OPT_GR | opts.OPT_NOUN | opts.OPT_RU} {

						w1 := ws1.WordByOpt(opts.Opt(nt))

						if w1 != nil {

							w2 := ws2.WordByOpt(opts.Opt(nt))
							if w2 != nil {

								w3 := ws3.WordByOpt(opts.Opt(nt))
								if w3 != nil {

									wr := words.NounNounNoun(frms, w1, w2, w3, w3.GetOpt())

									if ws4 := index.Get(wr.Form(0)); ws4 != nil {
										ws3.TT2Word(wr)
									} else {
										wr.AddType(types.TP_MAN)
									}

									ws := &index.Record{
										Name:  ws1.Name + " " + ws2.Name + " " + ws3.Name,
										Words: []*words.Word{wr},
									}

									output <- ws
									buf.Remove(buf.Front())
									buf.Remove(buf.Front())
									buf.Remove(buf.Front())
									return
								}
							}
						}
					}
				}

				for _, nt := range []int32{opts.OPT_MR | opts.OPT_NOUN | opts.OPT_RU, opts.OPT_GR | opts.OPT_NOUN | opts.OPT_RU} {

					w1 := s1.WordByOpt(opts.Opt(nt))

					if w1 != nil {

						w2 := s2.WordByOpt(opts.Opt(nt))
						if w2 != nil {

							w3 := words.NounNoun(frms, w1, w2, w2.GetOpt())

							if ws3 := index.Get(w3.Form(0)); ws3 != nil {
								ws3.TT2Word(w3)
							} else {
								w3.AddType(types.TP_MAN)
							}

							ws := &index.Record{
								Name:  ws1.Name + " " + ws2.Name,
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

			// check adj + noun from predef
			if ws1.HasOpt(opts.Opt(opts.OPT_ADJ)) && ws2.HasOpt(opts.Opt(opts.OPT_NOUN)) {

				for _, v := range []int32{opts.OPT_MR, opts.OPT_GR, opts.OPT_SR, opts.OPT_ML} {

					w1 := ws1.WordByOpt(opts.Opt(v | opts.OPT_RU | opts.OPT_ADJ))
					if w1 == nil {
						continue
					}

					w2 := ws2.WordByOpt(opts.Opt(v | opts.OPT_RU | opts.OPT_NOUN))
					if w2 == nil {
						continue
					}

					title := w1.Form(0) + " " + w2.Form(0)
					ws3 := index.Get(title)

					if ws3 != nil {
						w3 := words.AdjNoun(frms, w1, w2, w2.GetOpt())
						w3.CloneTT(ws3.Words[0])

						ws3 := &index.Record{
							Name:  ws1.Name + " " + ws2.Name,
							Words: []*words.Word{w3},
						}

						output <- ws3
						buf.Remove(buf.Front())
						buf.Remove(buf.Front())
						return
					}
				}
			}

			// noun + noun / noun + adj + noun / noun + pretext + noun
			if ws1.HasOpt(opts.Opt(opts.OPT_NOUN)) && ws2.HasOpt(opts.Opt(opts.OPT_NOUN)) {

				ws3 := wsFromList(second.Next())

				lst := []string{ws2.Name}

				if ws3 != nil {
					lst = append(lst, ws3.Name)
				}

				if len(lst) > 1 {
					lst[0], lst[1] = lst[0]+" "+lst[1], lst[0]
				}

				for i, str := range lst {

					for _, w := range ws1.Words {

						title := w.Form(0) + " " + str

						ws := index.Get(title)
						if ws == nil {
							continue
						}

						wr := words.NounStr(frms, w, str, w.GetOpt())

						ws.TT2Word(wr)

						ws3 := &index.Record{
							Name:  ws1.Name + " " + str,
							Words: []*words.Word{wr},
						}

						output <- ws3
						buf.Remove(buf.Front())
						buf.Remove(buf.Front())
						if i == 0 {
							buf.Remove(buf.Front())
						}

						return
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
