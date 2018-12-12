package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.006
// @date    2018-12-12

import (
	"strings"
)

const OPT_EN int64 = 0x00000001      // английский язык
const OPT_RU int64 = 0x00000002      // русский язык
const OPT_NOUN int64 = 0x00000004    // сществительное
const OPT_ADJ int64 = 0x00000008     // прилагательное
const OPT_VERB int64 = 0x00000010    // глагол
const OPT_ADV int64 = 0x00000020     // наречие
const OPT_UNION int64 = 0x00000040   // союз
const OPT_PRETEXT int64 = 0x00000080 // предлог
const OPT_PRONOUN int64 = 0x00000100 // местоимение
const OPT_ART int64 = 0x00000200     // артикль
const OPT_PART int64 = 0x00000400    // частица
const OPT_ADVPART int64 = 0x00000800 // деепричастие
const OPT_INTER int64 = 0x00001000   // междометие
const OPT_PADJ int64 = 0x00002000    // причастие
const OPT_NUMERAL int64 = 0x00004000 // числительное
const OPT_MR int64 = 0x00008000      // мужской род
const OPT_GR int64 = 0x00010000      // женский род
const OPT_SR int64 = 0x00020000      // средний род
const OPT_ML int64 = 0x00040000      // множественное число
const OPT_NUM int64 = 0x00080000     // число
const OPT_EOS int64 = 0x00100000     // конец предложения
const OPT_ALIVE int64 = 0x00200000   // объект является живым
const OPT_IP int64 = 0x00400000      // ИП
const OPT_RP int64 = 0x00800000      // РП
const OPT_DP int64 = 0x01000000      // ДП
const OPT_VP int64 = 0x02000000      // ВП
const OPT_TP int64 = 0x04000000      // ТП
const OPT_PP int64 = 0x08000000      // ПП
const OPT_ROMAN int64 = 0x10000000   // ПП
const OPT_SEP int64 = 0x20000000     // SEP (NOT STETEMENT END!)

var optList []string = []string{
	"en", "ru",
	"noun", "adj", "verb", "adv", "union", "pretext", "pronoun", "art", "part", "advpart", "inter", "padj", "numeral",
	"mr", "gr", "sr", "ml", "num",
	"eos",
	"alive",
	"ip", "rp", "dp", "vp", "tp", "pp",
	"roman",
	"sep",
}

var nameToCode map[string]int64
var codeToName map[int64]string

func init() {

	code := int64(0x00000001)

	nameToCode = make(map[string]int64, 32)
	codeToName = make(map[int64]string, 32)

	for _, opt := range optList {
		nameToCode[opt] = code
		codeToName[code] = opt

		code = code << 1
	}
}

func ToString(o int64) string {
	builder := strings.Builder{}

	val := o
	cur := int64(0x00000001)

	for _, v := range optList {
		if val&cur != 0 {
			if builder.Len() != 0 {
				builder.WriteRune('.')
			}
			builder.WriteString(v)
		}
		cur = cur << 1
	}

	return builder.String()
}

// opt include all opts from other
func Include(o1, o2 int64) bool {
	return (o1 & o2) == o2
}

func Parse(src string) int64 {
	res := int64(0)

	for _, v := range strings.Split(src, ".") {
		if v, has := nameToCode[v]; has {
			res = res | v
		}
	}

	return res
}
