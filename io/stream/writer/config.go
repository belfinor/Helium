package stream


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-24


type WriterConfig struct {
  Path   string `json:"path"`
  Index  string `json:"index"`
  Buffer int    `json:"buffer"`
  Save   int64  `json:"save"`
  Period int64  `json:"period"`
  LogId  int64  `json:"id"`
}

