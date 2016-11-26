// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"regexp"
	"strings"
)

// ------------ SCRIPT TAGS

func startOfScriptTag(line string) bool {
	re := `<script(.+(<\/script>))?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "</script>"
}

func endOfScriptTag(line string) bool {
	re := `((<script).+)?<\/script>`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "<script"
}

// ------------ LEFT BRACKET
func isLeftBracket(line string) bool {
	re := `(\{)(.*(\}))?`
	matches := regexp.MustCompile(re).FindAllStringSubmatch(line, -1)

	if matches == nil {
		return false
	}

	match := matches[len(matches)-1]

	return match != nil && len(match) == 4 && match[3] != "}"
}

func parseLeftBracket(line string) string {
	var nLine string
	re := `(.*)(\{)(.*(}))?(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) != 6 {
		return line
	}

	if matches[1] != "" {
		nLine = parseLeftBracket(matches[1])
	}

	var matchDelim bool

	if matches[2] == "{" {
		if matches[4] != "}" {
			matchDelim = true
			nLine += "{ldelim}"
		}
	}

	if !matchDelim {
		nLine += matches[2]
	}

	return nLine + matches[3] + matches[5]
}

// ------------ RIGHT BRACKET
func isRightBracket(line string) bool {
	re := `((\{).+)?\}`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 3 && match[2] != "{"
}

func parseRightBracket(line string) string {
	var nLine string
	re := `(.*)((\{)(.*))(\})(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) == 7 {
		// take first regex actions

		if matches[1] != "" {
			nLine = parseRightBracket(matches[1])
		}

		if matches[5] == "}" {
			hasLeft := matches[3] == "{"

			if hasLeft && strings.Contains(matches[4], "}") {
				nLine += matches[2] + "{rdelim}"
			} else {
				nLine += matches[2] + matches[5]
			}
		}

		return nLine + matches[6]
	}

	re = `(.*)(})(.*)`
	matches = regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) != 4 {
		return line
	}

	if matches[1] != "" {
		nLine = parseRightBracket(matches[1])
	}

	return nLine + "{rdelim}" + matches[3]
}
