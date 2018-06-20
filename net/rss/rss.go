package rss

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-18

import (
	"bytes"
	"encoding/xml"
	"github.com/belfinor/Helium/net/http/client"
	"golang.org/x/net/html/charset"
	"time"
)

type Rss struct {
	Channel RssChannel `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type RssChannel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RssItem `xml:"item"`
}

func Get(url string) []RssItem {
	ua := client.New()

	ua.UserAgent = "Mozilla/5.0 (Windows NT 6.1) Gecko/20100101 Thunderbird/52.7.0 Lightning/5.4"
	ua.Timeout = time.Second * 15

	resp, err := ua.Request("GET", url, nil, nil)

	if err != nil || resp == nil || resp.Content == nil {
		return nil
	}

	buffer := bytes.NewBuffer([]byte(resp.Content))
	xml := xml.NewDecoder(buffer)

	xml.CharsetReader = charset.NewReaderLabel

	rss := new(Rss)

	if err = xml.Decode(rss); err != nil {
		return nil
	}

	return rss.Channel.Items
}
