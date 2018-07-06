package sql

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-07-06

type Config struct {
	Driver  string `json:"driver"`
	Connect string `json:"connect"`
}
