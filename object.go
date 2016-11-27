// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

// ------------ INLINE OBJECT
func parseInlineObject(line string) (string, bool) {
	var nLine string
	re := `(.*[\=\(\[]\s?)({)(\w+\s?:.+)(})(.*)`
	matches := regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) == 6 {
		if matches[3] != "" {
			nLine, _ = parseInlineObject(matches[3])
		}

		return matches[1] + "{ldelim}" + nLine + "{rdelim}" + matches[5], true
	}

	re = `(.*)({)([\w]|[^\}\$]+\s?:.+)(})(.*)`
	matches = regexp.MustCompile(re).FindStringSubmatch(line)

	if len(matches) != 6 {
		return line, false
	}

	if matches[1] != "" {
		nLine, _ = parseInlineObject(matches[1])
	}

	return nLine + "{ldelim}" + matches[3] + "{rdelim}" + matches[5], true
}
