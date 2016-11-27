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
	`r2 regex /(.*?[^\/])\/(.+)\/(.*)/`,
	`commen ml start /(.*)\/\*\*(.*) (.*)\*\/(.*)/`,
	`commen ml start /(.*)\{\*\*(.*) (.*)\*}(.*)/ smarty`,
	`line comment /(.*)(\/\/.*)/ [1]check [2]append`,
	`partial comment /(.*)(\/\*.*\*\/)(.*)/`,
}

var expectRegExpPatterns = [][]string{
	[]string{`return exec(`, `/^[0-9]{11}$/`, `, value)`},
	[]string{`return exec(`, `/^[0-9]{2}$/`, `, value)`},
	[]string{`return exec(`, `/^[a-zA-Z]{1,2}[0-9]{2,3}$/`, `, value)`},
	[]string{`return exec(`, `/^[0-9]{7,10}$/`, `, value)`},
	[]string{`return exec(`, `/^\w{6}$/`, `, value)`},
	[]string{`r2 regex `, `/(.*?[^\/])\/(.+)\/(.*)/`, ``},
	[]string{`commen ml start `, `/(.*)\/\*\*(.*) (.*)\*\/(.*)/`, ``},
	[]string{`commen ml start `, `/(.*)\{\*\*(.*) (.*)\*}(.*)/`, ` smarty`},
	[]string{`line comment `, `/(.*)(\/\/.*)/`, ` [1]check [2]append`},
	[]string{`partial comment `, `/(.*)(\/\*.*\*\/)(.*)/`, ``},
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

func TestRegExpSplit(t *testing.T) {
	for i, p := range regExpPatterns {

		result, match := parseRegExp(p)

		if !match {
			t.Fatalf("Should match RegExp to be: %s", p)
		}

		for z, ep := range expectRegExpPatterns[i] {
			if ep != result[z] {
				t.Fatalf("Expected RegExp to be: %s; got: %s", ep, result[z])
			}
		}

	}
}
