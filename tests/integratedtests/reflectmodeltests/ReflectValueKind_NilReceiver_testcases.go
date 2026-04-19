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
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
)

// =============================================================================
// ReflectValueKind nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var reflectValueKindNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsInvalid on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasError on nil returns false",
		Func:  (*reflectmodel.ReflectValueKind).HasError,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyError on nil returns true",
		Func:  (*reflectmodel.ReflectValueKind).IsEmptyError,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ActualInstance on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).ActualInstance,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
	{
		Title: "PkgPath on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).PkgPath,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "TypeName on nil returns empty",
		Func:  (*reflectmodel.ReflectValueKind).TypeName,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "PointerRv on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerRv,
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "PointerInterface on nil returns nil",
		Func:  (*reflectmodel.ReflectValueKind).PointerInterface,
		Expected: results.ResultAny{
			Value:    "<nil>",
			Panicked: false,
		},
	},
}
