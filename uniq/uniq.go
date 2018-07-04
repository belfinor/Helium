package uniq

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.002
// @date    2018-07-04

import (
	"context"
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"

	"github.com/belfinor/Helium/math/num/fibo"
)

type Uniq struct {
	next   chan string
	cancel context.CancelFunc
}

func New() *Uniq {
	obj := &Uniq{
		next: make(chan string, 10),
	}

	ctx, cancel := context.WithCancel(context.Background())

	obj.cancel = cancel

	go maker(ctx, obj.next)

	return obj
}

func maker(ctx context.Context, stream chan string) {
	fb := fibo.New()
	defer fb.Close()

	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	tact := time.Now().Unix() & 0xffff
	crctab := crc32.MakeTable(crc32.IEEE)

	calc := func() string {
		ts := time.Now()
		mod := ts.UnixNano() & 0xffff
		epoch := ts.Unix()
		fbv := fb.Next() & 0xffffffff
		str := fmt.Sprintf("%08x-%04x-%04x-%04x-%08x", epoch, tact, rnd.Intn(0x10000), mod, fbv)
		tact = (tact + 1) & 0xffff
		return fmt.Sprintf("%s-%08x", str, crc32.Checksum([]byte(str), crctab))
	}

	for {
		select {
		case stream <- calc():
		case <-ctx.Done():
			return
		}
	}
}

func (u *Uniq) Next() string {
	return <-u.next
}

func (u *Uniq) Close() {
	u.cancel()
}
