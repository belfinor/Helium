package statements

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-06

import (
	"math"

	"github.com/belfinor/Helium/slice"
)

// statements agg struct
type Statements struct {
	cur   map[uint16]bool
	tags  map[uint16]int
	verts map[uint16]int
	ribs  map[uint32]bool
}

// make new Statement object
func New() *Statements {

	return &Statements{
		cur:   make(map[uint16]bool),
		tags:  make(map[uint16]int),
		verts: make(map[uint16]int),
		ribs:  make(map[uint32]bool),
	}
}

// add code to current statement
func (s *Statements) Add(code uint16) {
	if code != 0 {
		s.cur[code] = true
	}
}

// handler on dot
func (s *Statements) Tact() {

	for v1 := range s.cur {
		s.tags[v1]++

		for v2 := range s.cur {

			if v1 < v2 {

				rib := (uint32(v1) << 16) + uint32(v2)
				if !s.ribs[rib] {
					s.ribs[rib] = true

					s.verts[v1]++
					s.verts[v2]++
				}

			}

		}
	}

	s.cur = map[uint16]bool{}
}

var empty []uint16 = []uint16{}

// get result categories
func (s *Statements) Finish() []uint16 {

	maxCnt := 0

	for _, v := range s.tags {
		if v > maxCnt {
			maxCnt = v
		}
	}

	if maxCnt == 0 {
		return empty
	}

	rank := make(map[uint16]float64)

	max := float64(1)

	for k, v := range s.tags {

		if maxCnt > 1 && v == 1 {
			continue
		}

		r := float64(v) * (1 + math.Log10(float64(s.verts[k]+1)))
		rank[k] = r
		max = math.Max(max, r)
	}

	return slice.FromMapCnt(rank, &slice.MapCntOpts{Limit: 10, MinVal: max / 2}).([]uint16)
}
