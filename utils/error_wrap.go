/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package utils

import (
	"fmt"
)

// WrapErr gives error some context msg
// returns nil if err is nil
func WrapErr(err error, msg string) (errWithContext error) {
	if err == nil {
		return
	}

	errWithContext = fmt.Errorf("%s: %v", msg, err)
	return
}

// WrapfErr gives error some context msg
// with desired format and content
// returns nil if err is nil
func WrapfErr(err error, format string, a ...interface{}) (errWithContext error) {
	if err == nil {
		return
	}

	errWithContext = WrapErr(err, fmt.Sprintf(format, a...))
	return
}
