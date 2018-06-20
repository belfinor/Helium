package jsonrpc2

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
