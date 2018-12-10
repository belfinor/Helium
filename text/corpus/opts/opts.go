package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.004
// @date    2018-12-10

import (
	"strings"
)

const OPT_EN int32 = 0x00000001      // английский язык
const OPT_RU int32 = 0x00000002      // русский язык
const OPT_NOUN int32 = 0x00000004    // сществительное
const OPT_ADJ int32 = 0x00000008     // прилагательное
const OPT_VERB int32 = 0x00000010    // глагол
const OPT_ADV int32 = 0x00000020     // наречие
const OPT_UNION int32 = 0x00000040   // союз
const OPT_PRETEXT int32 = 0x00000080 // предлог
const OPT_PRONOUN int32 = 0x00000100 // местоимение
const OPT_ART int32 = 0x00000200     // артикль
const OPT_PART int32 = 0x00000400    // частица
const OPT_ADVPART int32 = 0x00000800 // деепричастие
const OPT_INTER int32 = 0x00001000   // междометие
const OPT_PADJ int32 = 0x00002000    // причастие
const OPT_NUMERAL int32 = 0x00004000 // числительное
const OPT_MR int32 = 0x00008000      // мужской род
const OPT_GR int32 = 0x00010000      // женский род
const OPT_SR int32 = 0x00020000      // средний род
const OPT_ML int32 = 0x00040000      // множественное число
const OPT_NUM int32 = 0x00080000     // число
const OPT_EOS int32 = 0x00100000     // конец предложения
const OPT_ALIVE int32 = 0x00200000   // объект является живым
const OPT_IP int32 = 0x00400000      // ИП
const OPT_RP int32 = 0x00800000      // РП
const OPT_DP int32 = 0x01000000      // ДП
const OPT_VP int32 = 0x02000000      // ВП
const OPT_TP int32 = 0x04000000      // ТП
const OPT_PP int32 = 0x08000000      // ПП

var optList []string = []string{
	"en", "ru",
	"noun", "adj", "verb", "adv", "union", "pretext", "pronoun", "art", "part", "advpart", "inter", "padj", "numeral",
	"mr", "gr", "sr", "ml", "num",
	"eos",
	"alive",
	"ip", "rp", "dp", "vp", "tp", "pp",
}

var nameToCode map[string]int32
var codeToName map[int32]string

func init() {

	code := int32(0x00000001)

	nameToCode = make(map[string]int32, 32)
	codeToName = make(map[int32]string, 32)

	for _, opt := range optList {
		nameToCode[opt] = code
		codeToName[code] = opt

		code = code << 1
	}
}

type Opt int32

func (o Opt) String() string {
	builder := strings.Builder{}

	val := int32(o)
	cur := int32(0x00000001)

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
func (o Opt) Include(other Opt) bool {
	return (o & other) == other
}

func Parse(src string) Opt {
	res := int32(0)

	for _, v := range strings.Split(src, ".") {
		if v, has := nameToCode[v]; has {
			res = res | v
		}
	}

	return Opt(res)
}
