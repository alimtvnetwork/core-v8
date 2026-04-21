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

package codestacktests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

var traceTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace basic properties -- valid trace returns correct fields",
		ArrangeInput: args.Map{
			"when":        "valid trace created",
			"packageName": "mypackage",
			"methodName":  "MyMethod",
			"pkgMethod":   "mypackage.MyMethod",
			"filePath":    "/src/mypackage/file.go",
			"line":        42,
		},
		ExpectedInput: args.Map{
			"packageName": "mypackage",
			"methodName":  "MyMethod",
			"pkgMethod":   "mypackage.MyMethod",
			"filePath":    "/src/mypackage/file.go",
			"lineNumber":  42,
			"isNil":       false,
			"isNotNil":    true,
			"hasIssues":   false,
		},
	},
}

var traceNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace nil -- nil trace returns correct defaults",
		ArrangeInput: args.Map{
			"when": "nil trace",
		},
		ExpectedInput: args.Map{
			"isNil":    true,
			"isNotNil": false,
			"string":   "",
		},
	},
}

var traceDisposeTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace Dispose -- clears all fields",
		ArrangeInput: args.Map{
			"when":        "dispose called on valid trace",
			"packageName": "pkg",
			"methodName":  "Method",
			"pkgMethod":   "pkg.Method",
			"filePath":    "/file.go",
			"line":        10,
		},
		ExpectedInput: args.Map{
			"packageName": "",
			"methodName":  "",
			"pkgMethod":   "",
			"filePath":    "",
			"lineNumber":  0,
			"isOkay":      false,
		},
	},
}

var traceCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Trace Clone -- creates independent copy",
		ArrangeInput: args.Map{
			"when":        "clone valid trace",
			"packageName": "clonepkg",
			"methodName":  "CloneMethod",
			"pkgMethod":   "clonepkg.CloneMethod",
			"filePath":    "/clone/file.go",
			"line":        99,
		},
		ExpectedInput: args.Map{
			"packageName": "clonepkg",
			"methodName":  "CloneMethod",
			"filePath":    "/clone/file.go",
			"lineNumber":  99,
		},
	},
}

var fileWithLineStringMethodTestCases = []coretestcases.CaseV1{
	{
		Title: "FileWithLine String method -- returns formatted path:line",
		ArrangeInput: args.Map{
			"when": "file with line created",
			"file": "/src/app.go",
			"line": 100,
		},
		ExpectedInput: args.Map{
			"isNil":    false,
			"isNotNil": true,
			"hasLine":  true,
		},
	},
}
