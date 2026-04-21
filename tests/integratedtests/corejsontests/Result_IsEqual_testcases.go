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

package corejsontests

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IsEqual test cases
// =============================================================================

var resultIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqual returns true -- same content",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "value"}),
			"b": corejson.New(map[string]string{"key": "value"}),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqual returns false -- different content",
		ArrangeInput: args.Map{
			"a": corejson.New(map[string]string{"key": "a"}),
			"b": corejson.New(map[string]string{"key": "b"}),
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// IsEqualPtr test cases
// =============================================================================

var resultIsEqualPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEqualPtr returns true -- both nil",
		ArrangeInput: args.Map{
			"aPtr": (*corejson.Result)(nil),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEqualPtr returns false -- one nil",
		ArrangeInput: args.Map{
			"aPtr": corejson.NewPtr(map[string]string{"k": "v"}),
			"bPtr": (*corejson.Result)(nil),
		},
		ExpectedInput: "false",
	},
}
