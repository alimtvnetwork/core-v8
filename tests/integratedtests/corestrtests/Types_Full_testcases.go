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
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ── SimpleSlice (C15) ──

var srcC15SimpleSliceTestCase = coretestcases.CaseV1{
	Title: "SimpleSlice Add AddSplit AddIf Adds Append AppendFmt IsEmpty Length LastIndex HasIndex List Strings return correct -- various",
	ExpectedInput: args.Map{
		"addLen":      2,
		"splitLen":    3,
		"addIfLen":    1,
		"addsLen":     2,
		"appendLen":   2,
		"fmtLen":      1,
		"fmtEmptyLen": 0,
		"isEmpty":     true,
		"length":      1,
		"lastIndex":   1,
		"hasIdx0":     true,
		"hasIdx5":     false,
		"listLen":     1,
		"stringsLen":  1,
	},
}

var srcC15AnyToStringTestCase = coretestcases.CaseV1{
	Title: "AnyToString returns non-empty -- int 42",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

var srcC15AllIndLenSlicesTestCase = coretestcases.CaseV1{
	Title: "AllIndividualsLengthOfSimpleSlices returns correct -- 2 slices and nil",
	ExpectedInput: args.Map{
		"length": 3,
		"nilLen": 0,
	},
}

var srcC15TypesTestCase = coretestcases.CaseV1{
	Title: "ValidValues ValueStatus TextWithLineNumber LeftRight LeftMiddleRight KeyValuePair return correct -- struct fields",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

var srcC15KeyValueCollectionTestCase = coretestcases.CaseV1{
	Title: "KeyValueCollection Add IsEmpty return correct -- various",
	ExpectedInput: args.Map{
		"addLen":  1,
		"isEmpty": true,
	},
}

var srcC15HashsetsCollectionTestCase = coretestcases.CaseV1{
	Title: "HashsetsCollection IsEmpty Add return correct -- various",
	ExpectedInput: args.Map{
		"isEmpty": true,
		"addLen":  1,
	},
}

var srcC15DataModelTestCase = coretestcases.CaseV1{
	Title: "HashmapDataModel NewHashmapUsingDataModel NewHashmapsDataModelUsing return correct -- items",
	ExpectedInput: args.Map{
		"hmLen": 1,
		"dmLen": 1,
	},
}

var srcC15SimpleStringOnceTestCase = coretestcases.CaseV1{
	Title: "SimpleStringOnce returns non-empty -- test",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

var srcC15CollectionCreatorsTestCase = coretestcases.CaseV1{
	Title: "Collection creators LenCap LineUsingSep LineDefault StringsPlusCap CapStrings CloneStrings return correct -- various",
	ExpectedInput: args.Map{
		"lenCapLen":    5,
		"lineLen":      3,
		"lineDefGe1":   true,
		"strPlusLen":   1,
		"capStrLen":    1,
		"cloneDeep":    "a",
	},
}

var srcC15HashmapCreatorsTestCase = coretestcases.CaseV1{
	Title: "Hashmap creators KeyAnyValues KeyValuesCollection KeyValuesStrings MapWithCap return correct -- various",
	ExpectedInput: args.Map{
		"anyLen":    1,
		"colLen":    1,
		"strLen":    1,
		"mapLen":    1,
	},
}

var srcC15HashsetCreatorsTestCase = coretestcases.CaseV1{
	Title: "Hashset creators StringsOption Empty return correct -- various",
	ExpectedInput: args.Map{
		"optLen":  2,
		"isEmpty": true,
	},
}

var srcC15CocCreatorsTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection creators Empty StringsOfStrings SpreadStrings CloneStrings StringsOptions return correct -- various",
	ExpectedInput: args.Map{
		"emptyIsEmpty":  true,
		"sosLen":        2,
		"spreadLen":     1,
		"cloneLen":      1,
		"optionsLen":    1,
	},
}

var srcC15CocJsonTestCase = coretestcases.CaseV1{
	Title: "CollectionsOfCollection MarshalJSON UnmarshalJSON Json ParseInjectUsingJson JsonParseSelfInject AddEmpty return correct -- various",
	ExpectedInput: args.Map{
		"noPanic":     true,
		"addEmptyLen": 0,
	},
}
