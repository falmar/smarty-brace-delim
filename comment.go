// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

func isCommentLine(line string) bool {
	return regexp.MustCompile(`(.*)\/\/(.*)`).MatchString(line)
}

func parseCommentLine(line string) ([]string, bool) {
	var nLine string
	var rc string
	re := regexp.MustCompile(`(.*)(\/\/.*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		return []string{nLine, line}, false
	}

	if matches[1] != "" {
		lMatches, matched := parseCommentLine(matches[1])

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

// -------------- single Multiline

func isMultilineCommentStart(line string, single bool) bool {
	p := `(.*)\/\*(.*)`

	if !single {
		p = `(.*)\{\*(.*)`
	}

	return regexp.MustCompile(p).MatchString(line)
}

func parseMultilineCommentStart(line string, single bool) ([]string, bool) {
	var nLine string
	var rc string
	p := `(.*)(\/\*\*.*)`

	if !single {
		p = `(.*)(\{\*.*)`
	}

	re := regexp.MustCompile(p)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		return []string{nLine, line}, false
	}

	if matches[1] != "" {
		lMatches, matched := parseMultilineCommentStart(matches[1], single)

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

func isMultilineCommentEnd(line string, single bool) bool {
	p := `(.*)\*\/(.*)`

	if !single {
		p = `(.*)\*\}(.*)`
	}

	return regexp.MustCompile(p).MatchString(line)
}

func parseMultilineCommentEnd(line string, single bool) ([]string, bool) {
	var nLine string
	var rc string
	p := `(.*\*\/)(.*)`

	if !single {
		p = `(.*\*\})(.*)`
	}

	re := regexp.MustCompile(p)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		return []string{nLine, line}, false
	}

	if matches[1] != "" {
		lMatches, matched := parseMultilineCommentStart(matches[1], single)

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
