// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

var regExpPatterns = []string{
	`return exec(/^[0-9]{11}$/, value)`,
	`return exec(/^[0-9]{2}$/, value)`,
	`return exec(/^[a-zA-Z]{1,2}[0-9]{2,3}$/, value)`,
	`return exec(/^[0-9]{7,10}$/, value)`,
	`return exec(/^\w{6}$/, value)`,
	`/^[0-9]{11}$/`,
	`/^[0-9]{2}$/`,
	`/^[a-zA-Z]{1,2}[0-9]{2,3}$/`,
	`/^[0-9]{7,10}$/`,
	`/^\w{6}$/`,
}

var nonRegExpPatterns = []string{
	`(.*)(\{)(.+(}))?(.*)(.+)?({)$ ldelim | $1{ldelim}`,
	`(.*)({)(.*)(})(.*)(\s+)(}) rdelim | $1{rdelim}`,
	`({)([^ldelim\n]\w+:\s?\"?\'?.+\"?\'?[^ldelim\n])(}) `,
	`rc1 (.*)(({).*)(\})(.*)`,
	`r2 (.*[\=\(\[\{]\s?({))(\w+\s*.+:\s*.+)(})`,
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`funcion () {`,
	`}`,
	`call({`,
	`hello: "world"`,
	`}, {`,
	`world: "hello"`,
	`})`,
	`let array = [{`,
	`hello: "world",`,
	`myObject:{`,
	`one: 1,`,
	`two: [2, 2]`,
	`}`,
	`}]`,
	`const strangeObject = {maybe: {it: {wont: {work: "?"`,
	`}, maybe: ""}, not: "did"}, work: "entirely"}`,
	`</script>`,
	`</body>`,
}

func TestIsRegExpMatch(t *testing.T) {
	for _, p := range regExpPatterns {
		if !isRegExp(p) {
			t.Fatalf("Should be a RegExp Pattern %s", p)
		}
	}
}

func TestIsRegExpNoMatch(t *testing.T) {
	for _, p := range nonRegExpPatterns {
		if isRegExp(p) {
			t.Fatalf("Should not be a RegExp Pattern %s", p)
		}
	}
}
