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

package codefuncstests

import (
	"fmt"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// GetFuncName — positive + negative
// =============================================================================

var getFuncNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncName returns non-empty name -- named function input",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: "true",
	},
}

var getFuncNameNilTestCase = coretestcases.CaseV1{
	Title: "GetFuncName returns empty string -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncName returns empty string -- non-function (int) input",
	ArrangeInput: args.Map{
		"when":  "given an integer instead of a function",
		"input": 42,
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

// =============================================================================
// GetFuncFullName — positive + negative
// =============================================================================

var getFuncFullNameTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFuncFullName returns full package-qualified name -- named function input",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: args.Map{
			"isNotEmpty":      true,
			"containsPackage": true,
		},
	},
}

var getFuncFullNameNilTestCase = coretestcases.CaseV1{
	Title: "GetFuncFullName returns empty string -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

var getFuncFullNameNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFuncFullName returns empty string -- non-function (string) input",
	ArrangeInput: args.Map{
		"when":  "given a string instead of a function",
		"input": "not-a-func",
	},
	ExpectedInput: args.Map{
		"result":   "",
		"panicked": false,
	},
}

// =============================================================================
// GetFunc — positive + negative
// =============================================================================

var getFuncTestCases = []coretestcases.CaseV1{
	{
		Title: "GetFunc returns non-nil -- named function input",
		ArrangeInput: args.Map{
			"when": "given a named function",
		},
		ExpectedInput: args.Map{
			"isNotNil": true,
		},
	},
}

var getFuncNilTestCase = coretestcases.CaseV1{
	Title: "GetFunc returns nil -- nil input",
	ArrangeInput: args.Map{
		"when": "given nil input",
	},
	ExpectedInput: args.Map{
		"isNil":    true,
		"panicked": false,
	},
}

var getFuncNonFuncTestCase = coretestcases.CaseV1{
	Title: "GetFunc returns nil -- non-function (struct) input",
	ArrangeInput: args.Map{
		"when":  "given a struct instead of a function",
		"input": struct{ Name string }{"test"},
	},
	ExpectedInput: args.Map{
		"isNil":    true,
		"panicked": false,
	},
}

// =============================================================================
// newCreator — factory methods
// =============================================================================

var newCreatorTestCases = []coretestcases.CaseV1{
	{
		Title: "New.ActionErr returns wrapper with correct name -- 'test-action' factory",
		ArrangeInput: args.Map{
			"method": "ActionErr",
			"name":   "test-action",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "New.IsSuccess returns wrapper that returns true -- 'test-check' factory",
		ArrangeInput: args.Map{
			"method": "IsSuccess",
			"name":   "test-check",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "New.NamedAction returns wrapper that calls with name -- 'test-named' factory",
		ArrangeInput: args.Map{
			"method": "NamedAction",
			"name":   "test-named",
		},
		ExpectedInput: args.Map{
			"calledWith": "test-named",
		},
	},
	{
		Title: "New.LegacyInOutErr returns wrapper with output 'processed' -- 'test-inout' factory",
		ArrangeInput: args.Map{
			"method": "LegacyInOutErr",
			"name":   "test-inout",
		},
		ExpectedInput: args.Map{
			"output":   "processed",
			"hasError": false,
		},
	},
	{
		Title: "New.LegacyResultDelegating returns wrapper without error -- 'test-delegate' factory",
		ArrangeInput: args.Map{
			"method": "LegacyResultDelegating",
			"name":   "test-delegate",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// Ensure fmt is used
var _ = fmt.Sprintf
