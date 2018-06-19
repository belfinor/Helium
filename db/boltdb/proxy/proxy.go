package proxy

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-01-24

import (
	"github.com/belfinor/Helium/log"
	"github.com/boltdb/bolt"
	"sync"
	"time"
)

type Proxy struct {
	sync.Mutex

	root     string
	bases    map[string]*DB
	cnt      int
	readonly bool
}

func New(root string, readonly bool) *Proxy {
	return &Proxy{
		cnt:      0,
		bases:    make(map[string]*DB),
		root:     root,
		readonly: readonly,
	}
}

func (p *Proxy) Open(name string) *DB {
	p.Lock()

	v, h := p.bases[name]
	if h {
		p.Unlock()
		v.Lock()
		return v
	}

	defer p.Unlock()

	db := &DB{
		proxy: p,
		name:  name,
	}

	opts := &bolt.Options{Timeout: time.Second, ReadOnly: p.readonly}

	var e error
	db.db, e = bolt.Open(p.root+name, 0644, opts)
	if e != nil {
		log.Error(e.Error())
		return nil
	}

	db.Lock()
	p.bases[name] = db

	return db
}

func (p *Proxy) Close() {
	p.Lock()
	b := p.bases
	p.bases = make(map[string]*DB)
	p.Unlock()

	for _, v := range b {
		v.Lock()
		v.db.Close()
	}
}
