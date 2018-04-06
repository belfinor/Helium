package sociation


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-10-07


import (
  "encoding/json"
  "net/url"
  "github.com/belfinor/Helium/log"
  "github.com/belfinor/Helium/net/http/client"
)


type SociumRec struct {
  Name string `json:"name"`
  Direct int64 `json:"popularity_direct"`
  Inverse int64 `json:"popularity_inverse"`
}


type SociumResp struct {
  Associations []SociumRec `json:"associations"`
  Word         string      `json:"word"`
}


func Get( phrase string ) []string {

  form := url.Values{}

  form.Add( "max_count", "0" )
  form.Add( "back", "false" )
  form.Add( "word", phrase )

  content := form.Encode()

  headers := map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
    "Referer": "http://sociation.org/graph/",
    "User-Agent": "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
  }


  res, err := client.New().Request( "POST", "http://sociation.org/ajax/word_associations/", headers , []byte(content) )
  if err != nil {
    log.Error( err.Error() )
    return []string{}
  }

  var resp SociumResp
  if err = json.Unmarshal( res.Content, &resp ) ; err != nil {
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


func GetFull( phrase string ) []SociumRec {

  form := url.Values{}

  form.Add( "max_count", "0" )
  form.Add( "back", "false" )
  form.Add( "word", phrase )

  content := form.Encode()

  headers := map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
    "Referer": "http://sociation.org/graph/",
    "User-Agent": "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36",
  }


  res, err := client.New().Request( "POST", "http://sociation.org/ajax/word_associations/", headers , []byte(content) )
  if err != nil {
    log.Error( err.Error() )
    return nil
  }

  var resp SociumResp
  if err = json.Unmarshal( res.Content, &resp ) ; err != nil {
    log.Error( err.Error() )
    return nil
  }

  return resp.Associations

}

