// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

func isLineComment(line string) bool {
	return regexp.MustCompile(`(.*)\/\/(.*)`).MatchString(line)
}

func parseLineComment(line string) ([]string, bool) {
	var nLine string
	var rc string
	re := regexp.MustCompile(`(.*)(\/\/.*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		return []string{nLine, line}, false
	}

	if matches[1] != "" {
		lMatches, matched := parseLineComment(matches[1])

		nLine = lMatches[1]
		rc = matches[2]

		if matched {
			nLine = lMatches[0]
			rc = lMatches[1] + matches[2]
		}

		return []string{nLine, rc}, true
	}

	return []string{nLine, matches[2]}, true
}
