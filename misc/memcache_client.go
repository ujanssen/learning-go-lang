package main

import (
	"flag"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

func main() {
	var (
		host    = flag.String("host", "memcached", "memcached host or ip address")
		command = flag.String("command", "get", "get set, state")
		key     = flag.String("key", "key", "key")
		value   = flag.String("value", "value", "value")
	)
	flag.Parse()

	mc := memcache.New(*host + ":11211")
	mc.Timeout = time.Second

	switch *command {
	case "get":
		item, err := mc.Get(*key)
		if err == nil {
			fmt.Printf("got %s -> %s\n", *key, string(item.Value))
		} else {
			fmt.Printf("error %v\n", err)
		}
	case "set":
		item := memcache.Item{Key: *key, Value: []byte(*value)}
		err := mc.Set(&item)
		if err == nil {
			fmt.Printf("set %s -> %s\n", *key, *value)
		} else {
			fmt.Printf("error %v\n", err)
		}
	}
}
