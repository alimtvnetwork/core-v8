// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package stringcompareastests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var globMatchTestCases = []coretestcases.CaseV1{
	// === Positive: basic wildcard ===
	{
		Title: "Glob - star matches any filename",
		ArrangeInput: args.Map{
			"when":         "given glob with star wildcard",
			"pattern":      "*.go",
			"content":      "main.go",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches path segment",
		ArrangeInput: args.Map{
			"when":         "given glob matching dynamic directory",
			"pattern":      "build-*/result.json",
			"content":      "build-20260303/result.json",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - question mark matches single char",
		ArrangeInput: args.Map{
			"when":         "given glob with question mark",
			"pattern":      "file?.txt",
			"content":      "fileA.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - character class matches range",
		ArrangeInput: args.Map{
			"when":         "given glob with character class",
			"pattern":      "log[0-9].txt",
			"content":      "log5.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - exact match without wildcards",
		ArrangeInput: args.Map{
			"when":         "given glob with no wildcards",
			"pattern":      "exact.txt",
			"content":      "exact.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},

	// === Case insensitive ===
	{
		Title: "Glob - ignore case matches different casing",
		ArrangeInput: args.Map{
			"when":         "given glob with ignore case",
			"pattern":      "*.GO",
			"content":      "main.go",
			"isIgnoreCase": true,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - case sensitive rejects different casing",
		ArrangeInput: args.Map{
			"when":         "given glob without ignore case",
			"pattern":      "*.GO",
			"content":      "main.go",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Negative: no match ===
	{
		Title: "Glob - no match returns false",
		ArrangeInput: args.Map{
			"when":         "given non-matching content",
			"pattern":      "*.go",
			"content":      "main.rs",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},
	{
		Title: "Glob - question mark rejects multi-char",
		ArrangeInput: args.Map{
			"when":         "given content longer than question mark allows",
			"pattern":      "file?.txt",
			"content":      "fileAB.txt",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Edge: invalid pattern ===
	{
		Title: "Glob - invalid pattern returns false gracefully",
		ArrangeInput: args.Map{
			"when":         "given malformed glob pattern",
			"pattern":      "[invalid",
			"content":      "anything",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "false",
			"isInverse": "true",
		},
	},

	// === Edge: empty ===
	{
		Title: "Glob - empty pattern matches empty content",
		ArrangeInput: args.Map{
			"when":         "given empty pattern and content",
			"pattern":      "",
			"content":      "",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches empty content",
		ArrangeInput: args.Map{
			"when":         "given star pattern with empty content",
			"pattern":      "*",
			"content":      "",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
	{
		Title: "Glob - star matches any content",
		ArrangeInput: args.Map{
			"when":         "given star pattern with any content",
			"pattern":      "*",
			"content":      "anything-at-all",
			"isIgnoreCase": false,
		},
		ExpectedInput: args.Map{
			"isMatch":   "true",
			"isInverse": "false",
		},
	},
}
