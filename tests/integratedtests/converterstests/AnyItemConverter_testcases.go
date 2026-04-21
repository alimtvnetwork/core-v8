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

package converterstests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// ToNonNullItems
// =============================================================================

var toNonNullItemsTestCases = []coretestcases.CaseV1{
	{
		Title: "ToNonNullItems returns count 0 -- nil input skipOnNil true",
		ArrangeInput: args.Map{
			"when":        "given nil input with isSkipOnNil true",
			"isSkipOnNil": true,
			"input":       nil,
		},
		ExpectedInput: args.Map{"count": 0},
	},
	{
		Title: "ToNonNullItems returns count 2 -- valid string slice",
		ArrangeInput: args.Map{
			"when":        "given valid string slice",
			"isSkipOnNil": false,
			"input":       []any{"hello", "world"},
		},
		ExpectedInput: args.Map{
			"count": 2,
			"item0": "hello",
			"item1": "world",
		},
	},
	{
		Title: "ToNonNullItems returns count 0 -- nil input skipOnNil false",
		ArrangeInput: args.Map{
			"when":        "given nil input with isSkipOnNil false - should still return empty for nil reflect",
			"isSkipOnNil": true,
			"input":       nil,
		},
		ExpectedInput: args.Map{"count": 0},
	},
}
