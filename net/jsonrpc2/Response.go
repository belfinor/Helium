package jsonrpc2


import (
    "encoding/json"
)


type Response struct {
    Id      *json.RawMessage `json:"id"`
    JsonRPC string           `json:"jsonrpc"`
    Result  interface{}      `json:"result"`
    Error   interface{}      `json:"error"`
}

