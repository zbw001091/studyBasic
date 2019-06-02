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
编译：go install studyBasic\src\registerAndDiscover
参考文档
https://segmentfault.com/a/1190000008672912?utm_source=tag-newest
https://segmentfault.com/a/1190000014501241
