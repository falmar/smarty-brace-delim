// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"fmt"
	"regexp"
)

func startOfScriptTag(line string) bool {
	// one line
	re := `<script.+[^<\/script>]>.*[<](\/script>)?`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	if match != nil && len(match) == 2 && match[1] == "/script>" {
		fmt.Println(match[1])
		return false
	}

	// single line
	return regexp.MustCompile(`<script.+[^<\/script>]>`).MatchString(line)
}

func endOfScriptTag(line string) bool {
	re := `(<script)?.*<\/script>`
	match := regexp.MustCompile(re).FindStringSubmatch(line)

	return match != nil && len(match) == 2 && match[1] != "<script"
}
