// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

// ------------ SCRIPT TAGS
var openPHPTags = []string{
	`{php} random test`,
	`rando{ñm test} {php}`,
	`  yo-!    {php}      // random comment`,
}

var closePHPTags = []string{
	`anything {/php} //  comment`,
	`{/php}`,
	`  yo-!    {/php}       // random comment`,
}

var nonPHPTags = []string{
	`<!doctype html>`,
	`<html class="no-js" lang="en">`,
	`<head>`,
	`<meta charset="utf-8" />`,
	`<meta http-equiv="x-ua-compatible" content="ie=edge">`,
	`<meta name="viewport" content="width=device-width, initial-scale=1.0" />`,
	`<title>React-App</title>`,
	`<link rel="stylesheet" href="/css/pace.css">`,
	`<link rel="stylesheet" href="/css/app.min.css">`,
	`</head>`,
	`<body>`,
	`<div id="app"></div>`,
	`<script src="js/pace.min.js"></script>`,
	`<script src="app.min.js"></script>`,
	`<script src="app.min.js"> function() {}</script>`,
	`</body>`,
	`</html>`,
	`$.fn.serializeObject = function () {ldelim}`,
	`var o = {};`,
	`var a = this.serializeArray();`,
	`$.each(a, function () {ldelim}`,
	`if (o[this.name] !== undefined) {ldelim}`,
	`if (!o[this.name].push) {ldelim}`,
	`o[this.name] = [o[this.name]];`,
	`{rdelim}`,
	`o[this.name].push(this.value || '');`,
	`{rdelim} else {ldelim}`,
	`o[this.name] = this.value || '';`,
	`{rdelim}`,
	`{rdelim});`,
	`return o;`,
	`{php}{/php}`,
}

func TestStartPHPTags(t *testing.T) {
	for _, line := range openLiteralTags {
		if !startOfLiteralTag(line) {
			t.Fatalf("Should be php tag %s", line)
		}
	}
}

func TestStartPHPTagNonTags(t *testing.T) {
	for _, line := range closeLiteralTags {
		if startOfLiteralTag(line) {
			t.Fatalf("Should not be php tag %s", line)
		}
	}
}

func TestEndPHPTags(t *testing.T) {
	for _, line := range closeLiteralTags {
		if !endOfLiteralTag(line) {
			t.Fatalf("Should be php tag %s", line)
		}
	}
}

func TestEndPHPTagNonTags(t *testing.T) {
	for _, line := range nonLiteralTags {
		if endOfLiteralTag(line) {
			t.Fatalf("Should not be php tag %s", line)
		}
	}
}
