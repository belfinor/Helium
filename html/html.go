package html


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-22


import (
    "bytes"
    ht "golang.org/x/net/html"
    "io"
)


type HTML struct {
    Links      map[string]bool
    TagErase   map[string]bool
    TagStack   []string
    EraseStack []string
    Document   string
    Text       string
}


func NewHtmlParser() *HTML {
    res := &HTML{}

    res.Links = make( map[string]bool )
    
    res.TagErase = map[string]bool{
        "form":    true,
        "iframe":  true,
        "link":    true,
        "meta":    true,
        "noscript":true,
        "option":  true,
        "script":  true,
        "select":  true,
        "style":   true,
    }

    res.TagStack   = make( []string, 0 )
    res.EraseStack = make( []string, 0 )

    res.Document = ""
    res.Text = ""

    return res
}


func (h *HTML) ProcessString(str string) string {
    r := bytes.NewReader( []byte(str) )
    return h.ProcessReader(r)
}

func (h *HTML) ProcessReader( r io.Reader ) string {

    parser := ht.NewTokenizer(r)

    
    for {
        tt := parser.Next()

        switch {
        case tt == ht.ErrorToken:
            return h.Text

        case tt == ht.StartTagToken:
            t := parser.Token()
            h.onStartTag(&t, string(parser.Raw()))

        case tt == ht.EndTagToken:
            t := parser.Token()
            h.onCloseTag(&t)

        case tt == ht.SelfClosingTagToken:
            t := parser.Token()
            h.onStartTag(&t, string(parser.Raw()))
            h.onCloseTag(&t)
        
        case tt == ht.TextToken:
            h.onText( string(parser.Text()), string(parser.Raw()) )

        }
    }

    return h.Text
}

func (h *HTML) onStartTag( t *ht.Token, raw string ) {

    if len(h.EraseStack) > 0 {
        h.EraseStack = append( h.EraseStack, t.Data )
        return
    }

    _, found := h.TagErase[ t.Data ]

    if found {
        if t.Data != "meta" && t.Data != "link" {
            h.EraseStack = append( h.EraseStack, t.Data )
        }
        return
    }

    h.TagStack = append( h.TagStack, t.Data )

    if t.Data == "a" {
        for _, a := range t.Attr {
            if a.Key == "href" {
                h.Links[ a.Val ] = true
            }
        }
    }

    h.Document += raw
}


func (h *HTML) onCloseTag( t *ht.Token ) {

    elen := len(h.EraseStack)

    if elen > 0 {

        if h.EraseStack[elen-1] == t.Data {
            h.EraseStack = h.EraseStack[0:elen-1]
            return
        }

        pos := -1

        for i, cur := range h.EraseStack {
            if cur == t.Data {
                pos = i
            }
        }

        if pos > -1 {
            h.EraseStack = h.EraseStack[0:pos]
        }

        return
    }
    
    elen = len(h.TagStack)

    if len(h.TagStack) > 0 {
        
        if( h.TagStack[elen-1] == t.Data ) {
            h.Document += "</" + t.Data + ">"
            h.TagStack = h.TagStack[0:elen-1]
            return
        }

        _, found := h.TagErase[t.Data]
        if found {
            return
        }

        pos := -1

        for i, cur := range h.TagStack {
            if cur == t.Data {
                pos = i
            }
        }

        if pos > -1 {
            
            for i := elen - 1 ; i >= pos ; i-- {
                h.Document += "</" + h.TagStack[i] + ">"
            }

            h.TagStack = h.TagStack[0:pos]
        }
    }
}


func (h *HTML) onText( str string, raw string ) {
    if len(h.EraseStack) == 0 {
        h.Document += raw
        h.Text += " " + str
    }
}

