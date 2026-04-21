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
// LeftRightFromSplit — edge cases
// ==========================================================================

var leftRightFromSplitNormalTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns valid split -- 'key=value' normal input",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "no-separator-here",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitEmptyTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty invalid -- empty input string",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitSepAtStartTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty left -- separator at start",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitSepAtEndTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns empty right -- separator at end",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "",
		"isValid": "true",
	},
}

var leftRightFromSplitMultipleSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplit returns first-left and last-right -- multiple separators",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b=c",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftRightFromSplitTrimmed — trimming edge cases
// ==========================================================================

var leftRightFromSplitTrimmedTrimsTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns trimmed parts -- whitespace around both",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitTrimmedNoSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns trimmed invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}

var leftRightFromSplitTrimmedWhitespaceTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitTrimmed returns empty parts -- whitespace-only values",
	ExpectedInput: args.Map{
		"left":    "",
		"right":   "",
		"isValid": "true",
	},
}

// ==========================================================================
// LeftRightFromSplitFull — remainder handling
// ==========================================================================

var leftRightFromSplitFullRemainderTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns remainder in right -- 'a:b:c:d' multi-separator",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b:c:d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullSingleSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns same as normal -- single separator",
	ExpectedInput: args.Map{
		"left":    "key",
		"right":   "value",
		"isValid": "true",
	},
}

var leftRightFromSplitFullMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFull returns invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "nosep",
		"right":   "",
		"isValid": "false",
	},
}

// ==========================================================================
// LeftRightFromSplitFullTrimmed — remainder + trimming
// ==========================================================================

var leftRightFromSplitFullTrimmedRemainderTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFullTrimmed returns trimmed remainder -- ' a : b : c : d ' with spaces",
	ExpectedInput: args.Map{
		"left":    "a",
		"right":   "b : c : d",
		"isValid": "true",
	},
}

var leftRightFromSplitFullTrimmedMissingSepTestCase = coretestcases.CaseV1{
	Title: "LeftRightFromSplitFullTrimmed returns trimmed invalid -- no separator found",
	ExpectedInput: args.Map{
		"left":    "hello",
		"right":   "",
		"isValid": "false",
	},
}
