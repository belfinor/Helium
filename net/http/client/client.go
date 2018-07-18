package client

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-10-06

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
)

type Client struct {
	Timeout   time.Duration
	UserAgent string
	NoDecode  bool
}

func New() *Client {
	return &Client{}
}

type Response struct {
	StatusCode int
	Content    []byte
	Header     http.Header
}

func (c *Client) Request(method string, url string, headers map[string]string, content []byte) (*Response, error) {

	timeout := time.Duration(5 * time.Second)

	if c.Timeout > 0 {
		timeout = c.Timeout
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ua := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	tr.DisableKeepAlives = true

	var rd io.Reader

	if content != nil && len(content) > 0 && (method == "POST" || method == "PUT") {
		rd = bytes.NewReader(content)
	}

	req, err := http.NewRequest(method, url, rd)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := ua.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var text []byte

	if c.NoDecode {
		text, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		utf8, err1 := charset.NewReader(resp.Body, resp.Header.Get("Content-Type"))
		if err1 != nil {
			return nil, err1
		}

		text, err = ioutil.ReadAll(utf8)
		if err != nil {
			return nil, err
		}
	}

	r := &Response{
		StatusCode: resp.StatusCode,
		Content:    text,
		Header:     resp.Header,
	}

	return r, nil
}

func AgentString() string {
	list := []string{
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; WOW64; Trident/6.0)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.84 Safari/537.36",
		"Opera/9.80 (Windows NT 6.1; WOW64; U; ru) Presto/2.10.289 Version/12.00",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.109 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.95 YaBrowser/13.10.1500.9323 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/58.0.3029.81 Chrome/58.0.3029.81 Safari/537.36",
	}

	return list[int(time.Now().UnixNano()%int64(len(list)))]
}
