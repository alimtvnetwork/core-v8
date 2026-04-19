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

package coredynamictests

import (
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// CastedResult nil receiver test cases
// (migrated from first element of CaseV1 slices in CastedResult_testcases.go)
// =============================================================================

var castedResultNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsInvalid true on nil receiver",
		Func:  (*coredynamic.CastedResult).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsNotNull false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotNull,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsNotPointer false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotPointer,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsNotMatchingAcceptedType false on nil receiver",
		Func:  (*coredynamic.CastedResult).IsNotMatchingAcceptedType,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasError false on nil receiver",
		Func:  (*coredynamic.CastedResult).HasError,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyIssues true on nil receiver",
		Func:  (*coredynamic.CastedResult).HasAnyIssues,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsSourceKind false on nil receiver",
		Func: func(cr *coredynamic.CastedResult) bool {
			return cr.IsSourceKind(0)
		},
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
