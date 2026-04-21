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

package enumimpltests

import (
	"github.com/alimtvnetwork/core-v8/coreimpl/enumimpl"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var enumByteTestCases = []coretestcases.CaseV1{
	{
		Title: "EnumByte returns min 0 and max 10 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   0,
				"A":         -2,
				"B":         8,
				"C":         5,
				"Something": 10,
			},
		},
		ExpectedInput: args.Map{
			"min": 0,
			"max": 10,
		},
	},
}

var enumInt8TestCases = []coretestcases.CaseV1{
	{
		Title: "EnumInt8 returns min -2 and max 12 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -2,
				"A":         -2,
				"B":         8,
				"C":         5,
				"Something": 12,
			},
		},
		ExpectedInput: args.Map{
			"min": -2,
			"max": 12,
		},
	},
}

var enumInt16TestCases = []coretestcases.CaseV1{
	{
		Title: "EnumInt16 returns min -3 and max 14 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -3,
				"A":         -2,
				"B":         -3,
				"C":         5,
				"Something": 14,
			},
		},
		ExpectedInput: args.Map{
			"min": -3,
			"max": 14,
		},
	},
}

var enumInt32TestCases = []coretestcases.CaseV1{
	{
		Title: "EnumInt32 returns min -4 and max 15 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":   -4,
				"A":         -2,
				"B":         -3,
				"C":         5,
				"Something": 15,
			},
		},
		ExpectedInput: args.Map{
			"min": -4,
			"max": 15,
		},
	},
}

var enumUInt16TestCases = []coretestcases.CaseV1{
	{
		Title: "EnumUInt16 returns min 0 and max 20 -- DynamicMap input",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":    0,
				"Something2": 15,
				"B":          15,
				"Something":  20,
			},
		},
		ExpectedInput: args.Map{
			"min": 0,
			"max": 20,
		},
	},
}

var enumStringTestCases = []coretestcases.CaseV1{
	{
		Title: "EnumString returns min empty and max 'Something2' -- DynamicMap input lexicographic",
		ArrangeInput: args.Map{
			"enum-map": enumimpl.DynamicMap{
				"Invalid":    0,
				"Something2": 15,
				"B":          15,
				"Something":  20,
			},
		},
		ExpectedInput: args.Map{
			"min": "B",
			"max": "Something2",
		},
	},
}
