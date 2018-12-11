package schemas

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-11

import (
	"container/list"
	"fmt"

	"github.com/belfinor/Helium/text/corpus/categorizer/statements"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/tags"
	"github.com/belfinor/Helium/text/corpus/types"
)

var tab map[int64][]uint16 = map[int64][]uint16{
	types.FromList(types.TP_FLY, types.TP_TO, types.TP_CITY):                      []uint16{tags.ToCode("путешествия")},
	types.FromList(types.TP_FLY, types.TP_TO, types.TP_COUNTRY):                   []uint16{tags.ToCode("путешествия")},
	types.FromList(types.TP_FLY, types.TP_ON, types.TP_PLANET):                    []uint16{tags.ToCode("космос")},
	types.FromList(types.TP_FLY, types.TP_TO, types.TP_PLANET):                    []uint16{tags.ToCode("космос")},
	types.FromList(types.TP_POLITIC, types.TP_FLY, types.TP_TO, types.TP_CITY):    []uint16{tags.ToCode("политика")},
	types.FromList(types.TP_POLITIC, types.TP_FLY, types.TP_TO, types.TP_COUNTRY): []uint16{tags.ToCode("политика")},
}

func Proc(buf *list.List, st *statements.Statements) int {

	i := 0
	j := 0

	checks := make([]int64, 0, 32)

	for e := buf.Front(); e != nil && i < 4; e = e.Next() {

		j++

		ws := e.Value.(*index.Record)
		if ws == nil {
			if i < 2 {
				return 0
			}
			j--
			break
		}

		data := make([]int64, 0, 8)

		ws.ForEachType(func(t uint16) {
			data = append(data, int64(t))
		})

		if len(data) == 0 {

			if ws.HasOpt(opts.Opt(opts.OPT_ADJ)) && i > 0 {
				continue
			}

			if i < 2 {
				return 0
			}

			j--
			break
		}

		if len(checks) == 0 {
			checks = data
		} else {

			res := make([]int64, 0, (len(checks)+1)*len(data))
			res = append(res, checks...)

			for _, v1 := range checks {
				for _, v2 := range data {
					res = append(res, types.AppendCode(v1, uint16(v2)))
				}
			}

			checks = res
		}

		i++
	}

	var last []uint16

	for _, t := range checks {
		if lst, h := tab[t]; h {
			last = lst
		}
	}

	if len(last) > 0 {

		for _, t := range last {
			st.Add(t)
		}

		str := ""

		for j > 0 {

			if str != "" {
				str += " "
			}

			str += buf.Front().Value.(*index.Record).Name

			buf.Remove(buf.Front())
			j--

		}

		fmt.Println(str)

		return j
	}

	return 0
}
