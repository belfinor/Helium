package io

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-04

import (
	"strings"
	"testing"
)

func TestReadBuf(t *testing.T) {
	buf := make([]byte, 4)
	src := "tes1tes2tes3te"
	r := strings.NewReader(src)

	for i := 0; i < 3; i++ {
		if !ReadBuf(r, buf) || src[i*4:i*4+4] != string(buf) {
			t.Fatal("Normal read error")
		}
	}

	if ReadBuf(r, buf) {
		t.Fatal("Not full read success")
	}
}
