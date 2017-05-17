package pack


// @author  Mikhail Kirillov
// @version 1.000
// @date    2017-05-17


import (
    "bytes"
    "encoding/binary"
)



func Int2Bytes( val int64 ) []byte {
    b := new(bytes.Buffer)
    binary.Write( b, binary.BigEndian, val )
    return b.Bytes()
}


func IntList2Bytes( vals []int64 ) []byte {
    if vals == nil {
        return make( []byte, 0 )
    }
    
    b := new(bytes.Buffer)
    
    for _, val := range vals {
        binary.Write( b, binary.BigEndian, val )
    }
    
    return b.Bytes()
}


func Bytes2Int( b []byte ) int64 {
    var val int64
    reader := bytes.NewReader(b)
    err := binary.Read( reader, binary.BigEndian, &val )
    if err != nil {
        val = 0
    }
    return val
}


func Bytes2IntList( b []byte ) []int64 {
    if b == nil {
        return make( []int64, 0 )
    }

    var val int64

    size   := int(len(b) / 8)
    res    := make( []int64, size )
    reader := bytes.NewReader(b)

    for i := 0 ; i < size ; i++ {
        binary.Read( reader, binary.BigEndian, &val )
        res[i] = val
    }

    return res
}


