package registerAndDiscover

import (
    "fmt"
    "context"
    etcd3 "go.etcd.io/etcd/clientv3"
    "google.golang.org/grpc/naming"
//    "github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// watcher is the implementaion of grpc.naming.Watcher
type watcher struct {
    re            *resolver // re: Etcd Resolver
    client        etcd3.Client
    isInitialized bool
}

// Close do nothing
func (w *watcher) Close() {
}

// micro-service discovery main function
// Next to return the updates
// data in etcd:  [key]=/zbw_naming/serviceName  [value]=relevant ip:port
// return: []*naming.Update, is the [value]=relevant ip:port of [key]=/zbw_naming/serviceName
func (w *watcher) Next() ([]*naming.Update, error) {
    // prefix is the etcd prefix/value to watch
    var Prefix = "zbw_naming"
    prefix := fmt.Sprintf("/%s/%s/", Prefix, w.re.serviceName)

    // if it is not initialized
    if !w.isInitialized {
        // query addresses from etcd
        // resp, is the [value]=ip:port of [key]=/zbw_naming/serviceName
        resp, err := w.client.Get(context.Background(), prefix, etcd3.WithPrefix())
        w.isInitialized = true
        if err == nil {
            addrs := extractAddrs(resp)
            //if not empty, return the updates or watcher new dir
            if l := len(addrs); l != 0 {
                updates := make([]*naming.Update, l)
                for i := range addrs {
                    updates[i] = &naming.Update{Op: naming.Add, Addr: addrs[i]}
                }
                return updates, nil
            }
        }
    }
	
    // generate etcd Watcher
    rch := w.client.Watch(context.Background(), prefix, etcd3.WithPrefix())
    for wresp := range rch {
        for _, ev := range wresp.Events {
            switch ev.Type {
            case mvccpb.PUT:
                return []*naming.Update{{Op: naming.Add, Addr: string(ev.Kv.Value)}}, nil
            case mvccpb.DELETE:
                return []*naming.Update{{Op: naming.Delete, Addr: string(ev.Kv.Value)}}, nil
            }
        }
    }
    return nil, nil
}

func extractAddrs(resp *etcd3.GetResponse) []string {
    addrs := []string{}

    if resp == nil || resp.Kvs == nil {
        return addrs
    }

    for i := range resp.Kvs {
        if v := resp.Kvs[i].Value; v != nil {
            addrs = append(addrs, string(v))
        }
    }

    return addrs
}