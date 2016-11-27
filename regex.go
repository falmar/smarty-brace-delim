// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

func isRegExp(line string) bool {
	re := `(.*?[^\/]?)\/(.+)\/(.*)`

	return regexp.MustCompile(re).MatchString(line)
}

func parseRegExp(line string) ([]string, bool) {
	re := `(.*?[^\/]?)(\/.+\/)(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if matches == nil {
		return nil, false
	}

	return []string{matches[1], matches[2], matches[3]}, true
}
