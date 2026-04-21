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
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/corevalidator"
)

// ==========================================
// Condition.IsSplitByWhitespace
// ==========================================

func Test_Condition_IsSplitByWhitespace_AllFalse_FromCondition(t *testing.T) {
	// Arrange
	tc := conditionAllFalseTestCase
	c := corevalidator.Condition{}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_UniqueWordOnly(t *testing.T) {
	// Arrange
	tc := conditionUniqueWordOnlyTestCase
	c := corevalidator.Condition{IsUniqueWordOnly: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_NonEmptyWhitespace(t *testing.T) {
	// Arrange
	tc := conditionNonEmptyWhitespaceTestCase
	c := corevalidator.Condition{IsNonEmptyWhitespace: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_SortBySpace(t *testing.T) {
	// Arrange
	tc := conditionSortBySpaceTestCase
	c := corevalidator.Condition{IsSortStringsBySpace: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_Condition_IsSplitByWhitespace_TrimOnlyNotEnough(t *testing.T) {
	// Arrange
	tc := conditionTrimOnlyTestCase
	c := corevalidator.Condition{IsTrimCompare: true}

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// Preset Conditions
// ==========================================

func Test_DefaultDisabledCondition_NoSplit(t *testing.T) {
	// Arrange
	tc := conditionDisabledTestCase
	c := corevalidator.DefaultDisabledCoreCondition

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultTrimCondition_NoSplit(t *testing.T) {
	// Arrange
	tc := conditionTrimTestCase
	c := corevalidator.DefaultTrimCoreCondition

	// Act
	actual := args.Map{
		"isSplit":       c.IsSplitByWhitespace(),
		"isTrimCompare": c.IsTrimCompare,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultSortTrimCondition_Split(t *testing.T) {
	// Arrange
	tc := conditionSortTrimTestCase
	c := corevalidator.DefaultSortTrimCoreCondition

	// Act
	actual := args.Map{"isSplit": c.IsSplitByWhitespace()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_DefaultUniqueWordsCondition_Split(t *testing.T) {
	// Arrange
	tc := conditionUniqueWordsTestCase
	c := corevalidator.DefaultUniqueWordsCoreCondition

	// Act
	actual := args.Map{
		"isSplit":          c.IsSplitByWhitespace(),
		"isUniqueWordOnly": c.IsUniqueWordOnly,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
