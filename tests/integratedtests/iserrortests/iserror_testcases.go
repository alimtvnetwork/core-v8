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

package iserrortests

import (
	"errors"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var (
	errSample1 = errors.New("sample error 1")
	errSample2 = errors.New("sample error 2")
	errSame    = errors.New("same message")
	errSameDup = errors.New("same message")
)

// ==========================================
// Empty / Defined / NotEmpty
// ==========================================

var emptyTestCases = []coretestcases.CaseV1{
	{
		Title: "Empty returns true for nil error",
		ArrangeInput: args.Map{
			"when": "given nil error",
		},
		ExpectedInput: args.Map{
			"isEmpty":    "true",
			"isDefined":  "false",
			"isNotEmpty": "false",
		},
	},
	{
		Title: "Empty returns false for non-nil error",
		ArrangeInput: args.Map{
			"when": "given non-nil error",
		},
		ExpectedInput: args.Map{
			"isEmpty":    "false",
			"isDefined":  "true",
			"isNotEmpty": "true",
		},
	},
}

// ==========================================
// Equal / NotEqual
// ==========================================

var equalTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal returns true for same error instance",
		ArrangeInput: args.Map{
			"when": "given same error on both sides",
		},
		ExpectedInput: args.Map{
			"isEqual":    "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "Equal returns true for both nil",
		ArrangeInput: args.Map{
			"when": "given both nil",
		},
		ExpectedInput: args.Map{
			"isEqual":    "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "Equal returns false for nil vs non-nil",
		ArrangeInput: args.Map{
			"when": "given nil vs non-nil",
		},
		ExpectedInput: args.Map{
			"isEqual":    "false",
			"isNotEqual": "true",
		},
	},
	{
		Title: "Equal returns true for same message different instances",
		ArrangeInput: args.Map{
			"when": "given same message different instances",
		},
		ExpectedInput: args.Map{
			"isEqual":    "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "Equal returns false for different messages",
		ArrangeInput: args.Map{
			"when": "given different messages",
		},
		ExpectedInput: args.Map{
			"isEqual":    "false",
			"isNotEqual": "true",
		},
	},
}

// ==========================================
// AllDefined / AnyDefined
// ==========================================

var allDefinedTestCases = []coretestcases.CaseV1{
	{
		Title: "AllDefined true when all errors are non-nil",
		ArrangeInput: args.Map{
			"when": "given all non-nil errors",
		},
		ExpectedInput: "true",
	},
	{
		Title: "AllDefined false when one is nil",
		ArrangeInput: args.Map{
			"when": "given one nil error among non-nil",
		},
		ExpectedInput: "false",
	},
	{
		Title: "AllDefined false for empty args",
		ArrangeInput: args.Map{
			"when": "given no arguments",
		},
		ExpectedInput: "false",
	},
}

var anyDefinedTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyDefined true when at least one non-nil",
		ArrangeInput: args.Map{
			"when": "given one non-nil among nils",
		},
		ExpectedInput: "true",
	},
	{
		Title: "AnyDefined false when all nil",
		ArrangeInput: args.Map{
			"when": "given all nil errors",
		},
		ExpectedInput: "false",
	},
	{
		Title: "AnyDefined false for empty args",
		ArrangeInput: args.Map{
			"when": "given no arguments",
		},
		ExpectedInput: "false",
	},
}

// ==========================================
// AllEmpty / AnyEmpty
// ==========================================

var allEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "AllEmpty true when all errors are nil",
		ArrangeInput: args.Map{
			"when": "given all nil errors",
		},
		ExpectedInput: "true",
	},
	{
		Title: "AllEmpty false when one is non-nil",
		ArrangeInput: args.Map{
			"when": "given one non-nil among nil",
		},
		ExpectedInput: "false",
	},
	{
		Title: "AllEmpty true for empty args",
		ArrangeInput: args.Map{
			"when": "given no arguments",
		},
		ExpectedInput: "true",
	},
}

var anyEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "AnyEmpty true when at least one nil",
		ArrangeInput: args.Map{
			"when": "given one nil among non-nil",
		},
		ExpectedInput: "true",
	},
	{
		Title: "AnyEmpty false when all non-nil",
		ArrangeInput: args.Map{
			"when": "given all non-nil errors",
		},
		ExpectedInput: "false",
	},
	{
		Title: "AnyEmpty true for empty args",
		ArrangeInput: args.Map{
			"when": "given no arguments",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// EqualString / NotEqualString
// ==========================================

var equalStringTestCases = []coretestcases.CaseV1{
	{
		Title: "EqualString true for same strings",
		ArrangeInput: args.Map{
			"when":  "given identical strings",
			"left":  "hello",
			"right": "hello",
		},
		ExpectedInput: args.Map{
			"isEqual":    "true",
			"isNotEqual": "false",
		},
	},
	{
		Title: "EqualString false for different strings",
		ArrangeInput: args.Map{
			"when":  "given different strings",
			"left":  "hello",
			"right": "world",
		},
		ExpectedInput: args.Map{
			"isEqual":    "false",
			"isNotEqual": "true",
		},
	},
	{
		Title: "EqualString true for empty strings",
		ArrangeInput: args.Map{
			"when":  "given both empty",
			"left":  "",
			"right": "",
		},
		ExpectedInput: args.Map{
			"isEqual":    "true",
			"isNotEqual": "false",
		},
	},
}
