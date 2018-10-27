package matrix

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2018-10-27

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {

	identity, e := Identity(10)
	if e != nil {
		t.Fatal("Identity not work")
	}

	if identity.rows != 10 || identity.cols != 10 {
		t.Fatal("Identity wrong size")
	}

	m2d := identity.Export()

	if len(m2d) != 10 || len(m2d[0]) != 10 {
		t.Fatal("Export not work")
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i != j && m2d[i][j] != 0 {
				t.Fatal("Wrond idetity data")
			}

			if i == j && m2d[i][j] != 1 {
				t.Fatal("wrong identty data")
			}
		}
	}

	cl := identity.Clone()

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			val, e := cl.Get(i, j)
			if e != nil {
				t.Fatal("Clone not work")
			}

			if i == j && val != 1 || i != j && val != 0 {
				t.Fatal("Clone not work")
			}
		}
	}

	det, e1 := identity.Determinant()
	if e1 != nil || det != 1 {
		fmt.Println(det, e1)
		t.Fatal("Determinant not work")
	}

	cl.Set(0, 0, 5)
	cl.Set(2, 2, 3)
	cl.Set(4, 4, 2)
	cl.Set(5, 4, 2)

	det, e1 = cl.Determinant()
	if e1 != nil || det != 30 {
		fmt.Println(det, e1)
		t.Fatal("Determinant not work")
	}

	cl, e1 = Import([][]float64{
		[]float64{10, 1, 1},
		[]float64{2, 10, 1},
		[]float64{2, 2, 10},
	})

	if e != nil {
		t.Fatal("matrix.Import not work")
	}

	det, e1 = cl.Determinant()
	if e1 != nil {
		t.Fatal("Detrminant not work")
	}

	if int(det+0.005) != 946 {
		t.Fatal("Determinant wrong value")
	}

	inv, err := cl.Inverse()
	if err != nil {
		t.Fatal("Inverse not work")
	}

	data := [][]int{
		[]int{1035, -84, -95},
		[]int{-190, 1035, -84},
		[]int{-169, -190, 1035},
	}

	for i := 0; i < inv.rows; i++ {
		for j := 0; j < inv.cols; j++ {
			v, _ := inv.Get(i, j)
			iv := int(v * 10000)
			if iv != data[i][j] {
				t.Fatal("wrong inverse matrix")
			}
		}
	}
}
