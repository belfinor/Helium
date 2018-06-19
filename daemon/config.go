package daemon

// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-24

type Config struct {
	PidFile string `json:"pid"`
	LogFile string `json:"log"`
	WordDir string `json:"dir"`
}
