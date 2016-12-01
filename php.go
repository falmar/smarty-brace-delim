// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

// ------------ SCRIPT TAGS
func startOfPHPTag(line string) bool {
	re := `{php}(.+({\/php}))?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "</php>"
}

func endOfPHPTag(line string) bool {
	re := `(({php).+)?{\/php}`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "{php"
}
