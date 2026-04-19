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

package coreapitests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// TypedRequest — NewTypedRequest
// ==========================================

var typedRequestNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedRequest creates valid typed request",
		ArrangeInput: args.Map{
			"when":    "given string payload",
			"payload": "my-request",
		},
		ExpectedInput: args.Map{
			"payload": "my-request",
			"isValid": "true",
		},
	},
}

// ==========================================
// TypedRequest — InvalidTypedRequest
// ==========================================

var typedRequestInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedRequest creates request with zero-value payload",
		ArrangeInput: args.Map{
			"when": "given nil attribute",
		},
		ExpectedInput: args.Map{
			"payload":      "",
			"isValid":      "false",
			"hasAttribute": "true",
		},
	},
}

// ==========================================
// TypedRequest — Clone
// ==========================================

var typedRequestCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy of TypedRequest",
		ArrangeInput: args.Map{
			"when":    "given valid request",
			"payload": "clone-payload",
		},
		ExpectedInput: args.Map{
			"payload":       "clone-payload",
			"isValid":       "true",
			"isIndependent": "true",
		},
	},
}

var typedRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedRequest returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// TypedResponseResult — ToTypedResponse
// ==========================================

var typedResponseResultToTypedResponseTestCases = []coretestcases.CaseV1{
	{
		Title: "ToTypedResponse converts result back to TypedResponse",
		ArrangeInput: args.Map{
			"when":     "given valid typed response result",
			"response": "back-convert",
		},
		ExpectedInput: args.Map{
			"response": "back-convert",
			"isValid":  "true",
		},
	},
}

var typedResponseResultToTypedResponseNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToTypedResponse on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// TypedResponseResult — Message
// ==========================================

var typedResponseResultMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "Message returns attribute message",
		ArrangeInput: args.Map{
			"when":     "given result with message",
			"response": "data",
			"message":  "operation completed",
		},
		ExpectedInput: "operation completed",
	},
	{
		Title: "Message returns empty string on nil result",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: "",
	},
}
