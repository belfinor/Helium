package errors


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-06-01


import (
    "github.com/belfinor/Helium/log"
)


type Error struct {
    Code string
    Text string
}


var codes map[string][]string = map[string][]string{
    "E0001": []string{ "Внутренняя ошибка", "Внутренняя ошибка" },
}


func New( code string ) Error {
    text, has := codes[code]

    if has {
        if len(text) > 1 {
            for _, msg := range text[1:] {
                log.Error( msg )
            }
        }
        return Error{ Code: code, Text: text[0] }
    }

    return Error{ Code: "E0001", Text: "Внутренняя ошибка" }
}


func (e Error) Error() string {
    return e.Code + ": " + e.Text
}


func SetCodes( tab map[string][]string ) {
    codes = tab
}

