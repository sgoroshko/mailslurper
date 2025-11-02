package mailslurper

import (
	"github.com/urfave/cli/v2"
)

// NewCommand
func NewCommand() *cli.Command {
	cfg := new(Config)
	return &cli.Command{
		Name:  "server",
		Flags: bindConfig(cfg),
		Action: func(c *cli.Context) error {
			return Serve(c, *cfg)
		},
	}
}

// Serve
func Serve(cc *cli.Context, cfg Config) error {
	return Run(cc.Context, cc.App.Version, cfg)
}
