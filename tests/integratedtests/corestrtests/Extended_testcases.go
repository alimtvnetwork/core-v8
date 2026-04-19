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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// Collection
// ==========================================

var extCollectionAddTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection Add single item",
		ArrangeInput: args.Map{
			"items": []string{"hello"},
		},
		ExpectedInput: "1",
	},
	{
		Title: "Collection Adds multiple items",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: "3",
	},
}

var extCollectionJoinTestCases = []coretestcases.CaseV1{
	{
		Title: "Collection Join with comma",
		ArrangeInput: args.Map{
			"items":  []string{"a", "b", "c"},
			"joiner": ",",
		},
		ExpectedInput: "a,b,c",
	},
}

// ==========================================
// SimpleSlice
// ==========================================

var extSimpleSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSlice Add and Length",
		ArrangeInput: args.Map{
			"items": []string{"a", "b"},
		},
		ExpectedInput: "2",
	},
}

// ==========================================
// LeftRight
// ==========================================

var extLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "LeftRight valid creation",
		ArrangeInput: args.Map{
			"left":  "key",
			"right": "value",
		},
		ExpectedInput: args.Map{
			"left": "key",
			"right": "value",
			"isValid": "true",
		},
	},
	{
		Title: "LeftRight from single-element slice",
		ArrangeInput: args.Map{
			"slice": []string{"only"},
		},
		ExpectedInput: args.Map{
			"left": "only",
			"right": "",
			"isValid": "false",
		},
	},
	{
		Title: "LeftRight from two-element slice",
		ArrangeInput: args.Map{
			"slice": []string{"first", "second"},
		},
		ExpectedInput: args.Map{
			"left": "first",
			"right": "second",
			"isValid": "true",
		},
	},
}

// ==========================================
// LeftMiddleRight
// ==========================================

var extLeftMiddleRightTestCases = []coretestcases.CaseV1{
	{
		Title: "LeftMiddleRight valid creation",
		ArrangeInput: args.Map{
			"left":   "L",
			"middle": "M",
			"right":  "R",
		},
		ExpectedInput: args.Map{
			"left": "L", "middle": "M", "right": "R",
			"isLeftEmpty": "false", "isMiddleEmpty": "false", "isRightEmpty": "false",
		},
	},
}

// ==========================================
// Hashset
// ==========================================

var extHashsetTestCases = []coretestcases.CaseV1{
	{
		Title: "Hashset Add and Has",
		ArrangeInput: args.Map{
			"items": []string{"a", "b", "c"},
		},
		ExpectedInput: args.Map{
			"length":  "3",
			"hasA":    "true",
			"hasB":    "true",
			"hasMiss": "false",
		},
	},
}

// ==========================================
// ValidValue
// ==========================================

var extValidValueTestCases = []coretestcases.CaseV1{
	{
		Title: "ValidValue valid",
		ArrangeInput: args.Map{
			"value":   "hello",
			"isValid": true,
		},
		ExpectedInput: args.Map{
			"value": "hello",
			"isValid": "true",
		},
	},
	{
		Title: "ValidValue invalid",
		ArrangeInput: args.Map{
			"value":   "",
			"isValid": false,
		},
		ExpectedInput: args.Map{
			"value": "",
			"isValid": "false",
		},
	},
}
