package registerAndDiscover

import (
    "fmt"
    "log"
    "strings"
    "time"
    "context"
    
    etcd3 "go.etcd.io/etcd/clientv3"
    "go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

// Prefix should start and end with no slash
var Prefix = "zbw_naming"
var client etcd3.Client
var serviceKey string

var stopSignal = make(chan bool, 1)

// Register
// target, is the url of etcd
// ttl, effective lifetime of one record in etcd
func Register(name string, host string, port int, target string, interval time.Duration, ttl int) error {
	// micro-service name resolution k/v, to be set into etcd
    serviceValue := fmt.Sprintf("%s:%d", host, port)
    serviceKey = fmt.Sprintf("/%s/%s/%s", Prefix, name, serviceValue)
	
    // get endpoints for register dial address
    var err error
    client, err := etcd3.New(etcd3.Config{
        Endpoints: strings.Split(target, ","),
    })
    if err != nil {
        return fmt.Errorf("grpclb: create etcd3 client failed: %v", err)
    }
	
    go func() {
        // invoke self-register with ticker
        ticker := time.NewTicker(interval)
        for {
            // minimum lease TTL is ttl-second
            resp, _ := client.Grant(context.TODO(), int64(ttl))
            // should get first, if not exist, set it
            _, err := client.Get(context.Background(), serviceKey)
            if err != nil {
                if err == rpctypes.ErrKeyNotFound {
                    if _, err := client.Put(context.TODO(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
                        log.Printf("grpclb: set service '%s' with ttl to etcd3 failed: %s", name, err.Error())
                    }
                } else {
                    log.Printf("grpclb: service '%s' connect to etcd3 failed: %s", name, err.Error())
                }
            } else {
                // refresh set to true for not notifying the watcher
                if _, err := client.Put(context.Background(), serviceKey, serviceValue, etcd3.WithLease(resp.ID)); err != nil {
                    log.Printf("grpclb: refresh service '%s' with ttl to etcd3 failed: %s", name, err.Error())
                }
            }
            select {
            case <-stopSignal:
                return
            case <-ticker.C:
            }
        }
    }()

    return nil
}

// UnRegister delete registered service from etcd
func UnRegister() error {
    stopSignal <- true
    stopSignal = make(chan bool, 1) // just a hack to avoid multi UnRegister deadlock
    var err error;
    if _, err := client.Delete(context.Background(), serviceKey); err != nil {
        log.Printf("grpclb: deregister '%s' failed: %s", serviceKey, err.Error())
    } else {
        log.Printf("grpclb: deregister '%s' ok.", serviceKey)
    }
    return err
}