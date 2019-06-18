package etcd

import (
	etcdc "go.etcd.io/etcd/clientv3"
)

type EtcdSource struct {
	inner *etcdc.Client
}
