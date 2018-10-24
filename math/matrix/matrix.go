package matrix

import "errors"

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-24

type Matrix struct {
	data []float64
	rows int
	cols int
}

func New(rows int, cols int) (*Matrix, error) {

	if rows <= 0 || cols <= 0 {
		return nil, errors.New("invalid matrix size")
	}

	return &Matrix{
		data: make([]float64, rows*cols),
		rows: rows,
		cols: cols,
	}, nil
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

func Identity(rows int) (*Matrix, error) {
	m, err := New(rows, rows)
	if err != nil {
		return nil, err
	}

	for i := 0; i < rows; i++ {
		m.data[i*rows+i] = 1
	}

	return m, nil
}

func (m *Matrix) Export() [][]float64 {
	res := make([][]float64, m.rows)

	for i, _ := range res {
		res[i] = make([]float64, m.cols)
		copy(res[i], m.data[i*m.rows:])
	}

	return res
}

func Import(m [][]float64) (*Matrix, error) {
	if m == nil || len(m) == 0 {
		return nil, errors.New("invalid input data")
	}

	rows := len(m)
	cols := len(m[0])

	r, e := New(rows, cols)
	if e != nil {
		return nil, e
	}

	k := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r.data[k] = m[i][j]
			k++
		}
	}

	return r, nil
}

func Sum(m1, m2 *Matrix) (*Matrix, error) {
	if m1 == nil || m2 == nil || m1.rows != m2.rows || m1.cols != m2.cols {
		return nil, errors.New("matrix.Sum icompatble matrix")
	}

	r, e := New(m1.rows, m2.cols)
	if e != nil {
		return nil, e
	}

	for i, v := range m1.data {
		r.data[i] = v + m2.data[i]
	}

	return r, nil
}

func Sub(m1, m2 *Matrix) (*Matrix, error) {
	if m1 == nil || m2 == nil || m1.rows != m2.rows || m1.cols != m2.cols {
		return nil, errors.New("matrix.Sub incompatible matrix")
	}

	r, e := New(m1.rows, m2.cols)
	if e != nil {
		return nil, e
	}

	for i, v := range m1.data {
		r.data[i] = v - m2.data[i]
	}

	return r, nil
}

func Mul(m1, m2 *Matrix) (*Matrix, error) {

	if m1 == nil || m2 == nil || m1.cols != m2.rows {
		return nil, errors.New("matrix.Mul incompatible matrix")
	}

	r, e := New(m1.rows, m2.cols)
	if e != nil {
		return nil, e
	}

	for i := 0; i < r.rows; i++ {
		for j := 0; j < r.cols; j++ {
			v := float64(0)
			for k := 0; k < m1.cols; k++ {
				v += m1.data[i*m1.rows+k] * m2.data[k*m2.rows+j]
			}
			r.data[i*r.rows+j] = v
		}
	}

	return r, nil
}

func (m *Matrix) Trans() *Matrix {
	r, _ := New(m.cols, m.rows)

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

func (m *Matrix) Determinant() (float64, error) {

	if m.rows != m.cols {
		return 0, errors.New("matrix.Determinant rows != cols")
	}

	res := float64(0)

	mt := m.Clone()

	for i := 0; i < mt.rows; i++ {

		cur := mt.data[i*mt.cols+i]

		if cur == 0 {

			for j := i + 1; j < mt.rows; j++ {
				cur = mt.data[j*mt.cols+i]
				if cur != 0 {
					mt.swapRows(i, j)
					break
				}
			}

			if cur == 0 {
				return 0, nil
			}

			res *= -1
		}

		res *= cur

		for j := i + 1; j < mt.rows; j++ {

			k1 := i*mt.cols + i
			k2 := j*mt.cols + i

			coeff := mt.data[(i+1)*mt.cols+i] / cur

			for k := i; k < mt.cols; k++ {
				mt.data[k2] -= mt.data[k1] * coeff
				k1++
				k2++
			}
		}
	}

	return res, nil
}

func (m *Matrix) Set(i, j int, v float64) error {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return errors.New("matrix.Set invalid cell")
	}

	m.data[i*m.cols+j] = v
	return nil
}

func (m *Matrix) Inverse() (*Matrix, error) {

	if m.rows != m.cols {
		return nil, errors.New("matrix.Inverse inverse matrix does not exists")
	}

	l := m.Clone()

	r, e := Identity(m.rows)
	if e != nil {
		return nil, e
	}

	for i := 0; i < l.rows; i++ {

		cur := l.data[i*l.cols+i]

		if cur == 0 {

			for j := i + 1; j < l.rows; j++ {
				cur = l.data[j*l.cols+i]
				if cur != 0 {
					l.swapRows(i, j)
					r.swapRows(i, j)
					break
				}
			}

			if cur == 0 {
				return nil, errors.New("matrix.Inverse inverse matrix does not exists")
			}
		}

		for j := i + 1; j < l.rows; j++ {
			coeff := l.data[j*l.cols+i] / cur

			if coeff == 0 {
				continue
			}

			k1 := i * l.cols
			k2 := j * l.cols

			for k := 0; k < l.cols; k++ {
				l.data[k2] -= l.data[k1] * coeff
				r.data[k2] -= r.data[k1] * coeff
				k1++
				k2++
			}
		}
	}

	for i := l.rows - 1; i >= 0; i-- {
		div := l.data[i*l.cols+i]

		k := i * l.cols

		for j := 0; j < l.cols; j++ {
			r.data[k] /= div
			k++
		}

		for j := i - 1; j >= 0; j-- {
			coeff := l.data[j*l.cols+i]
			k1 := i * l.cols
			k2 := j * l.cols
			for n := 0; n < l.cols; n++ {
				r.data[k2] -= r.data[k1] * coeff
				k1++
				k2++
			}
		}
	}

	return r, nil
}
