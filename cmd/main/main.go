package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mailslurper/mailslurper/cmd/mailslurper"
)

var (
	version = "1.14.1"
)

func main() {
	app := cli.NewApp()
	app.Usage = ""
	app.Version = version
	app.Commands = []*cli.Command{
		mailslurper.NewCommand(),
	}

	ctx := context.Background()
	err := app.RunContext(ctx, os.Args)
	if err != nil {
		slog.Error("main", slog.Group("fail",
			"vsn", version,
			"reason", err,
		))
		os.Exit(1)
	}
}
