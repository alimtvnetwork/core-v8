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

package defaulterrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var defaultErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "DefaultErr returns non-nil error with message -- Marshalling",
		ArrangeInput: args.Map{
			"when":  "checking Marshalling error",
			"error": "Marshalling",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnMarshalling",
		ArrangeInput: args.Map{
			"when":  "checking UnMarshalling error",
			"error": "UnMarshalling",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- OutOfRange",
		ArrangeInput: args.Map{
			"when":  "checking OutOfRange error",
			"error": "OutOfRange",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- CannotProcessNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking CannotProcessNilOrEmpty error",
			"error": "CannotProcessNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- NegativeDataCannotProcess",
		ArrangeInput: args.Map{
			"when":  "checking NegativeDataCannotProcess error",
			"error": "NegativeDataCannotProcess",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- NilResult",
		ArrangeInput: args.Map{
			"when":  "checking NilResult error",
			"error": "NilResult",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnexpectedValue",
		ArrangeInput: args.Map{
			"when":  "checking UnexpectedValue error",
			"error": "UnexpectedValue",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- CannotRemoveFromEmptyCollection",
		ArrangeInput: args.Map{
			"when":  "checking CannotRemoveFromEmptyCollection error",
			"error": "CannotRemoveFromEmptyCollection",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- MarshallingFailedDueToNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking MarshallingFailedDueToNilOrEmpty error",
			"error": "MarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- UnmarshallingFailedDueToNilOrEmpty",
		ArrangeInput: args.Map{
			"when":  "checking UnmarshallingFailedDueToNilOrEmpty error",
			"error": "UnmarshallingFailedDueToNilOrEmpty",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
	{
		Title: "DefaultErr returns non-nil error with message -- KeyNotExistInMap",
		ArrangeInput: args.Map{
			"when":  "checking KeyNotExistInMap error",
			"error": "KeyNotExistInMap",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
			"hasMessage": true,
		},
	},
}
