package matrix

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-10-19

type Matrix struct {
	data []float64
	cols int
	rows int
}

func New(rows, cols int) *Matrix {

	if rows < 1 || cols < 1 {
		return nil
	}

	return &Matrix{
		data: make([]float64, rows*cols),
		cols: cols,
		rows: rows,
	}
}

func Identity(size int) *Matrix {
	m := New(size, size)

	for i := 0; i < size; i++ {
		m.data[i*size+i] = 1
	}

	return m
}

func FromSlice2(data [][]float64) *Matrix {
	if data == nil {
		return nil
	}

	rows := len(data)

	if rows == 0 {
		return nil
	}

	cols := len(data[0])

	r := New(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r.data[i*cols+j] = data[i][j]
		}
	}

	return r
}

func (m *Matrix) Clone() *Matrix {

	r := &Matrix{
		data: make([]float64, m.cols*m.rows),
		cols: m.cols,
		rows: m.rows,
	}

	copy(r.data, r.data)

	return r
}

func Add(m1, m2 *Matrix) *Matrix {
	if m1.cols != m2.cols || m1.rows != m2.rows {
		return nil
	}

	r := New(m1.cols, m2.rows)

	for i, v := range m1.data {
		r.data[i] = v + m2.data[i]
	}

	return r
}

func Sub(m1, m2 *Matrix) *Matrix {
	if m1.cols != m2.cols || m1.rows != m2.rows {
		return nil
	}

	r := New(m1.cols, m2.rows)

	for i, v := range m1.data {
		r.data[i] = v - m2.data[i]
	}

	return r
}

func Mul(m1, m2 *Matrix) *Matrix {
	if m1.cols != m2.rows {
		return nil
	}

	r := New(m1.rows, m2.cols)

	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m2.cols; j++ {
			pos := i*r.cols + j
			for k := 0; k < m1.cols; k++ {
				r.data[pos] += m1.data[i*m1.cols+k] * m2.data[k*m2.cols+j]
			}
		}
	}

	return r
}

func MulNum(m *Matrix, num float64) *Matrix {
	r := New(m.rows, m.cols)

	for i, v := range m.data {
		r.data[i] = v * num
	}

	return r
}

func (m *Matrix) Trans() *Matrix {
	r := New(m.cols, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r.data[j*r.cols+i] = m.data[i*m.cols+j]
		}
	}

	return r
}

func (m *Matrix) Slice2() [][]float64 {

	res := make([][]float64, m.rows)

	for i := 0; i < m.rows; i++ {

		res[i] = make([]float64, m.cols)

		for j := 0; j < m.cols; j++ {
			res[i][j] = m.data[i*m.cols+j]
		}
	}

	return res
}

func (m *Matrix) Get(i, j int) float64 {
	return m.data[i*m.cols+j]
}

func (m *Matrix) Set(i, j int, v float64) {
	m.data[i*m.cols+j] = v
}

func (m *Matrix) Row() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Det() float64 {

	r := m.Clone()

	res := float64(1)

	for i := 0; i < r.cols; i++ {
		if r.data[i*r.cols+i] == 0 {
			ok := false
			for k := i + 1; k < r.rows; k++ {
				if r.data[k*r.cols+i] != 0 {
					ok = true
					r.swapRows(i, k)
					res *= -1
					break
				}
			}
			if !ok {
				return 0
			}
		}

		res *= r.data[i*r.cols+i]

		for k := i + 1; k < r.rows; k++ {
			v := r.data[k*r.cols+i] / r.data[i*r.cols+i]
			for j := i; j < r.cols; j++ {
				r.data[k*r.cols+j] -= v * r.data[i*r.cols+j]
			}
		}
	}

	return res
}

func (m *Matrix) swapRows(i1, i2 int) {
	for j := 0; j < m.cols; j++ {
		m.data[i1*m.cols+j], m.data[i2*m.cols+j] = m.data[i2*m.cols+j], m.data[i1*m.cols+j]
	}
}
