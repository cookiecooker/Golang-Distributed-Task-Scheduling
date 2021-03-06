package main

import (
	"os/exec" 
	"fmt"
)

import (
	"github.com/cores/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.client
		err error
	)

	config = clientv3.Config{
		Endpoints: []string{"36.111.184.221:2379"},
		DialTimeout: 5 * time.Second
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	client = client
}