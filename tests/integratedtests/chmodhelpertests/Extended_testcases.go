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

package chmodhelpertests

import (
	"github.com/alimtvnetwork/core-v8/chmodhelper"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
)

type pathExistStatType = chmodhelper.PathExistStat
type nilSafeResult = results.ResultAny

// ── PathExistStat nil-safety tests ──

var extPathExistStatNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title: "HasError on nil receiver returns false",
		Func:  (*pathExistStatType).HasError,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsEmptyError on nil receiver returns true",
		Func:  (*pathExistStatType).IsEmptyError,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasFileInfo on nil receiver returns false",
		Func:  (*pathExistStatType).HasFileInfo,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalidFileInfo on nil receiver returns true",
		Func:  (*pathExistStatType).IsInvalidFileInfo,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsFile on nil receiver returns false",
		Func:  (*pathExistStatType).IsFile,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsDir on nil receiver returns false",
		Func:  (*pathExistStatType).IsDir,
		Expected: nilSafeResult{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsInvalid on nil receiver returns true",
		Func:  (*pathExistStatType).IsInvalid,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyIssues on nil receiver returns true",
		Func:  (*pathExistStatType).HasAnyIssues,
		Expected: nilSafeResult{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Dispose on nil receiver does not panic",
		Func:  (*pathExistStatType).Dispose,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "String on nil receiver returns empty",
		Func:  (*pathExistStatType).String,
		Expected: nilSafeResult{
			Value:    "",
			Panicked: false,
		},
	},
	{
		Title: "NotExistError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotExistError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "NotAFileError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotAFileError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "NotADirError on nil receiver returns nil",
		Func:  (*pathExistStatType).NotADirError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "MeaningFullError on nil receiver returns nil",
		Func:  (*pathExistStatType).MeaningFullError,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "LastModifiedDate on nil receiver returns nil",
		Func:  (*pathExistStatType).LastModifiedDate,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "FileMode on nil receiver returns nil",
		Func:  (*pathExistStatType).FileMode,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
	{
		Title: "Size on nil receiver returns nil",
		Func:  (*pathExistStatType).Size,
		Expected: nilSafeResult{
			Panicked: false,
		},
	},
}

// ── PathExistStat method tests ──

var extPathExistStatTestCases = []coretestcases.CaseV1{
	{
		Title: "GetPathExistStat on invalid path -- returns not exist",
		ArrangeInput: args.Map{
			"when": "invalid path",
			"path": "/nonexistent/path/that/does/not/exist/at/all",
		},
		ExpectedInput: args.Map{
			"isExist":  false,
			"isInvalid": true,
		},
	},
}

// ── chmodVerifier test cases ──

var extChmodVerifierTestCases = []coretestcases.CaseV1{
	{
		Title: "GetRwx9 returns 9 char string for standard filemode",
		ArrangeInput: args.Map{
			"when": "standard 0755 filemode",
		},
		ExpectedInput: args.Map{
			"rwx9Length": 9,
		},
	},
	{
		Title: "GetRwxFull returns 10 char string for standard filemode",
		ArrangeInput: args.Map{
			"when": "standard 0755 filemode",
		},
		ExpectedInput: args.Map{
			"rwxFullLength": 10,
		},
	},
}
