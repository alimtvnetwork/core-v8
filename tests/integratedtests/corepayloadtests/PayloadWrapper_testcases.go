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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================
// PayloadWrapper — Create and Serialize
// ==========================================

var payloadWrapperCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "PayloadWrapper returns valid JSON -- name 'test-payload', id 'pay-1'",
		ArrangeInput: args.Map{
			"when": "given payload with name and id",
			"name": "test-payload",
			"id":   "pay-1",
		},
		ExpectedInput: args.Map{
			"name":       "test-payload",
			"identifier": "pay-1",
			"hasJson":    true,
		},
	},
}

// ==========================================
// PayloadWrapper — Deserialization roundtrip
// ==========================================

var payloadWrapperDeserializeRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "PayloadWrapper returns preserved data -- roundtrip name 'roundtrip-payload', id 'rt-1'",
		ArrangeInput: args.Map{
			"when": "given payload serialized then deserialized",
			"name": "roundtrip-payload",
			"id":   "rt-1",
		},
		ExpectedInput: args.Map{
			"restoredName":       "roundtrip-payload",
			"restoredIdentifier": "rt-1",
			"jsonIsEqual":        true,
		},
	},
}

// ==========================================
// PayloadWrapper — Deep Clone
// ==========================================

var payloadWrapperCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr returns independent copy -- original 'original-pay' mutated to 'mutated-pay'",
		ArrangeInput: args.Map{
			"when":     "given payload cloned and mutated",
			"name":     "original-pay",
			"id":       "clone-1",
			"new_name": "mutated-pay",
		},
		ExpectedInput: args.Map{
			"originalName":  "original-pay",
			"clonedName":    "mutated-pay",
			"isIndependent": true,
		},
	},
}

// ==========================================
// PayloadWrapper — DeserializeToMany
// ==========================================

var payloadWrapperDeserializeToManyTestCases = []coretestcases.CaseV1{
	{
		Title: "DeserializeToMany returns 3 payloads -- array of 3 serialized",
		ArrangeInput: args.Map{
			"when":  "given 3 payloads serialized as array",
			"count": 3,
		},
		ExpectedInput: args.Map{
			"deserializedCount": 3,
		},
	},
}
