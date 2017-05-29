package uniq


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-29


import (
    "fmt"
    "github.com/belfinor/Helium/math/fibonacci"
    "hash/crc32"
    "math/rand"
    "sync"
    "time"
)


type Uniq struct {
    random *rand.Rand
    fibo   *fibonacci.Fibonacci
    tact   int64
    sync.Mutex
}


func New() *Uniq {
    return &Uniq{
        random: rand.New( rand.NewSource( time.Now().Unix() ) ),
        fibo:   fibonacci.New(),
        tact:   0,
    }
}


func (u *Uniq) Next() string {
    u.Lock()
    defer u.Unlock()

    u.tact++

    str := fmt.Sprintf( "%x-%x-%x-%x", time.Now().UnixNano(), u.tact, u.random.Intn(0x10000), u.fibo.Next() )

    tab := crc32.MakeTable(crc32.IEEE)
    str = fmt.Sprintf( "%s-%x", str, crc32.Checksum( []byte(str), tab ) )

    return str
}

