// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "regexp"

func isRegExp(line string) bool {
	re := `(.*?[^\/]?)\/(.+)\/(.*)`

	return regexp.MustCompile(re).MatchString(line)
}
