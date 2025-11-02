package mailslurper

import (
	"github.com/mailslurper/mailslurper/pkg/storage"
	"github.com/urfave/cli/v2"
)

// Config
type Config struct {
	Debug        bool           //
	ConfigFile   string         //
	Storage      storage.Config // TODO
	SMTPListen   string         //
	APIListen    string         //
	APIPublic    string         //
	AdminListen  string         //
	AdminPublic  string         //
	StartBrowser bool           //
}

func bindConfig(cfg *Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "use debug mode",
			Destination: &cfg.Debug,
			Value:       false,
		},
		&cli.StringFlag{
			Name:        "config",
			Usage:       "configuration file",
			Destination: &cfg.ConfigFile,
			Value:       "config.json",
		},
		&cli.StringFlag{
			Name:        "storage.dsn",
			Usage:       "database connection string",
			Destination: &cfg.Storage.DSN,
			Value:       "sqlite3:///mailslurper.db",
			Hidden:      true,
		},
		&cli.BoolFlag{
			Name:        "browser.start",
			Usage:       "auto-start browser",
			EnvVars:     []string{"BROWSER_START"},
			Destination: &cfg.StartBrowser,
			Value:       false,
		},
		&cli.StringFlag{
			Name:        "smtp.ln",
			Usage:       "SMTP server listen address",
			EnvVars:     []string{"SMTP_LISTEN"},
			Destination: &cfg.SMTPListen,
			Value:       "0.0.0.0:2500",
		},
		&cli.StringFlag{
			Name:        "api.ln",
			Usage:       "API server listen address",
			EnvVars:     []string{"API_LISTEN"},
			Destination: &cfg.APIListen,
			Value:       "0.0.0.0:8085",
		},
		&cli.StringFlag{
			Name:        "api.pub",
			Usage:       "API server public address",
			EnvVars:     []string{"API_PUBLIC"},
			Destination: &cfg.APIPublic,
			Value:       "localhost:8085",
		},
		&cli.StringFlag{
			Name:        "admin.ln",
			Usage:       "admin panel listen address",
			EnvVars:     []string{"ADMIN_LISTEN"},
			Destination: &cfg.AdminListen,
			Value:       "0.0.0.0:8080",
		},
		&cli.StringFlag{
			Name:        "admin.pub",
			Usage:       "admin panel public address",
			EnvVars:     []string{"ADMIN_PUBLIC"},
			Destination: &cfg.AdminPublic,
			Value:       "localhost:8080",
		},
	}
}
