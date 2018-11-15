package trgm

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-15

import (
	"github.com/belfinor/Helium/text/stemmer"
)

type SkipData map[string]bool

func (s SkipData) Has(val string) bool {
	_, h := s[val]
	return h
}

var skipData []string = []string{
	"а",
	"бы",
	"возможно",
	"же",
	"и",
	"или",
	"итак",
	"ли",
	"либо",
	"может быть",
	"но",
	"однако",
	"пусть",
	"так",
	"также",
	"такой",
	"то",
	"тоже",
	"что",
	"чтоб",
	"чтобы",
	"я",
}

func getSkipData() SkipData {

	res := make(map[string]bool, len(skipData))

	for _, v := range skipData {
		res[stemmer.TextToCode(v)] = true
	}

	return res
}
