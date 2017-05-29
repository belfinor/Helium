package fibonacci


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-05-29


type Fibonacci struct {
    current int64
    prev    int64
    number  int64
}


func New() *Fibonacci {
    return &Fibonacci { current: 1, prev: 0, number: 0 }
}


func (f *Fibonacci) GetCurrent() int64 {
    return f.current
}


func (f *Fibonacci) GetNumber() int64 {
    return f.number
}


func (f *Fibonacci) Next() int64 {
    f.current, f.prev = f.current + f.prev, f.current
    f.number++
    return f.current
}


func (f *Fibonacci) Reset() {
    f.current = 1
    f.number  = 0
    f.prev    = 0
}

