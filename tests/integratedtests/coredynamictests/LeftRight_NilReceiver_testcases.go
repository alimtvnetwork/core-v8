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

package coredynamictests

import (
	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/coretests/results"
)

// =============================================================================
// LeftRight nil receiver test cases
// (migrated from first element of CaseV1 slices in LeftRight_testcases.go)
// =============================================================================

var leftRightNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "IsEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasLeft false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasLeft,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "HasRight false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasRight,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "IsLeftEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsLeftEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "IsRightEmpty true on nil receiver",
		Func:  (*coredynamic.LeftRight).IsRightEmpty,
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "HasAnyItem false on nil receiver",
		Func:  (*coredynamic.LeftRight).HasAnyItem,
		Expected: results.ResultAny{
			Value:    "false",
			Panicked: false,
		},
	},
	{
		Title: "DeserializeLeft on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.DeserializeLeft() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "DeserializeRight on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.DeserializeRight() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "LeftToDynamic on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.LeftToDynamic() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "RightToDynamic on nil returns nil",
		Func: func(lr *coredynamic.LeftRight) bool {
			return lr.RightToDynamic() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
