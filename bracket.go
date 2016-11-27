// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

// ------------ LEFT BRACKET
func parseLeftBracket(line string) (string, bool) {
	var nLine string
	re := `(.*)(\{)((.*)(\}))?(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) != 7 {
		return line, false
	}

	if matches[1] != "" {
		nLine, _ = parseLeftBracket(matches[1])
	}

	var matchDelim bool

	if matches[2] == "{" {
		if matches[5] != "}" {
			matchDelim = true
			nLine += "{ldelim}"
		} else if matches[5] == "}" && matches[4] == "" {
			matchDelim = true
			nLine += "{ldelim}"
		}
	}

	if !matchDelim {
		nLine += matches[2]
	}

	return nLine + matches[3] + matches[6], matchDelim
}

// ------------ RIGHT BRACKET
func parseRightBracket(line string) (string, bool) {
	var nLine string
	re := `(.*)((\{)(.*))(\})(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) == 7 {
		// take first regex actions

		if matches[1] != "" {
			nLine, _ = parseRightBracket(matches[1])
		}

		if matches[5] == "}" {
			hasLeft := matches[3] == "{"
			if hasLeft && (strings.Contains(matches[4], "}") || matches[4] == "") {
				nLine += matches[2] + "{rdelim}"
			} else if hasLeft {
				nLine += matches[2] + matches[5]
			}
		}

		return nLine + matches[6], true
	}

	re = `(.*)(})(.*)`
	matches = regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) != 4 {
		return line, false
	}

	if matches[1] != "" {
		nLine, _ = parseRightBracket(matches[1])
	}

	return nLine + "{rdelim}" + matches[3], true
}
