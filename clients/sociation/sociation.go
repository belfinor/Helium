package sociation


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-09-20


import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "golang.org/x/net/html/charset"
  "strings"
  "github.com/belfinor/Helium/log"
)


type SociumRec struct {
  Name string `json:"name"`
}


type SociumResp struct {
  Associations []SociumRec `json:"associations"`
  Word         string      `json:"word"`
}


func Get( phrase string ) []string {

  client := http.Client{}

  form := url.Values{}

  form.Add( "max_count", "0" )
  form.Add( "back", "false" )
  form.Add( "word", phrase )

  req, err := http.NewRequest("POST", "http://sociation.org/ajax/word_associations/", strings.NewReader(form.Encode()))
  if err != nil {
    log.Error( "sociation: " + err.Error() )
    return []string{}
  }

  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Referer", "http://sociation.org/graph/")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")

  res, err1 := client.Do( req )

  if err1 != nil {
    log.Error( "sociation: " + err1.Error() )
    return []string{}
  }

  defer res.Body.Close()

  utf8, err := charset.NewReader(res.Body, res.Header.Get("Content-Type"))
  if err != nil {
    log.Error( "sociation: " + err.Error() )
    return []string{}
  }

  var text []byte

  text, err = ioutil.ReadAll(utf8)
  if err != nil {
    log.Debug( "sociation: " + err.Error() )
    return []string{}
  }

  var resp SociumResp
  if err = json.Unmarshal( text, &resp ) ; err != nil {
    log.Error( err.Error() )
    return []string{}
  }

  if resp.Associations == nil {
    log.Debug( "sociation: no associatios" )
    return []string{}
  }

  result := make( []string, len(resp.Associations) )
  for i, v := range resp.Associations {
    result[i] = v.Name
  }

  if len(result) == 0 {
    log.Debug( "sociation: empty response" )
  }

  return result
}

