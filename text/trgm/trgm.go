package trgm

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-15

import (
	"hash/crc64"
	"strings"

	"github.com/belfinor/Helium/text"
	"github.com/belfinor/Helium/text/buffer"
	"github.com/belfinor/Helium/text/stemmer"
)

const (
	BUFFER_SIZE int = 3
)

var crcTab *crc64.Table = crc64.MakeTable(crc64.ECMA)

func trgmStream(in <-chan string) <-chan uint64 {

	out := make(chan uint64, 10000)

	skip := getSkipData()

	bufIn := buffer.New(BUFFER_SIZE)
	bufOut := buffer.New(BUFFER_SIZE)

	procIn := func(skipFullCheck bool) {

		if !skipFullCheck && !bufIn.Full() {
			return
		}

		if bufIn.Empty() {
			return
		}

		skipCnt := 0
		cur := ""

		for i := 0; i < bufIn.Size(); i++ {
			if i > 0 {
				cur += " "
			}
			cur += bufIn.Get(i)

			if _, h := skip[cur]; h {
				skipCnt = i + 1
			}
		}

		if skipCnt > 0 {
			bufIn.Shift(skipCnt)
			return
		}

		bufOut.Add(bufIn.Get(0))
		bufIn.Shift(1)
	}

	procOut := func() {
		if !bufOut.Full() {
			return
		}

		str := bufOut.Join(" ")
		bufOut.Shift(1)

		out <- crc64.Checksum([]byte(str), crcTab)

	}

	go func() {

		for w := range in {

			procIn(false)
			procOut()

			bufIn.Add(w)
		}

		for !bufIn.Empty() {
			procIn(true)
			procOut()
		}

		procOut()

		close(out)
	}()

	return out
}

func MakeTrgms(txt string) map[uint64]bool {

	exists := make(map[uint64]bool, 128)

	wordStream := text.WordStream(strings.NewReader(txt), text.WSO_NO_URLS, text.WSO_HASHTAG)
	tokStream := stemmer.Stream(wordStream)

	for val := range trgmStream(tokStream) {
		exists[val] = true
	}

	return exists
}

func UsedPart(txt string, source string) float64 {

	first := MakeTrgms(txt)
	second := MakeTrgms(source)

	return HashPart(first, second)
}

func HashPart(hash map[uint64]bool, source map[uint64]bool) float64 {
	has := 0
	size := len(hash)

	if size == 0 {
		return 0
	}

	for k := range hash {
		if _, h := source[k]; h {
			has++
		}
	}

	return float64(has) / float64(size)
}

func Sim(txt1 string, txt2 string) float64 {
	hash1 := MakeTrgms(txt1)
	hash2 := MakeTrgms(txt2)

	return HashSim(hash1, hash2)
}

func HashSim(hash1 map[uint64]bool, hash2 map[uint64]bool) float64 {
	has := 0

	size1 := len(hash1)
	if size1 == 0 {
		return 0
	}

	size2 := len(hash2)
	if size2 == 0 {
		return 0
	}

	for k := range hash1 {
		if _, h := hash2[k]; h {
			has++
		}
	}

	return float64(has) / float64(size1+size2-has)
}
