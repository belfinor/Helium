package ldb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-06-29

type Config struct {
	Path        string `json:"path"`
	Compression bool   `json:"compression"`
	FileSize    int    `json:"filesize"`
	ReadOnly    bool   `json:"readonly"`
}

func (c *Config) Clone() *Config {
	return &Config{
		Path:        c.Path,
		Compression: c.Compression,
		FileSize:    c.FileSize,
		ReadOnly:    c.ReadOnly,
	}
}
