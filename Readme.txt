aa.go
1) interface����������
2) go���ӱ���etcd������


bb.go
goroutine��context��ѧϰ
context.WithCancel��ʹ��


registerAndDiscoverĿ¼
�Լ�����ʵ��gRPC��ܵķ���ע��/����/���ؾ���
�ͻ��ˣ�ͨ��etcd��Get�ӿں�Watch�ӿڣ�resolver.go��watcher.go��
����ˣ�ͨ��etcd��Put�ӿں�KeepAlive�ӿڣ�register.go��
resolver.go: ������etcdClient��ʼ���ӵ�etcd
watcher.go: ����1)naming����������->URL�ķ��롣2)watch etcd��������Ⱥ�ı仯������1��list��������Ⱥ��ip list�����ؾ����ɿͻ��˴����Լ�ȥʵ�֣���������ѯ����������������
register.go: �����micro-service register/unregister��etcd����registerʱ����TTL����Ҫ������ά�ַ���ע�������(������δʵ��)��
���룺go install studyBasic\src\registerAndDiscover
�ο��ĵ�
https://segmentfault.com/a/1190000008672912?utm_source=tag-newest
https://segmentfault.com/a/1190000014501241
