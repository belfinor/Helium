package log

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-06-21

type Config struct {
	Template string `json:"template"`
	Period   int    `json:"period"`
	Save     int    `json:"save"`
	Level    string `json:"level"`
	StdOut   bool   `json:"stdout"`
	StdErr   bool   `json:"stderr"`
}
