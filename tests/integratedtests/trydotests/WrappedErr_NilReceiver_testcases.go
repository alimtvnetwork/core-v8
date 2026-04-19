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

package trydotests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/internal/trydo"
)

// =============================================================================
// WrappedErr nil receiver test cases
// (migrated from first element of CaseV1 slices in WrappedErr_testcases.go)
// =============================================================================

var wrappedErrNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsDefined on nil returns false",
		Func:  (*trydo.WrappedErr).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*trydo.WrappedErr).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ErrorString on nil returns empty",
		Func:  (*trydo.WrappedErr).ErrorString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExceptionString on nil returns empty",
		Func:  (*trydo.WrappedErr).ExceptionString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*trydo.WrappedErr).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "ExceptionType on nil returns nil",
		Func: func(w *trydo.WrappedErr) bool {
			return w.ExceptionType() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalidException on nil returns true",
		Func:  (*trydo.WrappedErr).IsInvalidException,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasErrorOrException on nil returns false",
		Func:  (*trydo.WrappedErr).HasErrorOrException,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsBothPresent on nil returns false",
		Func:  (*trydo.WrappedErr).IsBothPresent,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasException on nil returns false",
		Func:  (*trydo.WrappedErr).HasException,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
}
