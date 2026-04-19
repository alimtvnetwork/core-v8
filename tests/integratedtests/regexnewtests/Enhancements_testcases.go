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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// 1. Create caching — repeated Create returns same pointer
// =============================================================================

var createCachingTestCases = []coretestcases.MapGherkins{
	{
		Title: "Create returns same pointer on repeated calls -- caching behavior",
		When:  "given the same valid pattern called twice via CreateLock",
		Input: args.Map{
			params.pattern: `\w+caching_test_unique`,
		},
		Expected: args.Map{
			params.samePointer: true,
			params.hasError:    false,
		},
	},
	{
		Title: "Create returns same pointer for previously cached pattern",
		When:  "given a digit pattern called twice",
		Input: args.Map{
			params.pattern: `\d+caching_test_unique2`,
		},
		Expected: args.Map{
			params.samePointer: true,
			params.hasError:    false,
		},
	},
}

// =============================================================================
// 2. ManyUsingLock — empty input returns empty map
// =============================================================================

var manyUsingLockEmptyTestCases = []coretestcases.MapGherkins{
	{
		Title: "ManyUsingLock returns empty map -- no patterns given",
		When:  "given zero patterns to ManyUsingLock",
		Expected: args.Map{
			params.mapLength: 0,
		},
	},
}

// =============================================================================
// 3. ManyUsingLock — valid multi-pattern
// =============================================================================

var manyUsingLockValidTestCases = []coretestcases.MapGherkins{
	{
		Title: "ManyUsingLock returns map with all patterns -- 3 valid patterns",
		When:  "given 3 valid patterns to ManyUsingLock",
		Expected: args.Map{
			params.mapLength: 3,
		},
	},
}

// =============================================================================
// 4. TwoLock — both valid patterns
// =============================================================================

var twoLockValidTestCases = []coretestcases.MapGherkins{
	{
		Title: "TwoLock returns two applicable LazyRegex -- both valid patterns",
		When:  "given two valid patterns to TwoLock",
		Input: args.Map{
			params.pattern:      `\d+`,
			params.compareInput: `[a-z]+`,
		},
		Expected: args.Map{
			params.isDefined:    true,
			params.isApplicable: true,
		},
	},
}

// =============================================================================
// 5. TwoLock — one invalid pattern
// =============================================================================

var twoLockInvalidTestCases = []coretestcases.MapGherkins{
	{
		Title: "TwoLock returns one applicable, one not -- first valid second invalid",
		When:  "given one valid and one invalid pattern to TwoLock",
		Input: args.Map{
			params.pattern:      `\d+`,
			params.compareInput: "[bad",
		},
		Expected: args.Map{
			// first should be applicable, second should not
			params.isDefined: true,
		},
	},
	{
		Title: "TwoLock returns both not applicable -- both invalid patterns",
		When:  "given two invalid patterns to TwoLock",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "(unclosed",
		},
		Expected: args.Map{
			params.isDefined: true,
		},
	},
}

// =============================================================================
// 6. MatchError error message verification -- compile error branch
// =============================================================================

var matchErrorMessageCompileTestCases = []coretestcases.MapGherkins{
	{
		Title: "MatchError error contains 'compile failed' -- invalid regex pattern",
		When:  "given invalid regex pattern to MatchError",
		Input: args.Map{
			params.pattern:      "[bad",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.hasError:      true,
			params.errorContains: "compile failed",
		},
	},
}

// =============================================================================
// 7. MatchError error message verification -- mismatch branch (valid regex, no match)
// =============================================================================

var matchErrorMessageMismatchTestCases = []coretestcases.MapGherkins{
	{
		Title: "MatchError error contains 'doesn't match' -- valid regex, non-matching input",
		When:  "given valid regex that doesn't match comparing input",
		Input: args.Map{
			params.pattern:      `^\d+$`,
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.hasError:      true,
			params.errorContains: "doesn't match",
		},
	},
}

// =============================================================================
// 8. LazyRegex.MatchError error message verification -- compile error branch
// =============================================================================

var lazyMatchErrorMessageTestCases = []coretestcases.MapGherkins{
	{
		Title: "LazyRegex.MatchError error contains 'compile failed' -- invalid pattern",
		When:  "given invalid pattern to LazyRegex.MatchError",
		Input: args.Map{
			params.pattern:      "[broken",
			params.compareInput: "test",
		},
		Expected: args.Map{
			params.hasError:      true,
			params.errorContains: "compile failed",
		},
	},
	{
		Title: "LazyRegex.MatchError error contains 'doesn't match' -- valid pattern no match",
		When:  "given valid pattern that doesn't match",
		Input: args.Map{
			params.pattern:      `^\d+$`,
			params.compareInput: "abc",
		},
		Expected: args.Map{
			params.hasError:      true,
			params.errorContains: "doesn't match",
		},
	},
}
