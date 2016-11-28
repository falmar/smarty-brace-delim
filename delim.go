// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

// ------------ LDELIM
func parseLDelim(line string) (string, bool, bool) {
	if !strings.Contains(line, "{ldelim}") {
		return line, false, false
	}

	chars := []string{`"`, `'`, "`"}
	for _, c := range chars {
		p := `(.*?)([\` + c + `](.*?(\{ldelim\})?.*?)[\` + c + `])(.*)`
		matches := regexp.MustCompile(p).FindStringSubmatch(line)

		if matches != nil {
			l, lm, _ := parseLDelim(matches[1])
			r, rm, _ := parseLDelim(matches[5])

			return l + matches[2] + r, lm || rm, true
		}
	}

	return strings.Replace(line, "{ldelim}", "{", -1), true, false
}

// ------------ RDELIM
func parseRDelim(line string) (string, bool, bool) {
	if !strings.Contains(line, "{rdelim}") {
		return line, false, false
	}

	chars := []string{`"`, `'`, "`"}
	for _, c := range chars {
		p := `(.*?)([\` + c + `](.*?(\{rdelim\})?.*?)[\` + c + `])(.*)`
		matches := regexp.MustCompile(p).FindStringSubmatch(line)

		if matches != nil {
			l, lm, _ := parseRDelim(matches[1])
			r, rm, _ := parseRDelim(matches[5])

			return l + matches[2] + r, lm || rm, true
		}
	}

	return strings.Replace(line, "{rdelim}", "}", -1), true, false
}
