package writer


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-24


import(
  "io/ioutil"
  "strconv"
  "strings"
)


type Config struct {
  Path   string `json:"path"`
  Index  string `json:"index"`
  Buffer int    `json:"buffer"`
  Save   int64  `json:"save"`
  Period int64  `json:"period"`
  LogId  int64  `json:"id"`
}


func (c *Config ) LoadLogId() {
  if c.Buffer < 10 {
    c.Buffer = 10
  }

  // load index file
  if data, err := ioutil.ReadFile( c.Index ) ; err != nil {
    c.LogId = 1
    ioutil.WriteFile( c.Index, []byte("1"), 0664 )
  } else {
    str := strings.TrimSpace( string(data) )
    if c.LogId, err = strconv.ParseInt( str, 10, 64 ) ; err != nil {
      return
    }
    c.LogId++
  }
}

