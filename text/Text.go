package text


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-20



import (
    "regexp"
    "strings"
)


func GetWords( text string ) []string {
    src   := strings.ToLower(text) 
    src    =  strings.Replace( src, "ё", "е", -1 )
    
    re,_  := regexp.Compile( "([a-z]+|\\d+|[а-я]+)" )
    words := re.FindAllString(src,-1)

    if words == nil {
        words = make( []string, 0 )
    }

    return words
}


func MakePrefSuf( str string ) [][]string {
    list := [][]string{}
    prefix := ""

    for _, c := range str {
        rune := string(c)
        for _, val := range list {
            val[1] = val[1] + rune
        }
        list = append( list, []string{ prefix, rune } )
        prefix = prefix + rune
    }

    list = append( list, []string{ prefix, "" } )

    return list
}

