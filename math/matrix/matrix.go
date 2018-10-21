package matrix

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-10-21

type Matrix struct {
	data []float64
	rows int
	cols int
}

func New(rows int, cols int) *Matrix {

	if rows <= 0 || cols <= 0 {
		return nil
	}

	return &Matrix{
		data: make([]float64, rows*cols),
		rows: rows,
		cols: cols,
	}
}

func (m *Matrix) Clone() *Matrix {
	r := &Matrix{
		data: make([]float64, m.rows*m.cols),
		rows: m.rows,
		cols: m.cols,
	}

	copy(r.data, m.data)

	return r
}

func Identity(rows int) *Matrix {
	m := New(rows, rows)
	if m == nil {
		return nil
	}

	for i := 0; i < rows; i++ {
		m.data[i*rows+i] = 1
	}

	return m
}

func (m *Matrix) Export() [][]float64 {
	res := make([][]float64, m.rows)

	for i, _ := range res {
		res[i] = make([]float64, m.cols)
		copy(res[i], m.data[i*m.rows:])
	}

	return res
}

func Import(m [][]float64) *Matrix {
	if m == nil || len(m) == 0 {
		return nil
	}

	rows := len(m)
	cols := len(m[0])

	r := New(rows, cols)

	k := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r.data[k] = m[i][j]
			k++
		}
	}

	return r
}

func Sum(m1, m2 *Matrix) *Matrix {
	if m1 == nil || m2 == nil || m1.rows != m2.rows || m1.cols != m2.cols {
		return nil
	}

	r := New(m1.rows, m2.cols)

	for i, v := range m1.data {
		r.data[i] = v + m2.data[i]
	}

	return r
}

func Sub(m1, m2 *Matrix) *Matrix {
	if m1 == nil || m2 == nil || m1.rows != m2.rows || m1.cols != m2.cols {
		return nil
	}

	r := New(m1.rows, m2.cols)

	for i, v := range m1.data {
		r.data[i] = v - m2.data[i]
	}

	return r
}

func Mul(m1, m2 *Matrix) *Matrix {

	if m1 == nil || m2 == nil || m1.cols != m2.rows {
		return nil
	}

	r := New(m1.rows, m2.cols)

	for i := 0; i < r.rows; i++ {
		for j := 0; j < r.cols; j++ {
			v := float64(0)
			for k := 0; k < m1.cols; k++ {
				v += m1.data[i*m1.rows+k] * m2.data[k*m2.rows+j]
			}
			r.data[i*r.rows+j] = v
		}
	}

	return r
}

func (m *Matrix) Trans() *Matrix {
	r := New(m.cols, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			r.data[j*r.rows+i] = m.data[i*m.rows+j]
		}
	}

	return r
}

func (m *Matrix) swapRows(i1, i2 int) {

	k1 := i1 * m.rows
	k2 := i2 * m.rows

	for j := 0; j < m.cols; j++ {
		m.data[k1], m.data[k2] = m.data[k2], m.data[k1]
		k1++
		k2++
	}

}
