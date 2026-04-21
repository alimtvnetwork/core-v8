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

package corestrtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// LeftMiddleRightFromSplit — edge cases
// ==========================================================================

var leftMiddleRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns valid triple -- 'a:b:c' three-part split",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTwoPartsTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns invalid -- 'a:b' two parts only",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitSinglePartTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns invalid -- 'hello' no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitFourPlusTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns middle=second, right=last -- 'a:b:c:d' four+ parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "d",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns all empty invalid -- empty string input",
	ExpectedInput: args.Map{
		"left":    "",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitEdgesTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplit returns valid empty parts -- separator at edges",
	ExpectedInput: args.Map{
		"left":    "",
		"middle":  "",
		"right":   "",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftMiddleRightFromSplitTrimmedAllTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplitTrimmed returns trimmed parts -- ' a : b : c ' with whitespace",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "LeftMiddleRightFromSplitTrimmed returns invalid -- ' a : b ' two parts with whitespace",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitN — remainder handling
// ==========================================================================

var leftMiddleRightFromSplitNRemainderTestCase = coretestcases.CaseV1{
	Title: "SplitN returns remainder in right -- 'a:b:c:d:e' keeps 'c:d:e'",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c:d:e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNExact3TestCase = coretestcases.CaseV1{
	Title: "SplitN returns exact parts -- 'a:b:c' exactly 3 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTwoOnlyTestCase = coretestcases.CaseV1{
	Title: "SplitN returns invalid -- 'a:b' only 2 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}

var leftMiddleRightFromSplitNMissingSepTestCase = coretestcases.CaseV1{
	Title: "SplitN returns invalid -- 'nosep' missing separator",
	ExpectedInput: args.Map{
		"left":    "nosep",
		"middle":  "",
		"right":   "",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftMiddleRightFromSplitNTrimmed — remainder + trimming
// ==========================================================================

var leftMiddleRightFromSplitNTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title: "SplitNTrimmed returns trimmed remainder -- ' a : b : c : d : e ' with whitespace",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "b",
		"right":   "c : d : e",
		"isValid": "true",
	},
}

var leftMiddleRightFromSplitNTrimmedTwoTestCase = coretestcases.CaseV1{
	Title: "SplitNTrimmed returns invalid -- ' a : b ' only 2 parts",
	ExpectedInput: args.Map{
		"left":    "a",
		"middle":  "",
		"right":   "b",
		"isValid": "false",
	},
}
