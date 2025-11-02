// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package contexts

import "github.com/labstack/echo/v4"

// AdminUserContext
type AdminUserContext struct {
	echo.Context
	User string
}

// GetAdminContext
func GetAdminContext(ctx echo.Context) *AdminUserContext {
	if vv, ok := ctx.(*AdminUserContext); ok {
		return vv
	}

	return &AdminUserContext{
		Context: ctx,
	}
}
