package categorizer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-14

import (
	"container/list"
	"strconv"

	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/tools"
	"github.com/belfinor/Helium/text/corpus/types"
	"github.com/belfinor/Helium/text/corpus/words"
)

func tryFind(pref string, e *list.Element, level int) (*index.Record, int) {

	if level > 4 || e == nil {
		return nil, 0
	}

	ws := tools.WsFromList(e)
	if ws == nil {
		return nil, 0
	}

	var res *index.Record
	var maxLevel int

	upRes := func(rec *index.Record, level int) {
		if rec != nil && level > maxLevel {
			res, maxLevel = rec, level
		}
	}

	if level > 1 {

		r := index.Get(pref + " " + ws.Name)
		if r != nil {
			upRes(r, level)
		} else {

			for _, w := range ws.Words {
				r := index.Get(pref + " " + w.Base)
				if r != nil {
					upRes(r, level)
					break
				}
			}

		}

	}

	if level == 1 {
		upRes(tryFind(ws.Name, e.Next(), level+1))
	} else {
		upRes(tryFind(pref+" "+ws.Name, e.Next(), level+1))
	}

	for _, w := range ws.Words {

		if level == 1 {
			upRes(tryFind(w.Base, e.Next(), level+1))
		} else {
			upRes(tryFind(pref+" "+w.Base, e.Next(), level+1))
		}

	}

	return res, maxLevel
}

var wsAgreed []int64 = []int64{opts.OPT_MR, opts.OPT_GR, opts.OPT_EN}

// original word stream with agg words to phrases
func wsStream(input <-chan string, slang *int) <-chan *index.Record {

	output := make(chan *index.Record, 2048)

	go func() {

		dot := &index.Record{Words: []*index.Word{words.Parse("eos . %.", func(f string, w *index.Word) {})}, Name: "."}

		bufSize := 4
		buf := list.New()

		tryAgg := func() {

			first := buf.Front()

			res, level := tryFind("", first, 1)

			if res != nil && level > 1 {

				if res.TypeMask&types.MTP_SKIP == 0 {
					output <- res
				}

				for level > 0 {
					buf.Remove(buf.Front())
					level--
				}
				return
			}

			ws1 := tools.WsFromList(first)
			if ws1 == nil {
				output <- nil
				buf.Remove(first)
				return
			}

			if ws1.TypeMask&types.MTP_SKIP != 0 {
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

			if ws1.TypeMask&types.MTP_NAME != 0 {

				// иван иванов
				if ws2.TypeMask&types.MTP_LASTNAME != 0 {

					for _, t := range wsAgreed {
						w1 := ws1.WordByOpt(t)
						if w1 != nil {
							if w2 := ws2.WordByOpt(t); w2 != nil {

								wr := words.Join(w1.GetOpt(), w1, w2)

								wr.AddType(types.TP_MAN)

								wsr := &index.Record{
									Name:     ws1.Name + " " + ws2.Name,
									Words:    []*words.Word{wr},
									TypeMask: wr.TypeMask(),
								}

								output <- wsr
								buf.Remove(buf.Front())
								buf.Remove(buf.Front())
								return
							}
						}
					}
				}

				// иван иванович
				if ws2.TypeMask&types.MTP_PATRONYMIC != 0 {

					// try иван иванович иванов
					ws3 := tools.WsFromList(second.Next())
					if ws3 != nil && ws3.TypeMask&types.MTP_LASTNAME != 0 {

						for _, t := range wsAgreed {
							w1 := ws1.WordByOpt(t)
							if w1 != nil {
								if w2 := ws2.WordByOpt(t); w2 != nil {
									if w3 := ws3.WordByOpt(t); w3 != nil {
										wr := words.Join(w1.GetOpt(), w1, w2, w3)

										wr.AddType(types.TP_MAN)

										wsr := &index.Record{
											Name:     ws1.Name + " " + ws2.Name + " " + ws3.Name,
											Words:    []*words.Word{wr},
											TypeMask: wr.TypeMask(),
										}

										output <- wsr
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										return
									}
								}
							}
						}

					}

					for _, t := range wsAgreed {
						w1 := ws1.WordByOpt(t)
						if w1 != nil {
							if w2 := ws2.WordByOpt(t); w2 != nil {

								wr := words.Join(w1.GetOpt(), w1, w2)

								wr.AddType(types.TP_MAN)

								wsr := &index.Record{
									Name:     ws1.Name + " " + ws2.Name,
									Words:    []*words.Word{wr},
									TypeMask: wr.TypeMask(),
								}

								output <- wsr
								buf.Remove(buf.Front())
								buf.Remove(buf.Front())
								return
							}
						}
					}
				}

				// петр i
				if ws2.TypeMask&types.MTP_ROMAN != 0 {

					w1 := ws1.WordByMask(types.MTP_NAME)

					wr := words.Join(w1.GetOpt(), w1, ws2.Words[0])

					wr.AddTag(tags.ToCode("история"))
					wr.AddType(types.TP_HPERSON)

					wsr := &index.Record{
						Name:     ws1.Name + " " + ws2.Name,
						Words:    []*words.Word{wr},
						TypeMask: wr.TypeMask(),
					}

					output <- wsr
					buf.Remove(buf.Front())
					buf.Remove(buf.Front())
					return
				}

			}

			if ws1.TypeMask&types.MTP_LASTNAME != 0 {

				// иванов иван -> иван иванов
				if ws2.TypeMask&types.MTP_NAME != 0 {

					// try иванов иван иванович -> иван иванович иванов
					ws3 := tools.WsFromList(second.Next())
					if ws3 != nil && ws3.TypeMask&types.MTP_PATRONYMIC != 0 {

						for _, t := range wsAgreed {
							w1 := ws1.WordByOpt(t)
							if w1 != nil {
								if w2 := ws2.WordByOpt(t); w2 != nil {
									if w3 := ws3.WordByOpt(t); w3 != nil {
										wr := words.Join(w1.GetOpt(), w2, w3, w1)

										wr.AddType(types.TP_MAN)

										wsr := &index.Record{
											Name:     ws2.Name + " " + ws3.Name + " " + ws1.Name,
											Words:    []*words.Word{wr},
											TypeMask: wr.TypeMask(),
										}

										output <- wsr
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										buf.Remove(buf.Front())
										return
									}
								}
							}
						}

					}

					for _, t := range wsAgreed {
						w1 := ws1.WordByOpt(t)
						if w1 != nil {
							if w2 := ws2.WordByOpt(t); w2 != nil {

								wr := words.Join(w1.GetOpt(), w2, w1)

								wr.AddType(types.TP_MAN)

								wsr := &index.Record{
									Name:     ws2.Name + " " + ws1.Name,
									Words:    []*words.Word{wr},
									TypeMask: wr.TypeMask(),
								}

								output <- wsr
								buf.Remove(buf.Front())
								buf.Remove(buf.Front())
								return
							}
						}
					}

				}

			}

			if ws1.TypeMask&types.MTP_NUMBER != 0 {

				v, _ := strconv.Atoi(ws1.Name)

				if v > 300 && v < 2000 {

					if ws2.TypeMask&types.MTP_YEAR != 0 {
						wr := words.Join(0, ws1.Words[0], ws2.Words[0])

						wr.AddType(types.TP_HDATE)
						wr.AddTag(tags.ToCode("история"))

						wsr := &index.Record{
							Name:     ws1.Name + " " + ws2.Name,
							Words:    []*words.Word{wr},
							TypeMask: wr.TypeMask(),
						}

						output <- wsr
						buf.Remove(buf.Front())
						buf.Remove(buf.Front())
						return
					}

				} else if v > 0 && v < 21 {

					if ws2.TypeMask&types.MTP_CENTURY != 0 {
						wr := words.Join(0, ws1.Words[0], ws2.Words[0])

						wr.AddType(types.TP_HDATE)
						wr.AddTag(tags.ToCode("история"))

						wsr := &index.Record{
							Name:     ws1.Name + " " + ws2.Name,
							Words:    []*words.Word{wr},
							TypeMask: wr.TypeMask(),
						}

						output <- wsr
						buf.Remove(buf.Front())
						buf.Remove(buf.Front())
						return
					}
				}
			}

			if ws1.TypeMask&types.MTP_ROMAN != 0 {

				for _, w := range ws2.Words {
					if w.Base == "век" {
						wr := words.Join(0, ws1.Words[0], w)

						wr.AddType(types.TP_HDATE)
						wr.AddTag(tags.ToCode("история"))

						wsr := &index.Record{
							Name:     ws1.Name + " " + ws2.Name,
							Words:    []*words.Word{wr},
							TypeMask: wr.TypeMask(),
						}

						output <- wsr
						buf.Remove(buf.Front())
						buf.Remove(buf.Front())
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
