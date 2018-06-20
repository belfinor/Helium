package jsonrpc2

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06

import (
	"encoding/json"
)

type Request struct {
	Method  string           `json:"method"`
	JsonRPC string           `json:"jsonrpc"`
	Params  *json.RawMessage `json:"params"`
	Id      *json.RawMessage `json:"id"`
}
