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

package corevalidatortests

import (
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

// ==========================================================================
// Condition.IsSplitByWhitespace
// ==========================================================================

var conditionAllFalseTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns false -- all flags false",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionUniqueWordOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsUniqueWordOnly enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionNonEmptyWhitespaceTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsNonEmptyWhitespace enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionSortBySpaceTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns true -- IsSortStringsBySpace enabled",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionTrimOnlyTestCase = coretestcases.CaseV1{
	Title:         "IsSplitByWhitespace returns false -- IsTrimCompare only",
	ExpectedInput: args.Map{"isSplit": false},
}

// ==========================================================================
// Preset Conditions
// ==========================================================================

var conditionDisabledTestCase = coretestcases.CaseV1{
	Title:         "DefaultDisabled returns isSplit false -- preset disabled",
	ExpectedInput: args.Map{"isSplit": false},
}

var conditionTrimTestCase = coretestcases.CaseV1{
	Title: "DefaultTrim returns isSplit false, isTrimCompare true -- preset trim",
	ExpectedInput: args.Map{
		"isSplit":       false,
		"isTrimCompare": true,
	},
}

var conditionSortTrimTestCase = coretestcases.CaseV1{
	Title:         "DefaultSortTrim returns isSplit true -- preset sort-trim",
	ExpectedInput: args.Map{"isSplit": true},
}

var conditionUniqueWordsTestCase = coretestcases.CaseV1{
	Title: "DefaultUniqueWords returns isSplit true, isUniqueWordOnly true -- preset unique-words",
	ExpectedInput: args.Map{
		"isSplit":          true,
		"isUniqueWordOnly": true,
	},
}
