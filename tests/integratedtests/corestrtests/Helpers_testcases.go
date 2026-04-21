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

package corestrtests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ── CloneSlice ──

var cloneSliceTestCases = []coretestcases.CaseV1{
	{
		Title:         "CloneSlice returns empty -- nil input",
		ExpectedInput: args.Map{"len": 0},
	},
	{
		Title:         "CloneSlice returns cloned items -- non-empty input",
		ExpectedInput: args.Map{
			"len": 2,
			"first": "a",
		},
	},
}

// ── CloneSliceIf ──

var cloneSliceIfTestCases = []coretestcases.CaseV1{
	{
		Title:         "CloneSliceIf returns empty -- empty variadic",
		ExpectedInput: args.Map{"len": 0},
	},
	{
		Title:         "CloneSliceIf returns original -- isClone false",
		ExpectedInput: args.Map{"len": 2},
	},
	{
		Title:         "CloneSliceIf returns clone -- isClone true",
		ExpectedInput: args.Map{"len": 2},
	},
}

// ── AnyToString ──

var anyToStringTestCases = []coretestcases.CaseV1{
	{
		Title:         "AnyToString returns empty -- empty input",
		ExpectedInput: args.Map{"isEmpty": true},
	},
	{
		Title:         "AnyToString returns non-empty -- with field name",
		ExpectedInput: args.Map{"isEmpty": false},
	},
	{
		Title:         "AnyToString returns non-empty -- without field name",
		ExpectedInput: args.Map{"isEmpty": false},
	},
	{
		Title:         "AnyToString returns non-empty -- pointer input",
		ExpectedInput: args.Map{"isEmpty": false},
	},
}

// ── AllIndividualStringsOfStringsLength ──

var allIndividualStringsOfStringsLengthTestCases = []coretestcases.CaseV1{
	{
		Title:         "AllIndividualStringsOfStringsLength returns 0 -- nil",
		ExpectedInput: args.Map{"len": 0},
	},
	{
		Title:         "AllIndividualStringsOfStringsLength returns 0 -- empty",
		ExpectedInput: args.Map{"len": 0},
	},
	{
		Title:         "AllIndividualStringsOfStringsLength returns 3 -- with items",
		ExpectedInput: args.Map{"len": 3},
	},
}

// ── AllIndividualsLengthOfSimpleSlices ──

var allIndividualsLengthOfSimpleSlicesTestCases = []coretestcases.CaseV1{
	{
		Title:         "AllIndividualsLengthOfSimpleSlices returns 0 -- nil",
		ExpectedInput: args.Map{"len": 0},
	},
	{
		Title:         "AllIndividualsLengthOfSimpleSlices returns 3 -- with items",
		ExpectedInput: args.Map{"len": 3},
	},
}

// ── utils ──

var utilsWrapTestCases = []coretestcases.CaseV1{
	{
		Title:         "WrapDouble returns quoted -- simple input",
		ExpectedInput: args.Map{"result": `"a"`},
	},
	{
		Title:         "WrapSingle returns quoted -- simple input",
		ExpectedInput: args.Map{"result": "'a'"},
	},
	{
		Title:         "WrapTilda returns backticked -- simple input",
		ExpectedInput: args.Map{"result": "`a`"},
	},
	{
		Title:         "WrapDoubleIfMissing returns quoted -- empty input",
		ExpectedInput: args.Map{"result": `""`},
	},
	{
		Title:         "WrapDoubleIfMissing returns same -- already wrapped",
		ExpectedInput: args.Map{"result": `"a"`},
	},
	{
		Title:         "WrapDoubleIfMissing returns quoted -- not wrapped",
		ExpectedInput: args.Map{"result": `"a"`},
	},
	{
		Title:         "WrapSingleIfMissing returns quoted -- empty input",
		ExpectedInput: args.Map{"result": "''"},
	},
	{
		Title:         "WrapSingleIfMissing returns same -- already wrapped",
		ExpectedInput: args.Map{"result": "'a'"},
	},
	{
		Title:         "WrapSingleIfMissing returns quoted -- not wrapped",
		ExpectedInput: args.Map{"result": "'a'"},
	},
}

// ── Empty creators ──

var emptyCreatorTestCases = []coretestcases.CaseV1{
	{
		Title:         "Empty creators return non-nil -- all types",
		ExpectedInput: args.Map{"allNonNil": true},
	},
}

// ── DataModels ──

var dataModelTestCases = []coretestcases.CaseV1{
	{
		Title:         "CharCollectionDataModel round-trips -- valid input",
		ExpectedInput: args.Map{"nonNil": true},
	},
	{
		Title:         "CharHashsetDataModel round-trips -- valid input",
		ExpectedInput: args.Map{"nonNil": true},
	},
	{
		Title:         "HashmapDataModel round-trips -- valid input",
		ExpectedInput: args.Map{
			"nonNil": true,
			"nonEmpty": true,
		},
	},
	{
		Title:         "HashsetDataModel round-trips -- valid input",
		ExpectedInput: args.Map{
			"nonNil": true,
			"nonEmpty": true,
		},
	},
	{
		Title:         "HashsetsCollectionDataModel round-trips -- valid input",
		ExpectedInput: args.Map{"nonNil": true},
	},
}

// ── SimpleStringOnceModel ──

var simpleStringOnceModelTestCases = []coretestcases.CaseV1{
	{
		Title:         "SimpleStringOnceModel stores value -- valid input",
		ExpectedInput: args.Map{"value": "hello"},
	},
}

// ── CollectionsOfCollectionModel ──

var collectionsOfCollectionModelTestCases = []coretestcases.CaseV1{
	{
		Title:         "CollectionsOfCollectionModel stores items -- valid input",
		ExpectedInput: args.Map{"nonNil": true},
	},
}
