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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================
// DynamicMap
// ==========================================

var extDynMapBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "DynamicMap Length and IsEmpty",
		ArrangeInput: args.Map{
			"items": map[string]any{"A": 1, "B": 2},
		},
		ExpectedInput: args.Map{
			"length":  "2",
			"isEmpty": "false",
		},
	},
	{
		Title: "DynamicMap empty",
		ArrangeInput: args.Map{
			"items": map[string]any{},
		},
		ExpectedInput: args.Map{
			"length":  "0",
			"isEmpty": "true",
		},
	},
}

// ==========================================
// DiffLeftRight
// ==========================================

var extDiffLeftRightTestCases = []coretestcases.CaseV1{
	{
		Title: "DiffLeftRight same values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "hello",
		},
		ExpectedInput: args.Map{
			"isSame":     "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "DiffLeftRight different values",
		ArrangeInput: args.Map{
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: args.Map{
			"isSame":     "false",
			"isNotEqual": "true",
		},
	},
}

// ==========================================
// KeyAnyVal
// ==========================================

var extKeyAnyValTestCases = []coretestcases.CaseV1{
	{
		Title: "KeyAnyVal with int value",
		ArrangeInput: args.Map{
			"key":   "Invalid",
			"value": 0,
		},
		ExpectedInput: args.Map{
			"key":      "Invalid",
			"valInt":   "0",
			"isString": "false",
		},
	},
	{
		Title: "KeyAnyVal with string value",
		ArrangeInput: args.Map{
			"key":   "Name",
			"value": "hello",
		},
		ExpectedInput: args.Map{
			"key":      "Name",
			"isString": "true",
		},
	},
}

// ==========================================
// KeyValInteger
// ==========================================

var extKeyValIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "KeyValInteger normal",
		ArrangeInput: args.Map{
			"key":   "Variant",
			"value": 5,
		},
		ExpectedInput: args.Map{
			"key":      "Variant",
			"isString": "false",
		},
	},
}
