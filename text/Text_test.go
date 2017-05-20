package text


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-20



import (
    "testing"
)


func TestTextGetWords( t *testing.T ) {
    text := "hello World! Сегодня хороший день для 1-ого теста. Ёлка в лесу"

    res := GetWords(text)

    wait := []string{ "hello", "world", "сегодня", "хороший", "день", "для", "1", "ого", "теста", "елка", "в", "лесу" }

    if len(res) != len(wait) {
        t.Fatal( "bad result list size" )
    }

    for i, val := range res {
        if val != wait[i] {
            t.Fatal( "invalid item in slice" )
        }
    }
}


func TestTextMakePrefSuf( t *testing.T ) {
    wait := [][]string{
        []string{ "", "тест" },
        []string{ "т", "ест" },
        []string{ "те", "ст" },
        []string{ "тес", "т" },
        []string{ "тест", "" },
    }

    res := MakePrefSuf( "тест" )

    if len(res) != len(wait ) {
        t.Fatal( "Invalid result length" )
    }

    for i, rec := range res {
        if wait[i][0] != rec[0] || wait[i][1] != rec[1] {
            t.Fatal( "Invalid result" )
        }
    }
}

