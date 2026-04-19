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

package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ══════════════════════════════════════════════════════════════════════════════
// MapStringStringOnce.JsonStringMust() — success path
// Covers MapStringStringOnce.go L306-317 (non-error path already covered,
// error panic path L309-314 is unreachable with valid map data)
// ══════════════════════════════════════════════════════════════════════════════

var cov13MapStringStringOnceJsonStringMustTestCase = coretestcases.CaseV1{
	Title: "JsonStringMust returns valid JSON -- non-empty map",
	ExpectedInput: args.Map{
		"nonEmpty": true,
		"noPanic":  true,
	},
}

// ══════════════════════════════════════════════════════════════════════════════
// StringsOnce.JsonStringMust() — success path
// Covers StringsOnce.go L248-258 (non-error path already covered,
// error panic path L251-256 is unreachable with valid string slice data)
// ══════════════════════════════════════════════════════════════════════════════

var cov13StringsOnceJsonStringMustTestCase = coretestcases.CaseV1{
	Title: "JsonStringMust returns valid JSON -- non-empty strings",
	ExpectedInput: args.Map{
		"nonEmpty": true,
		"noPanic":  true,
	},
}
