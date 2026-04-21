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

package coregenerictests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// MapCollection
// ==========================================================================

var mapCollectionIntToStringTestCase = coretestcases.CaseV1{
	Title:         "MapCollection int to string",
	ExpectedInput: args.Map{
		"length": 3,
		"first": "v1",
		"last": "v3",
	},
}

var mapCollectionNilSourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection nil source",
	ExpectedInput: args.Map{"isEmpty": true},
}

var mapCollectionEmptySourceTestCase = coretestcases.CaseV1{
	Title:         "MapCollection empty source",
	ExpectedInput: args.Map{"isEmpty": true},
}

// ==========================================================================
// FlatMapCollection
// ==========================================================================

var flatMapCollectionFlattensTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection flattens",
	ExpectedInput: args.Map{"length": 6},
}

var flatMapCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "FlatMapCollection nil",
	ExpectedInput: args.Map{"isEmpty": true},
}

// ==========================================================================
// ReduceCollection
// ==========================================================================

var reduceCollectionSumTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection sum",
	ExpectedInput: args.Map{"result": 10},
}

var reduceCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection nil returns initial",
	ExpectedInput: args.Map{"result": 99},
}

var reduceCollectionConcatTestCase = coretestcases.CaseV1{
	Title:         "ReduceCollection string concat",
	ExpectedInput: args.Map{"result": "abc"},
}

// ==========================================================================
// GroupByCollection
// ==========================================================================

var groupByCollectionGroupsTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection groups",
	ExpectedInput: args.Map{
		"groupCount": 2,
		"evenCount": 3,
		"oddCount": 3,
	},
}

var groupByCollectionNilTestCase = coretestcases.CaseV1{
	Title:         "GroupByCollection nil",
	ExpectedInput: args.Map{"groupCount": 0},
}

// ==========================================================================
// ContainsFunc
// ==========================================================================

var containsFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc found",
	ExpectedInput: args.Map{"result": true},
}

var containsFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc not found",
	ExpectedInput: args.Map{"result": false},
}

var containsFuncNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsFunc nil",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// ContainsItem
// ==========================================================================

var containsItemFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem found",
	ExpectedInput: args.Map{"result": true},
}

var containsItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem not found",
	ExpectedInput: args.Map{"result": false},
}

var containsItemNilTestCase = coretestcases.CaseV1{
	Title:         "ContainsItem nil",
	ExpectedInput: args.Map{"result": false},
}

// ==========================================================================
// IndexOfFunc
// ==========================================================================

var indexOfFuncFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc found",
	ExpectedInput: args.Map{"index": 1},
}

var indexOfFuncNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc not found",
	ExpectedInput: args.Map{"index": -1},
}

var indexOfFuncNilTestCase = coretestcases.CaseV1{
	Title:         "IndexOfFunc nil",
	ExpectedInput: args.Map{"index": -1},
}

// ==========================================================================
// IndexOfItem
// ==========================================================================

var indexOfItemFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem found",
	ExpectedInput: args.Map{"index": 2},
}

var indexOfItemNotFoundTestCase = coretestcases.CaseV1{
	Title:         "IndexOfItem not found",
	ExpectedInput: args.Map{"index": -1},
}

// ==========================================================================
// Distinct
// ==========================================================================

var distinctRemovesDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct removes duplicates",
	ExpectedInput: args.Map{"length": 3},
}

var distinctNilTestCase = coretestcases.CaseV1{
	Title:         "Distinct nil",
	ExpectedInput: args.Map{"isEmpty": true},
}

var distinctNoDuplicatesTestCase = coretestcases.CaseV1{
	Title:         "Distinct no duplicates",
	ExpectedInput: args.Map{"length": 3},
}

// ==========================================================================
// MapSimpleSlice
// ==========================================================================

var mapSimpleSliceTransformsTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice transforms",
	ExpectedInput: args.Map{"length": 3},
}

var mapSimpleSliceNilTestCase = coretestcases.CaseV1{
	Title:         "MapSimpleSlice nil",
	ExpectedInput: args.Map{"isEmpty": true},
}
