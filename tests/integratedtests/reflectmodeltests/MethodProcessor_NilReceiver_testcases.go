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

package reflectmodeltests

import (
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
	"github.com/alimtvnetwork/core-v8/reflectcore/reflectmodel"
)

// =============================================================================
// MethodProcessor nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var methodProcessorNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasValidFunc on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).HasValidFunc,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.MethodProcessor).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Func on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).Func,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "ReturnLength on nil returns -1",
		Func:  (*reflectmodel.MethodProcessor).ReturnLength,
		Expected: results.ResultAny{
			Value:    "-1",
			Panicked: false,
		},
	},
	{
		Title: "IsPublicMethod on nil returns false",
		Func:  (*reflectmodel.MethodProcessor).IsPublicMethod,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "GetType on nil returns nil",
		Func:  (*reflectmodel.MethodProcessor).GetType,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "GetInArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypes,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "GetOutArgsTypes on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetOutArgsTypes,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "GetInArgsTypesNames on nil returns empty",
		Func:  (*reflectmodel.MethodProcessor).GetInArgsTypesNames,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "Invoke on nil returns error",
		Func:  (*reflectmodel.MethodProcessor).Invoke,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
}
