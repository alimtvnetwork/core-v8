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

package coretests

import (
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/errcore"
)

type ComparingInstruction struct {
	FunName,
	Header,
	TestCaseName,
	MatchingAsEqualExpectation string
	ComparingItems                   []Compare
	HasWhitespace, IsMatchingAsEqual bool
	actualHashset                    *corestr.Hashset
	actual                           string
}

func (it *ComparingInstruction) Actual() string {
	return it.actual
}

func (it *ComparingInstruction) SetActual(actual string) {
	it.actual = actual
	it.actualHashset = nil
}

func (it *ComparingInstruction) ActualHashset() *corestr.Hashset {
	if it.actualHashset != nil {
		return it.actualHashset
	}

	whitespaceRemovedSplits := GetAssert.SortedArray(
		false,
		true,
		it.Actual(),
	)

	it.actualHashset = corestr.New.Hashset.Strings(whitespaceRemovedSplits)

	return it.actualHashset
}

func (it *ComparingInstruction) IsMatch(
	caseIndexPlusIsPrint *CaseIndexPlusIsPrint,
) bool {
	isMatchesEqual := it.isMatchingEqual(caseIndexPlusIsPrint)

	for i, item := range it.ComparingItems {
		isMatchesEqual = item.IsMatch(
			caseIndexPlusIsPrint.IsPrint,
			i,
			it,
		) &&
			isMatchesEqual
	}

	return isMatchesEqual
}

func (it *ComparingInstruction) isMatchingEqual(caseIndexPlusIsPrint *CaseIndexPlusIsPrint) bool {
	if !it.IsMatchingAsEqual {
		return true
	}

	expectation := &errcore.ExpectationMessageDef{
		CaseIndex:      caseIndexPlusIsPrint.CaseIndex,
		FuncName:       it.FunName,
		TestCaseName:   it.TestCaseName,
		When:           it.Header,
		Expected:       it.MatchingAsEqualExpectation,
		IsNonWhiteSort: it.HasWhitespace,
	}

	return IsStrMsgNonWhiteSortedEqual(
		caseIndexPlusIsPrint.IsPrint,
		it.actual,
		expectation,
	)
}
