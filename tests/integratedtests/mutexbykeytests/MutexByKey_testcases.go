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

package mutexbykeytests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var getAndDeleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Get returns non-nil mutex -- new key 'test-key-1'",
		ArrangeInput: args.Map{
			"when": "given a new key",
			"key":  "test-key-1",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Get returns same mutex -- same key 'test-key-same' requested twice",
		ArrangeInput: args.Map{
			"when": "given same key twice",
			"key":  "test-key-same",
		},
		ExpectedInput: "true",
	},
}

var deleteTestCases = []coretestcases.CaseV1{
	{
		Title: "Delete returns true -- existing key 'test-key-del'",
		ArrangeInput: args.Map{
			"when": "given existing key to delete",
			"key":  "test-key-del",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Delete returns false -- non-existing key 'test-key-nonexistent'",
		ArrangeInput: args.Map{
			"when": "given non-existing key to delete",
			"key":  "test-key-nonexistent",
		},
		ExpectedInput: "false",
	},
}
