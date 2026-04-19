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

package casenilsafetests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// Nil-safe pointer receiver methods (should NOT panic)
// =============================================================================

var nilSafePointerReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsValid on nil returns false",
		Func:  (*sampleStruct).IsValid,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
	{
		Title: "Length on nil returns 0",
		Func:  (*sampleStruct).Length,
		Expected: results.ResultAny{
			Value:    "0",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
	{
		Title: "String on nil returns empty",
		Func:  (*sampleStruct).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
		CompareFields: []string{"value", "panicked", "isSafe"},
	},
}

// =============================================================================
// Void methods (no return values)
// =============================================================================

var nilSafeVoidTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Reset on nil does not panic",
		Func:  (*sampleStruct).Reset,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked", "returnCount"},
	},
}

// =============================================================================
// Multi-return methods
// =============================================================================

var nilSafeMultiReturnTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "Parse on nil returns (0, nil)",
		Func:  (*sampleStruct).Parse,
		Args:  []any{"hello"},
		Expected: results.ResultAny{
			Value:       "0",
			Panicked:    false,
			ReturnCount: 2,
		},
		CompareFields: []string{"value", "panicked", "hasError", "returnCount"},
	},
	{
		Title: "Lookup on nil returns empty false",
		Func:  (*sampleStruct).Lookup,
		Args:  []any{"key"},
		Expected: results.ResultAny{
			Value:       "",
			Panicked:    false,
			ReturnCount: 2,
		},
	},
}

// =============================================================================
// Unsafe methods (SHOULD panic on nil)
// =============================================================================

var nilUnsafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "UnsafeMethod on nil panics",
		Func:  (*sampleStruct).UnsafeMethod,
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
	{
		Title: "ValueString on nil panics (value receiver)",
		Func:  (*sampleStruct).ValueString,
		Expected: results.ResultAny{
			Panicked: true,
		},
	},
}

// =============================================================================
// MethodName extraction — uses CompareFields for non-standard assertions
// =============================================================================

var methodNameTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "MethodName extracts IsValid",
		Func:  (*sampleStruct).IsValid,
		Expected: results.ResultAny{
			Value: "IsValid",
		},
	},
	{
		Title: "MethodName extracts Parse",
		Func:  (*sampleStruct).Parse,
		Args:  []any{"x"},
		Expected: results.ResultAny{
			Value: "Parse",
		},
	},
	{
		Title: "MethodName extracts Reset",
		Func:  (*sampleStruct).Reset,
		Expected: results.ResultAny{
			Value: "Reset",
		},
	},
}
