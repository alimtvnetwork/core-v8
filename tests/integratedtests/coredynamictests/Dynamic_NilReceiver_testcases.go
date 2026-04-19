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
// Dynamic nil receiver test cases
// (migrated from standalone CaseV1 variables in Dynamic_testcases.go)
// =============================================================================

var dynamicNilSafeTestCases = []coretestcases.CaseNilSafe{
	{
		Title: "ClonePtr returns nil on nil receiver",
		Func: func(d *coredynamic.Dynamic) bool {
			return d.ClonePtr() == nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Bytes returns nil,false on nil receiver",
		Func: func(d *coredynamic.Dynamic) bool {
			raw, ok := d.Bytes()
			return raw == nil && !ok
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ValueNullErr returns error on nil receiver",
		Func: func(d *coredynamic.Dynamic) error {
			return d.ValueNullErr()
		},
		Expected: results.ResultAny{
			Panicked: false,
		},
	},
	{
		Title: "ValueMarshal on nil returns nil bytes and error",
		Func: func(d *coredynamic.Dynamic) bool {
			bytes, err := d.ValueMarshal()
			return bytes == nil && err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "ReflectSetTo on nil returns error",
		Func: func(d *coredynamic.Dynamic) bool {
			var target string
			err := d.ReflectSetTo(&target)
			return err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "Deserialize on nil returns invalid and error",
		Func: func(d *coredynamic.Dynamic) bool {
			result, err := d.Deserialize([]byte(`{}`))
			return result != nil && err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
	{
		Title: "UnmarshalJSON on nil returns error",
		Func: func(d *coredynamic.Dynamic) bool {
			err := d.UnmarshalJSON([]byte(`{}`))
			return err != nil
		},
		Expected: results.ResultAny{
			Value:    "true",
			Panicked: false,
		},
	},
}
