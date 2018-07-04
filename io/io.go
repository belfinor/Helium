package io

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-04

import (
	i "io"
)

// if read full buffer return true else false
func ReadBuf(r i.Reader, dest []byte) bool {

	size := len(dest)
	proc := 0

	for proc < size {
		n, err := r.Read(dest[proc:])
		if err != nil {
			return false
		}
		if proc+n < size {
			proc += n
		} else {
			return true
		}
	}

	return true
}
