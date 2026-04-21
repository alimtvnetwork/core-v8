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
	"errors"
	"io"
	"os"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// Sentinel errors for chain verification
var (
	errSentinelA = errors.New("sentinel-a")
	errSentinelB = errors.New("sentinel-b")
	errSentinelC = errors.New("sentinel-c")
)

// =============================================================================
// MergeErrors + errors.Is
// =============================================================================

var mergeErrorsIsTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors preserves sentinel via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given sentinel among multiple errors",
			"sentinel": errSentinelA,
			"errors":   []error{errors.New("other"), errSentinelA, errors.New("another")},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "true",
		},
	},
	{
		Title: "MergeErrors single error matches itself via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given single sentinel",
			"sentinel": errSentinelB,
			"errors":   []error{errSentinelB},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "true",
		},
	},
	{
		Title: "MergeErrors does not match absent sentinel",
		ArrangeInput: args.Map{
			"when":     "given sentinel not in the list",
			"sentinel": errSentinelC,
			"errors":   []error{errSentinelA, errSentinelB},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr + errors.Is
// =============================================================================

var concatMessageErrorsIsTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr preserves sentinel via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given sentinel wrapped with message",
			"sentinel": errSentinelA,
			"message":  "context:",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"errorsIsOk":      "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "ConcatMessageWithErr preserves io.EOF via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given io.EOF wrapped",
			"sentinel": io.EOF,
			"message":  "read failed:",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"errorsIsOk":      "true",
			"containsMessage": "true",
		},
	},
}

// =============================================================================
// MergeErrors + errors.As
// =============================================================================

var mergeErrorsAsTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors preserves *os.PathError via errors.As",
		ArrangeInput: args.Map{
			"when": "given PathError among plain errors",
			"errors": []error{
				errors.New("plain"),
				&os.PathError{Op: "open", Path: "/tmp/test", Err: errors.New("not found")},
			},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "true",
		},
	},
	{
		Title: "MergeErrors errors.As returns false when type absent",
		ArrangeInput: args.Map{
			"when":   "given only plain errors",
			"errors": []error{errors.New("a"), errors.New("b")},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr + errors.As
// =============================================================================

var concatMessageErrorsAsTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr preserves *os.PathError via errors.As",
		ArrangeInput: args.Map{
			"when":    "given PathError wrapped with message",
			"error":   &os.PathError{Op: "read", Path: "/etc/cfg", Err: errors.New("denied")},
			"message": "config load:",
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "true",
		},
	},
	{
		Title: "ConcatMessageWithErr errors.As false for plain error",
		ArrangeInput: args.Map{
			"when":    "given plain error",
			"error":   errors.New("simple"),
			"message": "wrap:",
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr nil passthrough
// =============================================================================

var concatMessageNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr nil returns nil",
		ArrangeInput: args.Map{
			"when":    "given nil error",
			"message": "should not appear",
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// MergeErrors multiple sentinels
// =============================================================================

var mergeErrorsMultiSentinelTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors all three sentinels found via errors.Is",
		ArrangeInput: args.Map{
			"when":      "given three distinct sentinels",
			"errors":    []error{errSentinelA, errSentinelB, errSentinelC},
			"sentinels": []error{errSentinelA, errSentinelB, errSentinelC},
		},
		ExpectedInput: args.Map{
			"hasError":       "true",
			"allSentinelsOk": "true",
		},
	},
	{
		Title: "MergeErrors partial sentinels -- missing one fails",
		ArrangeInput: args.Map{
			"when":      "given two of three sentinels",
			"errors":    []error{errSentinelA, errSentinelB},
			"sentinels": []error{errSentinelA, errSentinelB, errSentinelC},
		},
		ExpectedInput: args.Map{
			"hasError":       "true",
			"allSentinelsOk": "false",
		},
	},
}
