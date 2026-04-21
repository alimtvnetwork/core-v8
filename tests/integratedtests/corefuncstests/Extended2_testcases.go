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

package corefuncstests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// GetFuncName
// ==========================================================================

var ext2GetFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetFuncName returns short name of named function",
		ArrangeInput:  args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// GetFuncFullName
// ==========================================================================

var ext2GetFuncFullNameTestCases = []coretestcases.CaseV1{
	{
		Title:         "GetFuncFullName returns full qualified name",
		ArrangeInput:  args.Map{"when": "given a named function"},
		ExpectedInput: args.Map{"notEmpty": true},
	},
}

// ==========================================================================
// NewInOutErrWrapper
// ==========================================================================

var ext2NewInOutErrWrapperTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewInOutErrWrapper.Exec returns output and nil error -- success",
		ArrangeInput: args.Map{"when": "given valid input"},
		ExpectedInput: args.Map{
			"result": 5,
			"isNil":  true,
			"name":   "parse",
		},
	},
}

var ext2NewInOutErrWrapperFailTestCases = []coretestcases.CaseV1{
	{
		Title:        "NewInOutErrWrapper.Exec returns error -- failure",
		ArrangeInput: args.Map{"when": "given empty input"},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// ==========================================================================
// InOutErrFuncWrapperOf.ToLegacy
// ==========================================================================

var ext2InOutErrToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title:        "InOutErrFuncWrapperOf.ToLegacy converts -- success",
		ArrangeInput: args.Map{"when": "given generic wrapper converted to legacy"},
		ExpectedInput: args.Map{
			"name":   "legacy-parse",
			"result": "5",
			"isNil":  true,
		},
	},
}

// ==========================================================================
// SerializeOutputFuncWrapperOf.AsActionReturnsErrorFunc
// ==========================================================================

var ext2SerializeAsErrFuncTestCases = []coretestcases.CaseV1{
	{
		Title:        "SerializeOutputFuncWrapperOf.AsActionReturnsErrorFunc success -- nil",
		ArrangeInput: args.Map{"when": "given serialize wrapper as err func"},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
}
