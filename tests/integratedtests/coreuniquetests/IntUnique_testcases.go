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

package coreuniquetests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var intUniqueGetRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title: "Get returns unique count -- slice with duplicates [1,2,2,3,3,3]",
	ArrangeInput: args.Map{
		"when":  "given slice with duplicates",
		"input": []int{1, 2, 2, 3, 3, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetAlreadyUniqueTestCase = coretestcases.CaseV1{
	Title: "Get returns same count -- already unique slice [1,2,3]",
	ArrangeInput: args.Map{
		"when":  "given slice without duplicates",
		"input": []int{1, 2, 3},
	},
	ExpectedInput: "3",
}

var intUniqueGetNilTestCase = coretestcases.CaseV1{
	Title: "Get returns nil-safe result -- nil slice input",
	ArrangeInput: args.Map{
		"when": "given nil slice",
	},
	ExpectedInput: "true",
}
