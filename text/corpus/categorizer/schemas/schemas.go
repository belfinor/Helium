package schemas

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-12-15

import (
	"container/list"
	"fmt"

	"github.com/belfinor/Helium/text/corpus/categorizer/statements"
	"github.com/belfinor/Helium/text/corpus/index"
	"github.com/belfinor/Helium/text/corpus/opts"
	"github.com/belfinor/Helium/text/corpus/types"
)

var tab map[int64][]uint16 = map[int64][]uint16{}

func fetchNext(e *list.Element) (*index.Record, *list.Element) {
	e = e.Next()
	if e == nil {
		return nil, nil
	}

	ws := e.Value.(*index.Record)
	return ws, e
}

func killN(l *list.List, n int) {

	for n > 0 {
		l.Remove(l.Front())
		n--
	}
}

func cleanHow(buf *list.List) int {

	cnt := 0

	e := buf.Front()
	if e == nil {
		return 0
	}

	ws := e.Value.(*index.Record)
	if ws == nil {
		return 0
	}

	if ws.TypeMask&types.MTP_COMMA == 0 {
		return 0
	}

	cnt++

	ws, e = fetchNext(e)
	if ws == nil {
		return 0
	}

	if ws.TypeMask&types.MTP_HOW == 0 {
		return 0
	}

	cnt++

	state := 0

	for {
		ws, e = fetchNext(e)
		if ws == nil {
			break
		}

		fmt.Println("PRE: " + ws.Name)

		switch state {
		case 0:

			if ws.HasOpt(opts.OPT_NOUN) || ws.HasOpt(opts.OPT_ADJ) {
				state = 1
				cnt++
			} else if ws.HasOpt(opts.OPT_PRETEXT) {
				state = 3
				cnt++
			} else {
				return 0
			}

		case 1:

			if ws.TypeMask&types.MTP_DOT != 0 {
				cnt++
				killN(buf, cnt)
				return cnt
			}

			if ws.HasOpt(opts.OPT_NOUN) {
				cnt++
				state = 2
			} else {
				return 0
			}

		case 2:

			if ws.TypeMask&types.MTP_DOT != 0 {
				cnt++
				killN(buf, cnt)
				return cnt
			}

			return 0

		case 3:

			if ws.HasOpt(opts.OPT_NOUN) || ws.HasOpt(opts.OPT_ADJ) {
				state = 1
				cnt++
			} else {
				return 0
			}
		}
	}

	return 0
}

func Proc(buf *list.List, st *statements.Statements, trace bool) int {

	res := cleanHow(buf)
	if res > 0 {
		return res
	}

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
