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

package corecomparatortests

// Extended test cases migrated from cmd/main/enumTesting.go and
// cmd/main/compareEnumTesting02.go.

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// -------------------------------------------------------------------------
// enumTesting — JSON marshal/unmarshal roundtrip for Compare enum
// -------------------------------------------------------------------------

var compareJsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Marshal returns quoted name and Unmarshal restores identity -- Equal value 0, unmarshal '3'",
		ArrangeInput: args.Map{
			"value":          0,
			"unmarshalInput": "3",
		},
		ExpectedInput: args.Map{
			"marshaledJson":    "\"Equal\"",
			"unmarshaledName":  "LeftLess",
			"unmarshaledValue": "3",
		},
	},
}

// -------------------------------------------------------------------------
// compareEnumTesting02 — OnlySupportedErr
// -------------------------------------------------------------------------

var onlySupportedErrTestCases = []coretestcases.CaseV1{
	{
		Title: "OnlySupportedErr returns error -- Equal (0) not in supported [6,3]",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "dining doesn't support more",
			"supported": []int{6, 3},
		},
		ExpectedInput: "true", // hasError
	},
	{
		Title: "OnlySupportedErr returns nil -- Equal (0) in supported [0,3]",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "some context message",
			"supported": []int{0, 3},
		},
		ExpectedInput: "false", // hasError
	},
	{
		Title: "OnlySupportedErr returns error -- Inconclusive (6) not in supported [0]",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "only equal allowed",
			"supported": []int{0},
		},
		ExpectedInput: "true", // hasError
	},
	{
		Title: "OnlySupportedErr returns nil -- Inconclusive (6) in supported [6,0,3]",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "inconclusive is fine",
			"supported": []int{6, 0, 3},
		},
		ExpectedInput: "false", // hasError
	},
}
