package router

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-10-02

import (
	"strconv"
)

type Params map[string]string

func (p Params) GetString(name string) string {
	if v, h := p[name]; h {
		return v
	}

	return ""
}

func (p Params) GetInt(name string) int64 {
	if v, h := p[name]; h {
		if val, err := strconv.ParseInt(v, 10, 64); err == nil {
			return val
		}

		return 0
	}

	return 0
}

func (p Params) GetFloat(name string) float64 {
	if v, h := p[name]; h {
		if val, err := strconv.ParseFloat(v, 64); err == nil {
			return val
		}

		return 0
	}

	return 0
}
