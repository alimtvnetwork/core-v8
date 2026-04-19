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

package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// Attributes.IsEqual — Regression for logic inversion bug
// ==========================================

var attributesIsEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "Attributes.IsEqual returns true -- both nil",
		ArrangeInput: args.Map{
			"when":      "both attributes are nil",
			"left_nil":  true,
			"right_nil": true,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Attributes.IsEqual returns false -- left nil right non-nil",
		ArrangeInput: args.Map{
			"when":      "left is nil right is not",
			"left_nil":  true,
			"right_nil": false,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Attributes.IsEqual returns false -- right nil left non-nil",
		ArrangeInput: args.Map{
			"when":      "right is nil left is not",
			"left_nil":  false,
			"right_nil": true,
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
	{
		Title: "Attributes.IsEqual returns true -- same pointer identity",
		ArrangeInput: args.Map{
			"when":         "same pointer identity",
			"same_pointer": true,
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Attributes.IsEqual returns true -- same dynamic payloads",
		ArrangeInput: args.Map{
			"when":    "same dynamic payloads on both",
			"payload": "test-data",
		},
		ExpectedInput: args.Map{
			"isEqual": true,
		},
	},
	{
		Title: "Attributes.IsEqual returns false -- different dynamic payloads",
		ArrangeInput: args.Map{
			"when":          "different dynamic payloads",
			"left_payload":  "data-a",
			"right_payload": "data-b",
		},
		ExpectedInput: args.Map{
			"isEqual": false,
		},
	},
}

// ==========================================
// Attributes.Clone — Regression for deep clone independence
// ==========================================

var attributesCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Attributes.Clone returns nil -- nil attributes shallow clone",
		ArrangeInput: args.Map{
			"when":     "attributes is nil",
			"nil_attr": true,
			"deep":     false,
		},
		ExpectedInput: args.Map{
			"isNil":    true,
			"hasError": false,
		},
	},
	{
		Title: "Attributes.Clone returns preserved payload -- shallow clone",
		ArrangeInput: args.Map{
			"when":    "shallow clone with dynamic payloads",
			"payload": "clone-payload",
			"deep":    false,
		},
		ExpectedInput: args.Map{
			"clonedPayload": "clone-payload",
			"isEqual":       true,
		},
	},
	{
		Title: "Attributes.Clone returns error -- deep clone",
		ArrangeInput: args.Map{
			"when":    "deep clone then mutate original",
			"payload": "deep-clone-data",
			"deep":    true,
		},
		ExpectedInput: args.Map{
			"clonedPayload": "deep-clone-data",
			"isEqual":       true,
		},
	},
}

// ==========================================
// Attributes.IsSafeValid — Regression for negation bug
// ==========================================

var attributesIsSafeValidTestCases = []coretestcases.CaseV1{
	{
		Title: "Attributes.IsSafeValid returns false -- nil attributes",
		ArrangeInput: args.Map{
			"when":     "attributes is nil",
			"nil_attr": true,
		},
		ExpectedInput: args.Map{
			"isSafeValid": false,
		},
	},
	{
		Title: "Attributes.IsSafeValid returns false -- empty attributes",
		ArrangeInput: args.Map{
			"when":  "attributes has no data",
			"empty": true,
		},
		ExpectedInput: args.Map{
			"isSafeValid": false,
		},
	},
	{
		Title: "Attributes.IsSafeValid returns true -- has dynamic payload",
		ArrangeInput: args.Map{
			"when":    "attributes has dynamic payload",
			"payload": "valid-data",
		},
		ExpectedInput: args.Map{
			"isSafeValid": true,
		},
	},
}

// ==========================================
// AuthInfo.Clone — Regression for missing Identifier field
// ==========================================

var authInfoCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "AuthInfo.ClonePtr returns nil -- nil receiver",
		ArrangeInput: args.Map{
			"when":     "auth info is nil",
			"nil_auth": true,
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "AuthInfo.ClonePtr returns preserved Identifier -- populated auth info",
		ArrangeInput: args.Map{
			"when":          "auth info has identifier",
			"identifier":    "user-42",
			"action_type":   "login",
			"resource_name": "/api/data",
		},
		ExpectedInput: args.Map{
			"identifier":   "user-42",
			"actionType":   "login",
			"resourceName": "/api/data",
		},
	},
	{
		Title: "AuthInfo.ClonePtr returns independent copy -- clone mutated after creation",
		ArrangeInput: args.Map{
			"when":            "clone mutated after creation",
			"identifier":      "original-id",
			"action_type":     "read",
			"new_action_type": "write",
		},
		ExpectedInput: args.Map{
			"originalAction": "read",
			"clonedAction":   "write",
		},
	},
	{
		Title: "AuthInfo.ClonePtr returns preserved fields -- empty Identifier",
		ArrangeInput: args.Map{
			"when":          "identifier is empty string",
			"identifier":    "",
			"action_type":   "delete",
			"resource_name": "/api/remove",
		},
		ExpectedInput: args.Map{
			"identifier":   "",
			"actionType":   "delete",
			"resourceName": "/api/remove",
		},
	},
}
