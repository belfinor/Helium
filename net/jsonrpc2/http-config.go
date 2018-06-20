package jsonrpc2

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06

type HttpConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Url  string `json:"url"`
}
