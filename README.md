# bbolt-cli
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cosmoer/bbolt-cli)](https://pkg.go.dev/github.com/cosmoer/bbolt-cli)

bbolt-cli是一个读取boltdb数据库的命令行工具。目前主要针对containerd的元数据进行解析并导出，可以用来观察containerd的元数据组织形式。



## 安装

```
go install github.com/cosmoer/bbolt-cli
```



## 使用示例

```
root@vm:/# bbolt-cli dump metadata.db 
/,v1
/v1,parents
/v1/parents,10 17=default/26/busybox
/v1,snapshots
/v1/snapshots,default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,createdat=2022-02-13 15:23:13.382463192 +0000 UTC
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,id=10
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,inodes=24
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,kind=Committed
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,size=1343488
/v1/snapshots/default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034,updatedat=2022-02-13 15:23:13.382463192 +0000 UTC
/v1/snapshots,default/26/busybox
/v1/snapshots/default/26/busybox,createdat=2022-02-28 15:47:04.95761334 +0000 UTC
/v1/snapshots/default/26/busybox,id=17
/v1/snapshots/default/26/busybox,kind=Active
/v1/snapshots/default/26/busybox,parent=default/15/sha256:ad6a27b3fb1554337b257ddb2d2e0061e3955e653b672000dcb73e2545912034
/v1/snapshots/default/26/busybox,updatedat=2022-02-28 15:47:04.95761334 +0000 UTC
root@vm:/home/bbolt-cli#
```

