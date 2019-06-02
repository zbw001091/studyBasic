package main

import (
	"fmt"
	// "github.com/coreos/go-etcd/etcd"
	// "github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
	"time"
	"context"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {
}

func (iPhone ApplePhone) call() {
	fmt.Println("I am Apple Phone, I can call you!")
}

func main1() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(ApplePhone)
	phone.call()

	//    machines := []string{"http://127.0.0.1:2379"}
	//    client := etcd.NewClient(machines)
	//    if _, err := client.Set("/foo", "bar", 0); err != nil {
	//        log.Fatal(err)
	//    }
	//    fmt.Println(client.Get("/foo"))
	
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
        fmt.Println("etcd connect failed, err:", err)
        return
    }
	fmt.Println("connect succ")
    defer cli.Close()
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    _, err = cli.Put(ctx, "/logagent/conf/", "sample_value")
    cancel()
    if err != nil {
        fmt.Println("put failed, err:", err)
        return
    }
    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    resp, err := cli.Get(ctx, "/logagent/conf/")
    cancel()
    if err != nil {
        fmt.Println("get failed, err:", err)
        return
    }
    for _, ev := range resp.Kvs {
        fmt.Printf("%s : %s\n", ev.Key, ev.Value)
    }
}
