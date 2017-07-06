package redis


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-07-06


type Config struct {
    Addr     string   `json:"addr"`
    Failover bool     `json:"failover"`
    Cluster  string   `json:"cluster"`
    Sentinel []string `json:"sentinel"`
}

