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

package regexnewtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
	"github.com/alimtvnetwork/core-v8/regexnew"
)

// =============================================================================
// LazyRegex nil receiver test cases (migrated from inline t.Error tests)
// =============================================================================

var lazyRegexNilReceiverTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "LazyRegex.IsNull returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsNull,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsUndefined returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsUndefined,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsDefined returns false -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsDefined,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsCompiled returns false -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsCompiled,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.String returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).String,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.Pattern returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).Pattern,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.HasAnyIssues returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).HasAnyIssues,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.IsInvalid returns true -- nil receiver",
		Func:  (*regexnew.LazyRegex).IsInvalid,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.OnRequiredCompiled returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).OnRequiredCompiled,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "LazyRegex.FullString returns empty -- nil receiver",
		Func:  (*regexnew.LazyRegex).FullString,
		Expected: results.ResultAny{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "LazyRegex.CompiledError returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).CompiledError,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
	{
		Title: "LazyRegex.Error returns error -- nil receiver",
		Func:  (*regexnew.LazyRegex).Error,
		Expected: results.ResultAny{
			Panicked: false,
			Error:    results.ExpectAnyError,
		},
	},
}
