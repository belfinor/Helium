package types

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-12

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/belfinor/Helium/log"
	"github.com/belfinor/Helium/time/timer"
)

type FOREACH_FUNC func(t uint16)

var fromCode map[uint16]string
var toCode map[string]uint16

var TP_HDATE uint16
var TP_HPERSON uint16
var TP_LASTNAME uint16
var TP_MAN uint16
var TP_NAME uint16
var TP_NUMBER uint16
var TP_PATRONYMIC uint16
var TP_ROMAN uint16
var TP_SLANG uint16

func init() {

	fromCode = make(map[uint16]string, 128)
	toCode = make(map[string]uint16, 128)

}

// reload types from file
func Load(filename string) {

	fh, err := os.Open(filename)
	if err != nil {
		log.Error("error load corpus types from " + filename)
		return
	}
	defer fh.Close()

	log.Info("reload corpus types from " + filename)

	load(fh)

}

// reload types from string
func LoadFromString(txt string) {

	log.Info("reload corpus types from text")

	load(strings.NewReader(txt))
}

func load(rh io.Reader) {

	tm := timer.New()

	rb := bufio.NewReader(rh)

	i := uint16(1)

	rCode := make(map[uint16]string, 128)
	rStr := make(map[string]uint16, 128)

	appender := func(t string) {
		if rStr[t] != 0 {
			return
		}

		rCode[i] = t
		rStr[t] = i

		i++
	}

	for {
		str, err := rb.ReadString('\n')
		if err != nil && str == "" {
			break
		}

		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}

		appender(str)
	}

	for _, t := range []string{"имя", "истдата", "истлицо", "мат", "отчество", "римскцифра", "фамилия", "человек", "число", ".", ","} {
		appender(t)
	}

	fromCode = rCode
	toCode = rStr

	TP_HDATE = toCode["истдата"]
	TP_HPERSON = toCode["истлицо"]
	TP_MAN = toCode["человек"]
	TP_NAME = toCode["имя"]
	TP_NUMBER = toCode["число"]
	TP_LASTNAME = toCode["фамилия"]
	TP_PATRONYMIC = toCode["отчество"]
	TP_ROMAN = toCode["римскцифра"]
	TP_SLANG = toCode["мат"]
	TP_HDATE = toCode["истдата"]

	log.Info(fmt.Sprintf("corpus types reloaded %.4fs", tm.DeltaFloat()))
	log.Info(fmt.Sprintf("corpus types size = %d", Total()))
}

func ToCode(str string) uint16 {
	if v, h := toCode[str]; h {
		return v
	}

	return 0
}

func FromCode(code uint16) string {
	if v, h := fromCode[code]; h {
		return v
	}

	return ""
}

func Join(lst ...uint16) int64 {
	val := int64(0)

	for _, v := range lst {
		val = Append(val, v)
	}

	return val
}

func Append(val int64, code uint16) int64 {

	if code == 0 {
		return val
	}

	nv := int64(code)

	for i := 0; i < 4; i++ {
		if (val>>uint(16*i))&0xffff == nv {
			return val
		}
	}

	return (val << 16) | nv
}

func AppendCircle(val int64, code uint16) int64 {

	if code == 0 {
		return val
	}

	nv := int64(code)

	return (val << 16) | nv
}

func Total() int {
	return len(fromCode)
}

func ForEach(v int64, fn FOREACH_FUNC) {
	for i := uint(0); i < 4; i++ {
		code := uint16((v >> (i * 16)) & 0xffff)
		if code != 0 {
			fn(code)
		}
	}
}

func Has(val int64, code uint16) bool {
	nv := int64(code)

	for i := 0; i < 4; i++ {
		if (val>>uint(16*i))&0xffff == nv {
			return true
		}
	}

	return false
}
