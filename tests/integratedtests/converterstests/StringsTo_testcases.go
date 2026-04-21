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

package converterstests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// =============================================================================
// IntegersWithDefaults
// =============================================================================

var integersWithDefaultsTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegersWithDefaults returns all converted -- valid integer strings",
		ArrangeInput: args.Map{
			"when":       "given all valid integer strings",
			"input":      []string{"1", "2", "3"},
			"defaultInt": -1,
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           1,
			"val1":           2,
			"val2":           3,
			"hadDefaultUsed": false,
		},
	},
	{
		Title: "IntegersWithDefaults returns default for invalid -- mix of valid and invalid strings",
		ArrangeInput: args.Map{
			"when":       "given mix of valid and invalid strings",
			"input":      []string{"10", "abc", "20"},
			"defaultInt": -1,
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           10,
			"val1":           -1,
			"val2":           20,
			"hadDefaultUsed": true,
		},
	},
	{
		Title: "IntegersWithDefaults returns empty -- empty input",
		ArrangeInput: args.Map{
			"when":       "given empty input",
			"input":      []string{},
			"defaultInt": 0,
		},
		ExpectedInput: args.Map{
			"count":          0,
			"hadDefaultUsed": false,
		},
	},
	{
		Title: "IntegersWithDefaults returns all defaults -- all non-numeric strings",
		ArrangeInput: args.Map{
			"when":       "given all non-numeric strings",
			"input":      []string{"x", "y", "z"},
			"defaultInt": 99,
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           99,
			"val1":           99,
			"val2":           99,
			"hadDefaultUsed": true,
		},
	},
	{
		Title: "IntegersWithDefaults returns correct values -- negative numbers",
		ArrangeInput: args.Map{
			"when":       "given negative number strings",
			"input":      []string{"-5", "0", "5"},
			"defaultInt": 0,
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           -5,
			"val1":           0,
			"val2":           5,
			"hadDefaultUsed": false,
		},
	},
}

// =============================================================================
// BytesWithDefaults
// =============================================================================

var bytesWithDefaultsTestCases = []coretestcases.CaseV1{
	{
		Title: "BytesWithDefaults returns all converted -- valid byte strings",
		ArrangeInput: args.Map{
			"when":        "given valid byte strings",
			"input":       []string{"0", "127", "255"},
			"defaultByte": byte(0),
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           0,
			"val1":           127,
			"val2":           255,
			"hadDefaultUsed": false,
		},
	},
	{
		Title: "BytesWithDefaults returns default -- out-of-range value > 255",
		ArrangeInput: args.Map{
			"when":        "given value > 255",
			"input":       []string{"100", "256", "50"},
			"defaultByte": byte(42),
		},
		ExpectedInput: args.Map{
			"count":          3,
			"val0":           100,
			"val1":           42,
			"val2":           50,
			"hadDefaultUsed": true,
		},
	},
	{
		Title: "BytesWithDefaults returns default -- negative value",
		ArrangeInput: args.Map{
			"when":        "given negative value",
			"input":       []string{"-1", "10"},
			"defaultByte": byte(0),
		},
		ExpectedInput: args.Map{
			"count":          2,
			"val0":           0,
			"val1":           10,
			"hadDefaultUsed": true,
		},
	},
	{
		Title: "BytesWithDefaults returns default -- non-numeric string",
		ArrangeInput: args.Map{
			"when":        "given non-numeric string",
			"input":       []string{"abc"},
			"defaultByte": byte(99),
		},
		ExpectedInput: args.Map{
			"count":          1,
			"val0":           99,
			"hadDefaultUsed": true,
		},
	},
	{
		Title: "BytesWithDefaults returns empty -- empty input",
		ArrangeInput: args.Map{
			"when":        "given empty input",
			"input":       []string{},
			"defaultByte": byte(0),
		},
		ExpectedInput: args.Map{
			"count":          0,
			"hadDefaultUsed": false,
		},
	},
}

// =============================================================================
// CloneIf
// =============================================================================

var cloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneIf returns independent clone -- isClone true",
		ArrangeInput: args.Map{
			"when":    "given isClone true",
			"input":   []string{"a", "b", "c"},
			"isClone": true,
		},
		ExpectedInput: args.Map{
			"count":         3,
			"item0":         "a",
			"item1":         "b",
			"item2":         "c",
			"isIndependent": true,
		},
	},
	{
		Title: "CloneIf returns same slice -- isClone false",
		ArrangeInput: args.Map{
			"when":    "given isClone false",
			"input":   []string{"x", "y"},
			"isClone": false,
		},
		ExpectedInput: args.Map{
			"count":         2,
			"item0":         "x",
			"item1":         "y",
			"isIndependent": false,
		},
	},
	{
		Title: "CloneIf returns empty -- empty input regardless of isClone",
		ArrangeInput: args.Map{
			"when":    "given empty input with isClone true",
			"input":   []string{},
			"isClone": true,
		},
		ExpectedInput: args.Map{
			"count":         0,
			"isIndependent": false,
		},
	},
}

// =============================================================================
// PtrOfPtrToPtrStrings
// =============================================================================

var ptrOfPtrToPtrStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOfPtrToPtrStrings returns converted strings -- valid pointer string array",
		ArrangeInput: args.Map{
			"when":  "given valid pointer string array",
			"input": []string{"hello", "world"},
		},
		ExpectedInput: args.Map{
			"count": 2,
			"item0": "hello",
			"item1": "world",
		},
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty for nil -- nil entry in array",
		ArrangeInput: args.Map{
			"when":   "given array with nil entry",
			"input":  []string{"hello"},
			"hasNil": true,
			"nilIdx": 1,
		},
		ExpectedInput: args.Map{
			"count": 2,
			"item0": "hello",
			"item1": "",
		},
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty -- nil outer pointer",
		ArrangeInput: args.Map{
			"when":  "given nil outer pointer",
			"isNil": true,
		},
		ExpectedInput: "0",
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty -- nil inner pointer",
		ArrangeInput: args.Map{
			"when":       "given nil inner pointer",
			"isNilInner": true,
		},
		ExpectedInput: "0",
	},
}

// =============================================================================
// PtrOfPtrToMapStringBool
// =============================================================================

var ptrOfPtrToMapStringBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOfPtrToMapStringBool returns converted entries -- valid pointer string array",
		ArrangeInput: args.Map{
			"when":  "given valid pointer string array",
			"input": []string{"key1", "key2"},
		},
		ExpectedInput: args.Map{
			"count":   2,
			"haskey1": true,
			"haskey2": true,
		},
	},
	{
		Title: "PtrOfPtrToMapStringBool returns without nil -- nil entry skipped",
		ArrangeInput: args.Map{
			"when":   "given array with nil entry",
			"input":  []string{"key1"},
			"hasNil": true,
		},
		ExpectedInput: args.Map{
			"count":   1,
			"haskey1": true,
		},
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty -- nil input",
		ArrangeInput: args.Map{
			"when":  "given nil outer pointer",
			"isNil": true,
		},
		ExpectedInput: "0",
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty -- empty array",
		ArrangeInput: args.Map{
			"when":  "given empty array",
			"input": []string{},
		},
		ExpectedInput: "0",
	},
}
