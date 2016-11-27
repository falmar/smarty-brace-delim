// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

var lineComments = []string{
	`function () { // Copyright 2016 David Lavieri. All rights reserved.`,
	`some code // Use of this source code is governed by a MIT License`,
	`// License that can be found in the LICENSE file.`,
	`<script type="text/javascript"> // Date: 0/0/0`,
	`console.log({include// file=$myCustomFile}) // Time: 0:0 PM`,
	`</script> // @author    David Lavieri (falmar) <daviddlavier@gmail.com>`,
	`// @copyright 2016 David Lavieri`,
	`let some = 0 // @license   http://opensource.org/licenses/MIT The MIT License (MIT)`,
}

var expLineComments = [][]string{
	[]string{`function () { `, `// Copyright 2016 David Lavieri. All rights reserved.`},
	[]string{`some code `, `// Use of this source code is governed by a MIT License`},
	[]string{``, `// License that can be found in the LICENSE file.`},
	[]string{`<script type="text/javascript"> `, `// Date: 0/0/0`},
	[]string{`console.log({include`, `// file=$myCustomFile}) // Time: 0:0 PM`},
	[]string{`</script> `, `// @author    David Lavieri (falmar) <daviddlavier@gmail.com>`},
	[]string{``, `// @copyright 2016 David Lavieri`},
	[]string{`let some = 0 `, `// @license   http://opensource.org/licenses/MIT The MIT License (MIT)`},
}

var expMatchLineComments = []bool{
	true,
	true,
	false,
	true,
	true,
	true,
	false,
	true,
}

var nonLineComments = []string{
	`<body>`,
	`{$some_variable}`,
	`Outside the script tag may be pure html or may not`,
	`<script type="text/javascript">`,
	`let myVar = {json_decode($jsonVariable)}`,
	`let myOtherVar = '{$wuuuu}'`,
	`console.log({include file=$myCustomFile})`,
	`let some = 0`,
	`{rdelim}`,
	`hello: "world"`,
	`world: "hello"`,
	`{rdelim})`,
	`hello: "world",`,
	`one: 1,`,
	`two: [2, 2]`,
	`{rdelim}`,
	`{rdelim}]`,
	`{rdelim}, maybe: ""{rdelim}, did: "not"{rdelim}, work: "entirely"{rdelim}`,
	`</script>`,
	`</body>`,
}

func TestIsLineCommentMatch(t *testing.T) {
	for _, c := range lineComments {
		if !isLineComment(c) {
			t.Fatalf("Should be/have a line comment %s", c)
		}
	}
}

func TestIsLineCommentNoMatch(t *testing.T) {
	for _, c := range nonLineComments {
		if isLineComment(c) {
			t.Fatalf("Should not be/have a line comment %s", c)
		}
	}
}

func TestLineCommentMatch(t *testing.T) {
	for i, c := range lineComments {
		left, match := parseLineComment(c)

		if match != expMatchLineComments[i] {
			t.Fatalf("Expected match %v; got: %v", expMatchLineComments[i], match)
		}

		if left[0] != expLineComments[i][0] {
			t.Fatalf("Expected left most parse: %s; got: %s", expLineComments[i][0], left[0])
		}

		if left[1] != expLineComments[i][1] {
			t.Fatalf("Expected right most parse: %s; got: %s", expLineComments[i][1], left[1])
		}
	}
}

func TestLineCommentNoMatch(t *testing.T) {
	for _, c := range nonLineComments {
		left, match := parseLineComment(c)

		if match {
			t.Fatal("Expected no match")
		}

		if left[1] != c {
			t.Fatalf("Expected left most be intact: %s; got: %s", c, left[1])
		}
	}
}
