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

func (m *Matrix) Trans() *Matrix {
	r := New(m.cols, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r.data[j*r.cols+i] = m.data[i*m.cols+j]
		}
	}

	return r
}
