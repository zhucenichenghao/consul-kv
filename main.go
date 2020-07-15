package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

var (
	consulAddr string
	key string
	waitIndex uint64
	waitTime int
)

func main() {
	flag.StringVar(&consulAddr, "consul", "127.0.0.1:8500", "consul addr")
	flag.StringVar(&key, "key", "", "key")
	flag.Uint64Var(&waitIndex, "waitIndex", 0, "key")
	flag.IntVar(&waitTime, "waitTime", 5, "key")

	flag.Parse()

	fmt.Printf("waitIndex:%v waitTime:%v\n", waitIndex, waitTime)

	c, err := api.NewClient(&api.Config{Address: consulAddr, Scheme: "http"})
	if err != nil {
		fmt.Printf("get client error:%v\n", err)
	}
	q := &api.QueryOptions{RequireConsistent: true, WaitIndex: waitIndex, WaitTime: time.Duration(waitTime) * time.Second}
	kvpairs, meta, err := c.KV().List(key, q)
	if err != nil {
		fmt.Printf("get kv error:%v\n", err)
	} else {
		for k, v := range kvpairs {
			fmt.Printf("kvpairs--k:%+v v:%+v value:%s\n", k, v, string(v.Value))
		}

		fmt.Printf("meta:%+v\n", *meta)
	}
}