// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

// ------------ SCRIPT TAGS
func startOfLiteralTag(line string) bool {
	re := `{literal}(.+({\/literal}))?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "</script>"
}

func endOfLiteralTag(line string) bool {
	re := `(({literal).+)?{\/literal}`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "{literal"
}
