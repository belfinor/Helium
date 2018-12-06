package forms

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-12-06

type RANGE_FUNC func(string)

var storage *Forms

type Forms struct {
	last int
	data []string
}

func New(alloc int, def bool) *Forms {
	if alloc < 0 {
		alloc = 0
	}

	s := &Forms{
		last: 0,
		data: make([]string, 0, alloc),
	}

	if def {
		storage = s
	}

	return s
}

func Default() *Forms {
	return storage
}

func SetDefault(s *Forms) {
	storage = s
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

func Range(from, to int) []string {
	return storage.Range(from, to)
}

func RangeFunc(from, to int, fn RANGE_FUNC) {
	storage.RangeFunc(from, to, fn)
}

func Get(index int) string {
	return storage.Get(index)
}

func Total() int {
	return storage.Total()
}
