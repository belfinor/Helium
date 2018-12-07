package opts

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.003
// @date    2018-12-07

import (
	"strings"
)

const OPT_EN int32 = 0x000001      // английский язык
const OPT_RU int32 = 0x000002      // русский язык
const OPT_NOUN int32 = 0x000004    // сществительное
const OPT_ADJ int32 = 0x000008     // прилагательное
const OPT_VERB int32 = 0x000010    // глагол
const OPT_ADV int32 = 0x000020     // наречие
const OPT_UNION int32 = 0x000040   // союз
const OPT_PRETEXT int32 = 0x000080 // предлог
const OPT_PRONOUN int32 = 0x000100 // местоимение
const OPT_ART int32 = 0x000200     // артикль
const OPT_PART int32 = 0x000400    // частица
const OPT_ADVPART int32 = 0x000800 // деепричастие
const OPT_INTER int32 = 0x001000   // междометие
const OPT_PADJ int32 = 0x002000    // причастие
const OPT_NUMERAL int32 = 0x004000 // числительное
const OPT_MR int32 = 0x008000      // мужской род
const OPT_GR int32 = 0x010000      // женский род
const OPT_SR int32 = 0x020000      // средний род
const OPT_ML int32 = 0x040000      // множественное число
const OPT_NUM int32 = 0x080000     // число
const OPT_EOS int32 = 0x100000     // конец предложения
const OPT_ALIVE int32 = 0x200000   // объект является живым

var optList []string = []string{
	"en", "ru",
	"noun", "adj", "verb", "adv", "union", "pretext", "pronoun", "art", "part", "advpart", "inter", "padj", "numeral",
	"mr", "gr", "sr", "ml", "num",
	"eos",
	"alive",
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
