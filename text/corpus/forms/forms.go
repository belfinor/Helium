package forms

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-12-07

type RANGE_FUNC func(string)

type Forms struct {
	last int
	data []string
}

func New(alloc int) *Forms {
	if alloc < 0 {
		alloc = 0
	}

	s := &Forms{
		last: 0,
		data: make([]string, 0, alloc),
	}

	return s
}

func (s *Forms) Total() int {

	if s == nil {
		return 0
	}

	return s.last
}

func (s *Forms) Add(form string) {
	s.data = append(s.data, form)
	s.last++
}

func (s *Forms) Range(from, to int) []string {
	if s == nil {
		return nil
	}

	return s.data[from:to]
}

func (s *Forms) RangeFunc(from, to int, fn RANGE_FUNC) {
	if s == nil {
		return
	}

	for _, v := range s.data[from:to] {
		fn(v)
	}
}

func (s *Forms) Get(index int) string {
	if s == nil {
		return ""
	}

	return s.data[index]
}
