package redis


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06


import (
    r "github.com/garyburd/redigo/redis"
    "github.com/FZambia/go-sentinel"
    "time"
)


var sntnl *sentinel.Sentinel


func RedisPool( cfg *Config ) *r.Pool {

    if cfg.Failover {

        sntnl = &sentinel.Sentinel{
            Addrs:      cfg.Sentinel,
            MasterName: cfg.Cluster,
            Dial: func(addr string) (r.Conn, error) {
                timeout := 500 * time.Millisecond
                c, err := r.DialTimeout("tcp", addr, timeout, timeout, timeout)
                if err != nil {
                    return nil, err
                }
                return c, nil
            },
       }

    }

    return &r.Pool{
        MaxIdle:   100, 
        MaxActive: 1200,
        Wait:      true,
        Dial: func() (r.Conn, error) {
            
            addr := cfg.Addr

            if cfg.Failover {
                addr, _ = sntnl.MasterAddr()
            } 

            c, err := r.Dial("tcp", addr )
            
            if err != nil {
                return nil, err
            }
            
            return c, err
        },
    }    
}

