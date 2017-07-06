package jsonrpc2


import (
    "encoding/json"
)


type Request struct {
    Method  string           `json:"method"`
    JsonRPC string           `json:"jsonrpc"`
    Params  *json.RawMessage `json:"params"`
    Id      *json.RawMessage `json:"id"`
}

