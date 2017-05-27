package test


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-27


import (
    "github.com/belfinor/Helium/db/ldb"
)


func Init() {
    ldb.Init( "/tmp/ldb.test" )

    for {
        list := ldb.List( []byte{}, 1000, 0, false )
        
        for _, key := range list {
            ldb.Del( key )
        }

        if len(list) < 1000 {
            break
        }
    }   
}

