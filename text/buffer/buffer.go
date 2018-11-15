package buffer

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-11-15

type Buffer struct {
	data  []string
	size  int
	first int
	limit int
}

func NewBuffer(length int) *Buffer {
	return &Buffer{
		data:  make([]string, length),
		size:  0,
		first: 0,
		limit: length,
	}
}

func (b *Buffer) Empty() bool {
	return b.size == 0
}

func (b *Buffer) Full() bool {
	return b.size == b.limit
}

func (b *Buffer) Get(index int) string {
	return b.data[(b.first+index)%b.limit]
}

func (b *Buffer) Shift(num int) {
	if num <= 0 {
		return
	}

	if num >= b.size {
		b.first = 0
		b.size = 0
	} else {
		b.first = (b.first + num) % b.limit
		b.size -= num
	}
}

func (b *Buffer) Pop(num int) {
	if num <= 0 {
		return
	}

	if num >= b.size {
		b.first = 0
		b.size = 0
	} else {
		b.size -= num
	}
}

func (b *Buffer) Size() int {
	return b.size
}

func (b *Buffer) Add(str string) {
	b.data[(b.first+b.size)%b.limit] = str
	if b.size < b.limit {
		b.size++
	}
}
