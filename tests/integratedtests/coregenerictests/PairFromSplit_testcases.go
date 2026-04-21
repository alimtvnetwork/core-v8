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

package coregenerictests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================
// PairFromSplit
// ==========================================

var pairFromSplitTestCases = []coretestcases.CaseV1{
	{
		Title: "Standard key=value split",
		ArrangeInput: args.Map{
			"input": "key=value",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "key",
			"right":   "value",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "No separator found produces invalid pair",
		ArrangeInput: args.Map{
			"input": "noseparator",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "noseparator",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Empty input produces invalid pair",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Multiple separators takes first two parts only",
		ArrangeInput: args.Map{
			"input": "a=b=c=d",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"right":   "b=c=d",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Separator at start produces empty left",
		ArrangeInput: args.Map{
			"input": "=value",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "",
			"right":   "value",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key=",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "key",
			"right":   "",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "hello::world",
			"sep":   "::",
		},
		ExpectedInput: args.Map{
			"left":    "hello",
			"right":   "world",
			"isValid": true,
			"message": "",
		},
	},
}

// ==========================================
// PairFromSplitTrimmed
// ==========================================

var pairFromSplitTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Trims whitespace from both parts",
		ArrangeInput: args.Map{
			"input": "  key  =  value  ",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "key",
			"right":   "value",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "No separator trims single part",
		ArrangeInput: args.Map{
			"input": "  onlypart  ",
			"sep":   "=",
		},
		ExpectedInput: args.Map{
			"left":    "onlypart",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
}

// ==========================================
// PairFromSplitFull
// ==========================================

var pairFromSplitFullTestCases = []coretestcases.CaseV1{
	{
		Title: "Splits at first separator only, right gets remainder",
		ArrangeInput: args.Map{
			"input": "a:b:c:d",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"right":   "b:c:d",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "No separator produces invalid pair",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "nosep",
			"right":   "",
			"isValid": false,
			"message": "separator not found",
		},
	},
	{
		Title: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key:",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "key",
			"right":   "",
			"isValid": true,
			"message": "",
		},
	},
}

// ==========================================
// PairFromSplitFullTrimmed
// ==========================================

var pairFromSplitFullTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Splits at first separator with trimming",
		ArrangeInput: args.Map{
			"input": "  a  :  b : c : d  ",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "a",
			"right":   "b : c : d",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "No separator trims and marks invalid",
		ArrangeInput: args.Map{
			"input": "  nosep  ",
			"sep":   ":",
		},
		ExpectedInput: args.Map{
			"left":    "nosep",
			"right":   "",
			"isValid": false,
			"message": "separator not found",
		},
	},
}

// ==========================================
// PairFromSlice
// ==========================================

var pairFromSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "Two-element slice produces valid pair",
		ArrangeInput: args.Map{
			"parts": []string{"left", "right"},
		},
		ExpectedInput: args.Map{
			"left":    "left",
			"right":   "right",
			"isValid": true,
			"message": "",
		},
	},
	{
		Title: "Single-element slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		ExpectedInput: args.Map{
			"left":    "only",
			"right":   "",
			"isValid": false,
			"message": "only one part found",
		},
	},
	{
		Title: "Empty slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		ExpectedInput: args.Map{
			"left":    "",
			"right":   "",
			"isValid": false,
			"message": "empty input",
		},
	},
	{
		Title: "Three-element slice uses first and last",
		ArrangeInput: args.Map{
			"parts": []string{"first", "middle", "last"},
		},
		ExpectedInput: args.Map{
			"left":    "first",
			"right":   "last",
			"isValid": true,
			"message": "",
		},
	},
}
