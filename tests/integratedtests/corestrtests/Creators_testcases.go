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

// Branch: newCollectionCreator factory methods
var srcC05CollectionCreatorTestCase = coretestcases.CaseV1{
	Title: "Collection creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty":      true,
		"capHasCap":         true,
		"stringsLen":        2,
		"createLen":         1,
		"cloneLen":          1,
		"optCloneLen":       1,
		"optNoCloneLen":     1,
		"optNilEmpty":       true,
		"sepLen":            2,
		"plusCapLen":         1,
		"lenCapLen":         2,
	},
}

// Branch: newSimpleSliceCreator factory methods
var srcC05SimpleSliceCreatorTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty":    true,
		"cap10Len":        0,
		"linesLen":        2,
		"noPanic":         true,
	},
}

// Branch: newSimpleStringOnceCreator factory methods
var srcC05SimpleStringOnceCreatorTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"initValue":         "hello",
		"initIsInitialized": true,
		"initPtrValue":      "hello",
		"uninitIsInit":      false,
	},
}

// Branch: newHashsetCreator factory methods
var srcC05HashsetCreatorTestCase = coretestcases.CaseV1{
	Title: "Hashset creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"stringsLen":   2,
		"noPanic":      true,
	},
}

// Branch: newHashmapCreator factory methods
var srcC05HashmapCreatorTestCase = coretestcases.CaseV1{
	Title: "Hashmap creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty":   true,
		"usingMapEmpty":  false,
		"noPanic":        true,
	},
}

// Branch: newLinkedListCreator factory methods
var srcC05LinkedListCreatorTestCase = coretestcases.CaseV1{
	Title: "LinkedList creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"createLen":  0,
		"stringsLen": 2,
		"noPanic":    true,
	},
}

// Branch: newLinkedCollectionCreator factory methods
var srcC05LinkedCollectionCreatorTestCase = coretestcases.CaseV1{
	Title: "LinkedCollection creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"createLen":  0,
		"stringsLen": 1,
		"noPanic":    true,
	},
}

// Branch: newKeyValuesCreator factory methods
var srcC05KeyValuesCreatorTestCase = coretestcases.CaseV1{
	Title: "KeyValues creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"usingMapLen":  1,
		"noPanic":      true,
	},
}

// Branch: newCollectionsOfCollectionCreator factory methods
var srcC05CollOfCollCreatorTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty": true,
		"noPanic":      true,
	},
}

// Branch: newHashsetsCollectionCreator factory methods
var srcC05HashsetsCollCreatorTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection creator factories return correct state -- various inputs",
	ExpectedInput: args.Map{
		"emptyIsEmpty":    true,
		"usingPtrsEmpty":  false,
		"noPanic":         true,
	},
}

// Branch: CharCollectionMap and CharHashsetMap creators
var srcC05CharMapCreatorsTestCase = coretestcases.CaseV1{
	Title: "CharCollectionMap and CharHashsetMap creators return correct state -- various inputs",
	ExpectedInput: args.Map{
		"ccmEmptyIsEmpty": true,
		"noPanic":         true,
	},
}
