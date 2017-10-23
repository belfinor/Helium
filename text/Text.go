package text


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2017-10-06



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


func Truncate( text string, limit int ) string {
  result := ""
  for i, rune := range text {
    if i >= limit {
      result += "..."
      break
    }
    result += string(rune)
  }
  return result
}


func GetPrefixes( str string ) []string {
  list := make( []string, 0, 16 )
  prefix := ""

  for _, c := range str {
    prefix = prefix + string(c)
    list = append( list, prefix )
  }

  return list
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


func SubStrDistVal( str string, minlen int ) map[string]int {

  runes := make( []string, 0, 20 )

  for _, r := range str {
    runes = append( runes, string(r) )
  }

  if minlen < 3 {
    minlen = 3
  }

  size := len(runes)
  res := make( map[string]int )

  if size < minlen {
    return res
  }

  for {
    clen := len(runes)
    data := runes

    if clen < minlen {
      break
    }

    for clen >= minlen {
      cstr := strings.Join( data, "" )
      res[cstr] = size - clen
      data = data[:clen-1]
      clen--
    }

    runes = runes[1:]
  }

  return res
}

