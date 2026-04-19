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

package corevalidatortests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var simpleSliceValidatorSetActualTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator SetActual returns same instance",
	ExpectedInput: args.Map{
		"sameInstance": true,
	},
}

var simpleSliceValidatorSliceValidatorTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator SliceValidator returns non-nil",
	ExpectedInput: args.Map{
		"isNotNil": true,
	},
}

var simpleSliceValidatorVerifyAllTestCases = []coretestcases.CaseV1{
	{
		Title: "SimpleSliceValidator VerifyAll matching returns nil",
		ArrangeInput: args.Map{
			"expected": []string{"a", "b"},
			"actual":   []string{"a", "b"},
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "SimpleSliceValidator VerifyAll mismatch returns error",
		ArrangeInput: args.Map{
			"expected": []string{"a", "b"},
			"actual":   []string{"x", "y"},
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

var simpleSliceValidatorVerifyFirstTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator VerifyFirst matching returns nil",
	ArrangeInput: args.Map{
		"expected": []string{"a"},
		"actual":   []string{"a"},
	},
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var simpleSliceValidatorVerifyUptoTestCase = coretestcases.CaseV1{
	Title: "SimpleSliceValidator VerifyUpto matching returns nil",
	ArrangeInput: args.Map{
		"expected": []string{"a", "b", "c"},
		"actual":   []string{"a", "b", "c"},
		"length":   2,
	},
	ExpectedInput: args.Map{
		"hasError": false,
	},
}
