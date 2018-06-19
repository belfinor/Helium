package ldb

type Config struct {
	Path        string `json:"path"`
	Compression bool   `json:"compression"`
	FileSize    int    `json:"filesize"`
	ReadOnly    bool   `json:"readonly"`
}
