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

package errcoretests

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// ConcatMessageWithErr nil passthrough test case
//
// Note: ConcatMessageWithErr is a package function, not a method.
// CaseNilSafe is inappropriate here — using CaseV1 instead.
// =============================================================================

var concatMessageWithErrNilPassthroughTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErr nil error returns nil",
	ArrangeInput: args.Map{
		"message": "should not appear",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var concatMessageWithErrNonNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErr non-nil error wraps message",
	ArrangeInput: args.Map{
		"message": "context:",
		"error":   "original error",
	},
	ExpectedInput: args.Map{
		"isNil":    false,
		"contains": true,
	},
}

var concatMessageWithErrWithStackTraceNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErrWithStackTrace nil error returns nil",
	ArrangeInput: args.Map{
		"message": "should not appear",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var concatMessageWithErrWithStackTraceNonNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErrWithStackTrace non-nil error wraps message",
	ArrangeInput: args.Map{
		"message": "context:",
		"error":   "original error",
	},
	ExpectedInput: args.Map{
		"isNil":    false,
		"contains": true,
	},
}

// =============================================================================
// Helper for constructing error from string
// =============================================================================

func errFromString(input args.Map) error {
	errStr, hasErr := input.GetAsString("error")

	if !hasErr {
		return nil
	}

	return fmt.Errorf("%s", errStr)
}
