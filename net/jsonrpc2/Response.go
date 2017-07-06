package jsonrpc2


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06


import (
    "encoding/json"
)


type Response struct {
    Id      *json.RawMessage `json:"id"`
    JsonRPC string           `json:"jsonrpc"`
    Result  interface{}      `json:"result"`
    Error   interface{}      `json:"error"`
}

