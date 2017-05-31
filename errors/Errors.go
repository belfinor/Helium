package errors


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-31


type Error struct {
    Code string
    Text string
}


var codes map[string]string = map[string]string{
    "E0001": "Внутренняя ошибка",
}


func New( code string ) Error {
    text, has := codes[code]

    if has {
        return Error{ Code: code, Text: text }
    }

    return Error{ Code: "E0001", Text: "Внутренняя ошибка" }
}


func (e *Error) Error() string {
    return e.Code + ": " + e.Text
}


func SetCodes( tab map[string]string ) {
    codes = tab
}

