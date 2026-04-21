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
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IsSuccessFuncWrapper — Exec
// =============================================================================

var isSuccessExecTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSuccessFuncWrapper.Exec returns true -- action succeeds",
		ArrangeInput: args.Map{
			"actionResult": true,
			"name":         "check-health",
		},
		ExpectedInput: args.Map{
			"result": true,
		},
	},
	{
		Title: "IsSuccessFuncWrapper.Exec returns false -- action fails",
		ArrangeInput: args.Map{
			"actionResult": false,
			"name":         "check-health",
		},
		ExpectedInput: args.Map{
			"result": false,
		},
	},
}

// =============================================================================
// IsSuccessFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var isSuccessAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "IsSuccessFuncWrapper.AsActionReturnsErrorFunc returns nil -- action succeeds",
		ArrangeInput: args.Map{
			"actionResult": true,
			"name":         "check-ok",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "IsSuccessFuncWrapper.AsActionReturnsErrorFunc returns error with name -- action fails",
		ArrangeInput: args.Map{
			"actionResult": false,
			"name":         "check-fail",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
}

// =============================================================================
// NamedActionFuncWrapper — Exec
// =============================================================================

var namedActionExecTestCases = []coretestcases.CaseV1{
	{
		Title: "NamedActionFuncWrapper.Exec returns calledWith 'my-task' -- wrapper name 'my-task'",
		ArrangeInput: args.Map{
			"name": "my-task",
		},
		ExpectedInput: args.Map{
			"calledWith": "my-task",
		},
	},
}

// =============================================================================
// NamedActionFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var namedActionAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "NamedActionFuncWrapper.AsActionReturnsErrorFunc returns nil -- always succeeds",
		ArrangeInput: args.Map{
			"name": "log-step",
		},
		ExpectedInput: args.Map{
			"hasError":   false,
			"calledWith": "log-step",
		},
	},
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — Exec
// =============================================================================

var actionReturnsErrorExecTestCases = []coretestcases.CaseV1{
	{
		Title: "ActionReturnsErrorFuncWrapper.Exec returns nil -- action succeeds",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "cleanup",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "ActionReturnsErrorFuncWrapper.Exec returns error -- action fails",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "cleanup",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// ActionReturnsErrorFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var actionReturnsErrorAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "ActionReturnsErrorFuncWrapper.AsActionReturnsErrorFunc returns error with name -- failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "deploy",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "ActionReturnsErrorFuncWrapper.AsActionReturnsErrorFunc returns nil -- success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "deploy",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapper — Exec
// =============================================================================

var inOutErrExecTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapper.Exec returns 'HELLO' -- success with 'hello' input",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": false,
			"name":         "transform",
		},
		ExpectedInput: args.Map{
			"output":   "HELLO",
			"hasError": false,
		},
	},
	{
		Title: "InOutErrFuncWrapper.Exec returns error -- failure with 'hello' input",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": true,
			"name":         "transform",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var inOutErrAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "InOutErrFuncWrapper.AsActionReturnsErrorFunc returns error with name -- failure",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": true,
			"name":         "process",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "InOutErrFuncWrapper.AsActionReturnsErrorFunc returns nil -- success",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": false,
			"name":         "process",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapper — Exec
// =============================================================================

var resultDelegatingExecTestCases = []coretestcases.CaseV1{
	{
		Title: "ResultDelegatingFuncWrapper.Exec returns nil -- success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "unmarshal",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "ResultDelegatingFuncWrapper.Exec returns error -- failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "unmarshal",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapper — AsActionReturnsErrorFunc
// =============================================================================

var resultDelegatingAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "ResultDelegatingFuncWrapper.AsActionReturnsErrorFunc returns error with name -- failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "deserialize",
		},
		ExpectedInput: args.Map{
			"hasError":     true,
			"containsName": true,
		},
	},
	{
		Title: "ResultDelegatingFuncWrapper.AsActionReturnsErrorFunc returns nil -- success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "deserialize",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}
