// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package ui

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/mailslurper/mailslurper/pkg/mailslurper"
	"github.com/skratchdot/open-golang/open"
)

// StartBrowser
func StartBrowser(config *mailslurper.Configuration, logger *logrus.Entry) {
	timer := time.NewTimer(time.Second)

	go func() {
		<-timer.C
		logger.Infof("opening web browser to %s", config.GetPublicWWWURL())
		err := open.Start(config.GetPublicWWWURL())
		if err != nil {
			logger.Infof("ERROR - Could not open browser - %s", err.Error())
		}
	}()
}
