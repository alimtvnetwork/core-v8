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
	"github.com/alimtvnetwork/core/codestack"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

type fileWithLineType = codestack.FileWithLine
type traceType = codestack.Trace
type traceCollectionType = codestack.TraceCollection
type nilResult = results.ResultAny

// ── FileWithLine nil-safety ──

var coverageFileWithLineNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title:    "FullFilePath on nil returns empty",
		Func:     (*fileWithLineType).FullFilePath,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "LineNumber on nil returns 0",
		Func:     (*fileWithLineType).LineNumber,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "IsNil on nil returns true",
		Func:     (*fileWithLineType).IsNil,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "IsNotNil on nil returns false",
		Func:     (*fileWithLineType).IsNotNil,
		Expected: nilResult{Value: "false", Panicked: false},
	},
	{
		Title:    "String on nil returns empty",
		Func:     (*fileWithLineType).String,
		Expected: nilResult{Value: "", Panicked: false},
	},
	{
		Title:    "JsonModelAny on nil panics",
		Func:     (*fileWithLineType).JsonModelAny,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "JsonString on nil panics",
		Func:     (*fileWithLineType).JsonString,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "AsFileLiner on nil returns nil",
		Func:     (*fileWithLineType).AsFileLiner,
		Expected: nilResult{Panicked: false},
	},
}

// ── Trace nil-safety ──

var coverageTraceNilSafeCases = []coretestcases.CaseNilSafe{
	{
		Title:    "Trace.IsNil on nil returns true",
		Func:     (*traceType).IsNil,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "Trace.IsNotNil on nil returns false",
		Func:     (*traceType).IsNotNil,
		Expected: nilResult{Value: "false", Panicked: false},
	},
	{
		Title:    "Trace.HasIssues on nil returns true",
		Func:     (*traceType).HasIssues,
		Expected: nilResult{Value: "true", Panicked: false},
	},
	{
		Title:    "Trace.String on nil returns empty",
		Func:     (*traceType).String,
		Expected: nilResult{Value: "", Panicked: false},
	},
	{
		Title:    "Trace.Dispose on nil does not panic",
		Func:     (*traceType).Dispose,
		Expected: nilResult{Panicked: false},
	},
	{
		Title:    "Trace.ClonePtr on nil returns nil",
		Func:     (*traceType).ClonePtr,
		Expected: nilResult{Panicked: false},
	},
	{
		Title:    "Trace.JsonModelAny on nil panics",
		Func:     (*traceType).JsonModelAny,
		Expected: nilResult{Panicked: true},
	},
	{
		Title:    "Trace.AsFileLiner on nil returns nil",
		Func:     (*traceType).AsFileLiner,
		Expected: nilResult{Panicked: false},
	},
}
