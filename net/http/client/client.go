package client


import (
  "bytes"
  "crypto/tls"
  "io"
  "io/ioutil"
  "golang.org/x/net/html/charset"
  "net/http"
  "time"
)


type Client struct {
}


func New() *Client {
  return &Client{}
}


type Response struct {
  StatusCode int
  Content   []byte
  Header    http.Header
}


func Request( method string, url string, headers map[string]string, content []byte ) ( *Response, error ) {

  timeout := time.Duration(5 * time.Second)

  tr := &http.Transport{
     TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  ua := &http.Client{
    Timeout:   timeout, 
    Transport: tr,
  }

  var rd io.Reader

  if content != nil && len(content) > 0 && ( method == "POST" || method == "PUT" ) {
    rd = bytes.NewReader(content)
  }

  req, err := http.NewRequest( method, url, rd )
  
  for k, v := range headers {
    req.Header.Set( k, v )
  }

  resp, err := ua.Do( req )

  defer resp.Body.Close()

  utf8, err1 := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
  if err1 != nil {
    return nil, err1
  }

  var text []byte

  text, err = ioutil.ReadAll(utf8)
  if err != nil {
    return nil, err
  }

  r := &Response {
    StatusCode: resp.StatusCode,
    Content:    text,
    Header:     resp.Header,
  }

  return r, nil
}

