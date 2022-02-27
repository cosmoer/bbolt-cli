package dump

import (
	"errors"

	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:      "dump",
	Usage:     "dump all bucket and key/value.",
	ArgsUsage: "<boltdb file>",
	Action: func(context *cli.Context) error {
		SrcPath := context.Args().Get(0)
		if SrcPath == "" {
			return errors.New("boltdb file need to be specified")
		}

		return nil
	},
}
