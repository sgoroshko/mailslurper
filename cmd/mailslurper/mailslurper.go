// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurper

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/mailslurper/mailslurper/pkg/mailslurper"
	"github.com/mailslurper/mailslurper/pkg/ui"
)

var config *mailslurper.Configuration
var database mailslurper.IStorage
var logger *logrus.Entry
var renderer *ui.TemplateRenderer
var mailItemChannel chan *mailslurper.MailItem
var smtpListenerContext context.Context
var smtpListenerCancel context.CancelFunc
var smtpListener *mailslurper.SMTPListener
var connectionManager *mailslurper.ConnectionManager
var cacheService *cache.Cache

var admin *echo.Echo
var service *echo.Echo

// var logFormat = flag.String("logformat", "simple", "Format for logging. 'simple' or 'json'.")
// var logLevel = flag.String("loglevel", "info", "Level of logs to write. Valid values are 'debug', 'info', or 'error'.")
// var configFile = flag.String("config", "config.json", "Absolute location of the config.json. Default is 'config.json' in the current directory")
var logFormat = "simple"
var logLevel = "info"
var configFile = "config.json"

// Run
func Run(ctx context.Context, vsn string, cfg Config) error {
	flag.Parse()

	logger = mailslurper.GetLogger(logLevel, logFormat, "MailSlurper")
	logger.Infof("Starting MailSlurper Server v%s", vsn)

	configFile = cfg.ConfigFile
	setupConfig(configFile)
	config.SMTPListen = cfg.SMTPListen
	config.ServiceListen = cfg.APIListen
	config.ServicePublic = cfg.APIPublic
	config.WWWListen = cfg.AdminListen
	config.WWWPublic = cfg.AdminPublic
	config.AutoStartBrowser = cfg.StartBrowser
	
	logger.Infof("Using configuration from %s", configFile)

	if err := config.Validate(); err != nil {
		logger.WithError(err).Fatalf("Invalid configuration")
	}

	renderer = ui.NewTemplateRenderer()
	cacheService = cache.New(time.Minute*time.Duration(config.AuthTimeoutInMinutes), time.Minute*time.Duration(config.AuthTimeoutInMinutes))

	setupDatabase()
	setupSMTP()
	setupAdminListener()
	setupServicesListener()

	defer database.Disconnect()

	if config.AutoStartBrowser {
		ui.StartBrowser(config, logger)
	}

	/*
	 * Block this thread until we get an interrupt signal. Once we have that
	 * start shutting everything down
	 */
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()

	ctx, cancel = context.WithTimeout(smtpListenerContext, 20*time.Second)
	defer cancel()

	smtpListenerCancel()

	if err := admin.Shutdown(ctx); err != nil {
		logger.Fatalf("Error shutting down admin listener: %s", err.Error())
	}

	if err := service.Shutdown(ctx); err != nil {
		logger.Fatalf("Error shutting down service listener: %s", err.Error())
	}

	return nil
}
