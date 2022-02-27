//Package schema
//根据containerd的metadata的database schema来解析containerd服务的元数据文件信息
//database schema: https://pkg.go.dev/github.com/containerd/containerd/metadata

package schema

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	Containerd = "containerd"
	Default    = "default"
)

// do not edit, copy from containerd-1.5.9/metadata/buckets.go
var (
	schemaVersion             = "v1"
	bucketKeyVersion          = []byte(schemaVersion)
	bucketKeyDBVersion        = []byte("version")    // stores the version of the schema
	bucketKeyObjectLabels     = []byte("labels")     // stores the labels for a namespace.
	bucketKeyObjectImages     = []byte("images")     // stores image objects
	bucketKeyObjectContainers = []byte("containers") // stores container objects
	bucketKeyObjectSnapshots  = []byte("snapshots")  // stores snapshot references
	bucketKeyObjectContent    = []byte("content")    // stores content references
	bucketKeyObjectBlob       = []byte("blob")       // stores content links
	bucketKeyObjectIngests    = []byte("ingests")    // stores ingest objects
	bucketKeyObjectLeases     = []byte("leases")     // stores leases

	bucketKeyDigest      = []byte("digest")
	bucketKeyMediaType   = []byte("mediatype")
	bucketKeySize        = []byte("size")
	bucketKeyImage       = []byte("image")
	bucketKeyRuntime     = []byte("runtime")
	bucketKeyName        = []byte("name")
	bucketKeyParent      = []byte("parent")
	bucketKeyChildren    = []byte("children")
	bucketKeyOptions     = []byte("options")
	bucketKeySpec        = []byte("spec")
	bucketKeySnapshotKey = []byte("snapshotKey")
	bucketKeySnapshotter = []byte("snapshotter")
	bucketKeyTarget      = []byte("target")
	bucketKeyExtensions  = []byte("extensions")
	bucketKeyCreatedAt   = []byte("createdat")
	bucketKeyExpected    = []byte("expected")
	bucketKeyRef         = []byte("ref")
	bucketKeyExpireAt    = []byte("expireat")

	deprecatedBucketKeyObjectIngest = []byte("ingest") // stores ingest links, deprecated in v1.2
)

// copy from containerd-1.5.9/metadata/boltutil/helpers.go
var (
	bucketKeyAnnotations = []byte("annotations")
	bucketKeyLabels      = []byte("labels")
	bucketKeyUpdatedAt   = []byte("updatedat")
)

// copy from containerd-1.5.9/snapshots/storage/bolt.go
var (
	bucketKeyStorageVersion = []byte("v1")
	bucketKeySnapshot       = []byte("snapshots")
	bucketKeyParents        = []byte("parents")

	bucketKeyID     = []byte("id")
	bucketKeyKind   = []byte("kind")
	bucketKeyInodes = []byte("inodes")
)

// definitions of snapshot kinds
const (
	KindUnknown uint8 = iota
	KindView
	KindActive
	KindCommitted
)

func toKindString(k uint8) string {
	switch k {
	case KindView:
		return "View"
	case KindActive:
		return "Active"
	case KindCommitted:
		return "Committed"
	}

	return "Unknown"
}

// ContainerdMeta represents the containerd meta object
type ContainerdMeta struct {
	dbfile string
}

// NewContainerdMetaParser returns a ContainerdMetaParser.
func NewContainerdMetaParser() *ContainerdMeta {
	return &ContainerdMeta{}
}

// Parse executes parse
func (cm *ContainerdMeta) Parse(keys [][]byte, k, v []byte) (retPath, retKey, retValue string, err error) {
	strbuilder := strings.Builder{}
	for _, key := range keys {
		strbuilder.WriteString("/")
		strbuilder.WriteString(string(key))
	}
	if keys == nil {
		strbuilder.WriteString("/")
	}
	retValue = "[not support]"
	if strings.HasSuffix(strbuilder.String(), string(bucketKeyLabels)) {
		retValue = string(v)
	}
	retKey = string(k)
	if strings.HasSuffix(strbuilder.String(), string(bucketKeyParents)) {
		parent, num := binary.Uvarint(k)
		if num <= 0 {
			return "", "", "", errors.New("parse parents key parent failed")
		}
		child, num := binary.Uvarint(k[num+1:])
		if num <= 0 {
			return "", "", "", errors.New("parse parents key child failed")
		}
		retKey = fmt.Sprintf("%d %d", parent, child)
		retValue = string(v)
	}

	switch string(k) {
	case string(bucketKeyCreatedAt), string(bucketKeyUpdatedAt):
		var t time.Time
		if v != nil {
			if err := t.UnmarshalBinary(v); err != nil {
				return "", "", "", err
			}
		}
		retValue = t.String()
	case string(bucketKeySize), string(bucketKeyInodes):
		size, _ := binary.Varint(v)
		retValue = fmt.Sprintf("%d", size)
	case string(bucketKeyKind):
		if len(v) == 1 {
			retValue = toKindString(v[0])
		}
	case string(bucketKeyID):
		id, _ := binary.Uvarint(v)
		retValue = fmt.Sprintf("%d", id)
	case string(bucketKeyDigest), string(bucketKeyMediaType), string(bucketKeyName), string(bucketKeyDBVersion), string(bucketKeyParent):
		retValue = string(v)
	}

	retPath = strbuilder.String()
	if v == nil {
		retValue = ""
	}

	return retPath, retKey, retValue, nil
}
