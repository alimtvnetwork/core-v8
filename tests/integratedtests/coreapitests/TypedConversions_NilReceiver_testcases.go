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
	"github.com/alimtvnetwork/core-v8/coredata/coreapi"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
)

// =============================================================================
// TypedSimpleGenericRequest nil receiver test cases
// (migrated from CaseV1 string-dispatch in TypedConversions_testcases.go)
//
// Note: Go does not support method expressions on generic types directly.
// We use function literal wrappers to achieve compile-time safety.
// Renaming the method still causes a build error at the call site.
// =============================================================================

var typedSimpleGenericRequestNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Nil receiver IsValid returns false",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) bool {
			return r.IsValid()
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver IsInvalid returns true",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) bool {
			return r.IsInvalid()
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver Message returns empty string",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) string {
			return r.Message()
		},
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "Nil receiver InvalidError returns nil",
		Func: func(r *coreapi.TypedSimpleGenericRequest[string]) error {
			return r.InvalidError()
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "hasError"},
	},
}
