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

package coreinstructiontests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IsMatch test cases
// =============================================================================

var stringSearchIsMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsMatch returns true -- equal match, search 'hello' in content 'hello'",
		ArrangeInput: args.Map{
			"when":    "given matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "hello",
		},
		ExpectedInput: args.Map{
			"isMatch":       true,
			"isMatchFailed": false,
		},
	},
	{
		Title: "IsMatch returns false -- equal no match, search 'hello' in content 'world'",
		ArrangeInput: args.Map{
			"when":    "given non-matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "world",
		},
		ExpectedInput: args.Map{
			"isMatch":       false,
			"isMatchFailed": true,
		},
	},
	{
		Title: "IsMatch returns true -- contains match, search 'world' in 'hello world'",
		ArrangeInput: args.Map{
			"when":    "given content containing search",
			"method":  "contains",
			"search":  "world",
			"content": "hello world",
		},
		ExpectedInput: args.Map{
			"isMatch":       true,
			"isMatchFailed": false,
		},
	},
	{
		Title: "IsMatch returns false -- contains no match, search 'xyz' in 'hello world'",
		ArrangeInput: args.Map{
			"when":    "given content not containing search",
			"method":  "contains",
			"search":  "xyz",
			"content": "hello world",
		},
		ExpectedInput: args.Map{
			"isMatch":       false,
			"isMatchFailed": true,
		},
	},
}

// =============================================================================
// IsAllMatch test cases
// =============================================================================

var stringSearchIsAllMatchTestCases = []coretestcases.CaseV1{
	{
		Title: "IsAllMatch returns true -- all contents contain 'o' in ['hello','world','foo']",
		ArrangeInput: args.Map{
			"when":     "given all contents containing search",
			"method":   "contains",
			"search":   "o",
			"contents": []string{"hello", "world", "foo"},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       true,
			"isAnyMatchFailed": false,
		},
	},
	{
		Title: "IsAllMatch returns false -- 'world' missing 'z' in ['hello','buzz','world']",
		ArrangeInput: args.Map{
			"when":     "given one content not containing search",
			"method":   "contains",
			"search":   "z",
			"contents": []string{"hello", "buzz", "world"},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       false,
			"isAnyMatchFailed": true,
		},
	},
	{
		Title: "IsAllMatch returns true -- empty contents slice",
		ArrangeInput: args.Map{
			"when":     "given empty contents",
			"method":   "equal",
			"search":   "hello",
			"contents": []string{},
		},
		ExpectedInput: args.Map{
			"isAllMatch":       true,
			"isAnyMatchFailed": false,
		},
	},
}

// =============================================================================
// IsEmpty / IsExist / Has test cases
// =============================================================================

var stringSearchStateTestCases = []coretestcases.CaseV1{
	{
		Title: "StringSearch returns isEmpty false, isExist true -- non-nil instance",
		ArrangeInput: args.Map{
			"when":   "given non-nil StringSearch",
			"method": "equal",
			"search": "test",
			"isNil":  false,
		},
		ExpectedInput: args.Map{
			"isEmpty": false,
			"isExist": true,
			"has":     true,
		},
	},
	{
		Title: "StringSearch returns isEmpty true, isExist false -- nil instance",
		ArrangeInput: args.Map{
			"when":  "given nil StringSearch",
			"isNil": true,
		},
		ExpectedInput: args.Map{
			"isEmpty": true,
			"isExist": false,
			"has":     false,
		},
	},
}

// =============================================================================
// VerifyError test cases
// =============================================================================

var stringSearchVerifyErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "VerifyError returns nil -- equal match, search 'hello' in 'hello'",
		ArrangeInput: args.Map{
			"when":    "given matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "hello",
			"isNil":   false,
		},
		ExpectedInput: "false",
	},
	{
		Title: "VerifyError returns error -- equal no match, search 'hello' in 'world'",
		ArrangeInput: args.Map{
			"when":    "given non-matching equal string",
			"method":  "equal",
			"search":  "hello",
			"content": "world",
			"isNil":   false,
		},
		ExpectedInput: "true",
	},
	{
		Title: "VerifyError returns nil -- nil receiver, content 'anything'",
		ArrangeInput: args.Map{
			"when":    "given nil StringSearch",
			"content": "anything",
			"isNil":   true,
		},
		ExpectedInput: "false",
	},
}
