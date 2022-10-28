// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package autostart

import (
	"os/user"
)

var (
	LoadAutostartDesktopFile = loadAutostartDesktopFile
	AutostartCmd             = autostartCmd
)

func MockUserCurrent(f func() (*user.User, error)) func() {
	origUserCurrent := userCurrent
	userCurrent = f
	return func() {
		userCurrent = origUserCurrent
	}
}

func MockCurrentDesktop(current string) func() {
	old := currentDesktop
	currentDesktop = splitSkippingEmpty(current, ':')
	return func() {
		currentDesktop = old
	}
}
