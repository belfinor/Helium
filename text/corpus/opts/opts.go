package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-12-04

import (
	"strings"
)

const OPT_EN int64 = 0x0000000000000001 // английский язык
const OPT_RU int64 = 0x0000000000000002 // русский язык

const OPT_MR int64 = 0x0000000000000004 // мужской род
const OPT_GR int64 = 0x0000000000000008 // женский род
const OPT_SR int64 = 0x0000000000000010 // средний род
const OPT_ML int64 = 0x0000000000000020 // множественное число

const OPT_IP int64 = 0x0000000000000040 // именительный падеж
const OPT_RP int64 = 0x0000000000000080 // родительный падеж
const OPT_DP int64 = 0x0000000000000100 // дательные падеж
const OPT_VP int64 = 0x0000000000000200 // винительный падеж
const OPT_TP int64 = 0x0000000000000400 // творительный падеж
const OPT_PP int64 = 0x0000000000000800 // предложный падеж

const OPT_NOUN int64 = 0x0000000000001000     // сществительное
const OPT_ADJ int64 = 0x0000000000002000      // прилагательное
const OPT_VERB int64 = 0x0000000000004000     // глагол
const OPT_ADV int64 = 0x0000000000008000      // наречие
const OPT_UNION int64 = 0x0000000000010000    // союз
const OPT_PRETEXT int64 = 0x0000000000020000  // предлог
const OPT_PRONOUN int64 = 0x0000000000040000  // местоимение
const OPT_ARTICLE int64 = 0x0000000000080000  // артикль
const OPT_PARTICLE int64 = 0x0000000000100000 // частица
const OPT_ADVPART int64 = 0x0000000000200000  // деепричастие
const OPT_INTER int64 = 0x0000000000400000    // междометие
const OPT_PADJ int64 = 0x0000000000800000     // причастие
const OPT_NUMERAL int64 = 0x0000000001000000  // числительное

const OPT_FACE1 int64 = 0x0000000002000000 // 1-ое лицо
const OPT_FACE2 int64 = 0x0000000004000000 // 2-ое лицо
const OPT_FACE3 int64 = 0x0000000008000000 // 3-ье лицо

const OPT_UNDEF int64 = 0x0000000010000000  // неопределенная форма
const OPT_PAST int64 = 0x0000000020000000   // прошедшее время
const OPT_PRE int64 = 0x0000000040000000    // настоящее время
const OPT_FUTURE int64 = 0x0000000080000000 // будущее время
const OPT_IMP int64 = 0x0000000100000000    // повелительное наклонение

const OPT_SHORT int64 = 0x0000000200000000 // краткая форма для прилагателных

var optList []string = []string{"en", "ru",
	"mr", "gr", "sr", "ml",
	"ip", "rp", "dp", "vp", "tp", "pp",
	"noun", "adj", "verb", "adv", "union", "pretext", "pronoun", "article", "particle", "advpart", "inter", "padj", "num",
	"f1", "f2", "f3",
	"undef", "past", "pre", "future", "imp",
	"short",
}

var nameToCode map[string]int64
var codeToName map[int64]string

func init() {

	code := int64(0x0000000000000001)

	nameToCode = make(map[string]int64)
	codeToName = make(map[int64]string)

	for _, opt := range optList {
		nameToCode[opt] = code
		codeToName[code] = opt

		code = code << 1
	}
}

type Opt int64

func (o Opt) String() string {
	builder := strings.Builder{}

	val := int64(o)
	cur := int64(0x0000000000000001)

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

func (o Opt) Include(other Opt) bool {
	return (o & other) == other
}

func Parse(src string) Opt {
	res := int64(0)

	for _, v := range strings.Split(src, ".") {
		if v, has := nameToCode[v]; has {
			res = res | v
		}
	}

	return Opt(res)
}
