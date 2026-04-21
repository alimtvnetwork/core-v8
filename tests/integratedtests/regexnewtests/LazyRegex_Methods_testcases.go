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

package regexnewtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var lazyRegexCompileTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.Compile returns no error -- valid pattern",
		When:  "given a valid pattern to Compile",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.regexNotNil: true,
			params.hasError:    false,
			params.isCompiled:  true,
		},
	},
	{
		Title: "LazyRegex.Compile returns error -- invalid pattern",
		When:  "given an invalid pattern to Compile",
		Input: args.Map{
			params.pattern: "[bad",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
			params.isCompiled:  true,
		},
	},
	{
		Title: "LazyRegex.Compile returns error -- empty pattern on undefined lazy",
		When:  "given empty pattern to Compile",
		Input: args.Map{
			params.pattern: "",
		},
		Expected: args.Map{
			params.regexNotNil: false,
			params.hasError:    true,
			params.isCompiled:  false,
		},
	},
}

var lazyRegexHasErrorTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.HasError returns false -- valid pattern",
		When:  "given valid pattern for HasError",
		Input: args.Map{
			params.pattern: "hello",
		},
		Expected: args.Map{
			params.hasError:  false,
			params.isInvalid: false,
		},
	},
	{
		Title: "LazyRegex.HasError returns true -- invalid pattern",
		When:  "given invalid pattern for HasError",
		Input: args.Map{
			params.pattern: "[broken",
		},
		Expected: args.Map{
			params.hasError:  true,
			params.isInvalid: true,
		},
	},
}

var lazyRegexMatchBytesTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.IsMatchBytes returns true -- matching bytes",
		When:  "given matching byte input",
		Input: args.Map{
			params.pattern:      "\\d+",
			params.compareInput: "abc123",
		},
		Expected: args.Map{
			params.isMatchBytes:       true,
			params.isFailedMatchBytes: false,
		},
	},
	{
		Title: "LazyRegex.IsMatchBytes returns false -- non-matching bytes",
		When:  "given non-matching byte input",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isMatchBytes:       false,
			params.isFailedMatchBytes: true,
		},
	},
}

var lazyRegexMatchErrorTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.MatchError returns nil -- matching input",
		When:  "given matching input to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "^hello$",
			params.compareInput: "hello",
		},
		Expected: args.Map{
			params.isNoError: true,
		},
	},
	{
		Title: "LazyRegex.MatchError returns error -- non-matching input",
		When:  "given non-matching input to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "^\\d+$",
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
	{
		Title: "LazyRegex.MatchError returns error -- invalid regex",
		When:  "given invalid regex to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.isNoError: false,
		},
	},
}

var lazyRegexStringTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.String returns pattern -- valid LazyRegex",
		When:  "given valid pattern for String",
		Input: args.Map{
			params.pattern: "\\d+",
		},
		Expected: args.Map{
			params.stringResult: "\\d+",
		},
	},
	{
		Title: "LazyRegex.Pattern returns original pattern -- email pattern",
		When:  "given email pattern for Pattern",
		Input: args.Map{
			params.pattern: `[a-z]+@[a-z]+\.[a-z]+`,
		},
		Expected: args.Map{
			params.stringResult: `[a-z]+@[a-z]+\.[a-z]+`,
		},
	},
}
