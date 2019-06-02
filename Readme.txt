aa.go
1) interface的入门例子
2) go连接本地etcd的例子


bb.go
goroutine和context的学习
context.WithCancel的使用


registerAndDiscover目录
自己动手实现gRPC框架的服务注册/发现/负载均衡
客户端：通过etcd的Get接口和Watch接口（resolver.go和watcher.go）
服务端：通过etcd的Put接口和KeepAlive接口（register.go）
resolver.go: 负责建立etcdClient初始连接到etcd
watcher.go: 负责，1)naming解析，名称->URL的翻译。2)watch etcd监听服务集群的变化。返回1个list，即服务集群的ip list。负载均衡由客户端代码自己去实现，可以用轮询，或者其他方法。
register.go: 负责把micro-service register/unregister到etcd，且register时带有TTL，需要心跳来维持服务注册的续命(心跳暂未实现)。
编译：go install studyBasic\src\registerAndDiscover
参考文档
https://segmentfault.com/a/1190000008672912?utm_source=tag-newest
https://segmentfault.com/a/1190000014501241
