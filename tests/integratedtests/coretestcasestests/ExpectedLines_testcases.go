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

package coretestcasestests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var expectedLinesTestCases = []coretestcases.CaseV1{
	{
		Title:         "ExpectedLines converts int to string",
		ArrangeInput:  args.Map{"inputType": "int"},
		ExpectedInput: 42,
	},
	{
		Title:         "ExpectedLines converts bool true",
		ArrangeInput:  args.Map{"inputType": "boolTrue"},
		ExpectedInput: true,
	},
	{
		Title:         "ExpectedLines converts bool false",
		ArrangeInput:  args.Map{"inputType": "boolFalse"},
		ExpectedInput: false,
	},
	{
		Title:         "ExpectedLines converts []int slice",
		ArrangeInput:  args.Map{"inputType": "intSlice"},
		ExpectedInput: []int{10, 20, 30},
	},
	{
		Title:         "ExpectedLines converts []bool slice",
		ArrangeInput:  args.Map{"inputType": "boolSlice"},
		ExpectedInput: []bool{true, false, true},
	},
	{
		Title:         "ExpectedLines wraps string into slice",
		ArrangeInput:  args.Map{"inputType": "string"},
		ExpectedInput: "hello",
	},
	{
		Title:         "ExpectedLines returns []string as-is",
		ArrangeInput:  args.Map{"inputType": "stringSlice"},
		ExpectedInput: []string{"a", "b", "c"},
	},
	{
		Title:        "ExpectedLines converts map[string]int sorted",
		ArrangeInput: args.Map{"inputType": "mapStringInt"},
		ExpectedInput: map[string]int{
			"age":   30,
			"count": 5,
		},
	},
}

// expectedLinesExpectedOutputs holds the expected output for each test case
// as args.Map with lineCount + indexed line keys.
var expectedLinesExpectedOutputs = []args.Map{
	{"lineCount": "1", "line0": "42"},
	{"lineCount": "1", "line0": "true"},
	{"lineCount": "1", "line0": "false"},
	{"lineCount": "3", "line0": "10", "line1": "20", "line2": "30"},
	{"lineCount": "3", "line0": "true", "line1": "false", "line2": "true"},
	{"lineCount": "1", "line0": "hello"},
	{"lineCount": "3", "line0": "a", "line1": "b", "line2": "c"},
	{"lineCount": "2", "line0": "age : 30", "line1": "count : 5"},
}
