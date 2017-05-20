package token


// @author  Mikhail Kirillov
// @email   mikkirillov@yandex.ru
// @version 1.000
// @date    2017-05-20

 

var mapper map[byte]byte = map[byte]byte{
    '\\' : '\\',
    '\'' : '\'',
    '"'  : '"',
    's'  : ' ',
    't'  : '\t',
    'r'  : '\r',
    'n'  : '\n',
    ' '  : ' ',
}


func isSpace( c byte ) bool {
    return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}


func Make( str string ) []string {    
    
    tokens := make( []string, 0, 10 )
    mode := 0
    bkmode := 0
    cur  := []byte{}

    for i := 0 ; i < len(str) ; i++ {
        
        c := str[i]

        switch mode {
        case 0:
            if isSpace(c) {
                continue
            }

            if c == '"' {
                mode = 1
                cur = []byte{}
                continue 
            }

            cur = []byte{ c }
            mode = 2

        case 1:

            if c == '"' {
                mode = 4
                tokens = append( tokens, string(cur) )
                continue
            }

            if c == '\\' {
                mode = 3
                bkmode = 1
                continue
            }

            cur = append( cur, c )

        case 2:
            
            if isSpace(c) {
                tokens = append( tokens, string(cur) )
                mode = 0
                continue
            }

            if c == '\\' {
                mode = 3
                bkmode = 2
                continue
            }

            cur = append( cur, c )

        case 3:
            
             nc, ok := mapper[c] 

             if !ok {
                 return nil
             }

             cur = append( cur, nc )
             mode = bkmode

        case 4:
             if isSpace(c) {
                 mode = 0
                 continue
             }

             return nil
        }
    }

    if mode == 1 || mode == 3 {
        return nil
    }

    if mode == 2 {
        tokens = append( tokens, string(cur) )
    }

    return tokens
}

