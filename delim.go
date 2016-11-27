// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "strings"

// ------------ LDELIM
func parseLDelim(line string) (string, bool) {
	if !strings.Contains(line, "{ldelim}") {
		return line, false
	}

	return strings.Replace(line, "{ldelim}", "{", -1), true
}

// ------------ RDELIM
func parseRDelim(line string) (string, bool) {
	if !strings.Contains(line, "{rdelim}") {
		return line, false
	}

	return strings.Replace(line, "{rdelim}", "}", -1), true
}
