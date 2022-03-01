module github.com/cosmoer/bbolt-cli

go 1.17

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/urfave/cli v1.22.5
	go.etcd.io/bbolt v1.3.6
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
)

retract (
	v0.0.0-20220301140752-59b71bcbf3a0
	v0.0.0-20220301064236-ced2e67245fa
	v0.0.0-20220228160738-3011b02d8b8c
)
