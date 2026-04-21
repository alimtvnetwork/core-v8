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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var resultUnmarshalValidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns deserialized struct -- valid JSON input",
	ExpectedInput: args.Map{
		"error":            "<nil>",
		"deserializedName": "Alice",
		"deserializedAge":  "30",
	},
}

var resultUnmarshalNilTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns error -- nil receiver",
	ExpectedInput: args.Map{
		"hasError":          true,
		"errorContainsNull": true,
	},
}

var resultUnmarshalInvalidTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns error -- invalid bytes input",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}

var resultUnmarshalExistingErrorTestCase = coretestcases.CaseV1{
	Title: "Unmarshal returns propagated error -- existing error on result",
	ExpectedInput: args.Map{
		"hasError":               true,
		"errorContainsUnmarshal": true,
	},
}
