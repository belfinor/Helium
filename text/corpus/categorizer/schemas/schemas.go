package schemas

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-12

import (
	"container/list"
	"fmt"

	"github.com/belfinor/Helium/text/corpus/categorizer/statements"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/types"
)

var tab map[int64][]uint16 = map[int64][]uint16{}

func Proc(buf *list.List, st *statements.Statements, trace bool) int {

	i := 0
	j := 0

	checks := make([]int64, 0, 32)
	places := make(map[int64]int, 32)

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

			if ws.HasOpt(opts.OPT_ADJ) && i > 0 {
				continue
			}

			if i < 2 {
				return 0
			}

			j--
			break
		}

		i++

		if len(checks) == 0 {
			checks = data

			for _, v := range data {
				checks = append(checks, v)
				places[v] = j
			}

		} else {

			res := make([]int64, 0, (len(checks)+1)*len(data))
			res = append(res, checks...)

			for _, v1 := range checks {
				for _, v2 := range data {
					code := types.AppendCircle(v1, uint16(v2))
					places[code] = j
					res = append(res, code)
				}
			}

			checks = res
		}
	}

	var last []uint16
	var pl int

	for _, t := range checks {
		if lst, h := tab[t]; h {
			last = lst
			pl = places[t]
		}
	}

	if len(last) > 0 {

		for _, t := range last {
			st.Add(t)
		}

		j = pl

		if trace {

			str := ""

			for pl > 0 {

				if str != "" {
					str += " "
				}

				str += buf.Front().Value.(*index.Record).Name

				buf.Remove(buf.Front())
				pl--

			}

			fmt.Println(str)

		} else {

			for pl > 0 {
				buf.Remove(buf.Front())
				pl--

			}
		}

		return j
	}

	return 0
}
