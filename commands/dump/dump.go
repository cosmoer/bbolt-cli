package dump

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cosmoer/bbolt-cli/boltutils"
	"github.com/cosmoer/bbolt-cli/schema"
	"github.com/urfave/cli"
	bolt "go.etcd.io/bbolt"
)

var Command = cli.Command{
	Name:      "dump",
	Usage:     "dump all bucket and key/value.",
	ArgsUsage: "<boltdb file>",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "schema",
			Value: "containerd",
			Usage: "boltdb schema: containerd,default. default is containerd.",
		},
	},
	Action: func(context *cli.Context) error {
		SrcPath := context.Args().Get(0)
		if SrcPath == "" {
			return errors.New("boltdb file need to be specified")
		}
		// Ensure boltdb file exists.
		_, err := os.Stat(SrcPath)
		if os.IsNotExist(err) {
			return errors.New("boltdb file is not exist")
		} else if err != nil {
			return err
		}

		// Open bolt database.
		src, err := bolt.Open(SrcPath, 0444, nil)
		if err != nil {
			return err
		}
		defer src.Close()

		sc := context.String("schema")
		if ret := strings.Compare(sc, schema.Containerd); ret == 0 {
			err = containerdMetaPrintAll(src)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func containerdMetaPrintAll(src *bolt.DB) error {
	parser := schema.NewContainerdMetaParser()

	if err := boltutils.Walk(src, func(keys [][]byte, k, v []byte, seq uint64) error {
		path, key, value, err := parser.Parse(keys, k, v)
		if err != nil {
			fmt.Errorf("parse failed. key:%s\n", string(k))
			return err
		}
		if v == nil {
			fmt.Printf("%s,%s\n", path, key)
		} else {
			fmt.Printf("%s,%s=%s\n", path, key, value)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
