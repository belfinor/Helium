package jsonrpc2


type Error struct {
  Code int64 `json:"code"`
  Message string `json:"message"`
}

