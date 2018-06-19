package boltdb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-22

type Config struct {
	Database string `json:"database"`
	Mask     int    `json:"mask"`
	ReadOnly bool   `json:"read_only"`
	Timeout  int    `json:"timeout"`
}
