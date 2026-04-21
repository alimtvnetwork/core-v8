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
	"strings"

	"github.com/alimtvnetwork/core-v8/errcore"
)

func IsErrorNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual error,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	var actualErrorMessage string

	if actual != nil {
		actualErrorMessage = actual.Error()
	}

	expectedString := expectationMessageDef.ExpectedString()

	if expectedString == "" && actualErrorMessage == "" {
		return true
	}

	return IsStrMsgNonWhiteSortedEqual(
		isPrintOnFail,
		actualErrorMessage,
		expectationMessageDef,
	)
}

func IsStrMsgNonWhiteSortedEqual(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	if expectationMessageDef.IsNonWhiteSort {
		return isStrMsgNonWhiteSortedEqualInternal(
			isPrintOnFail,
			actual,
			expectationMessageDef,
		)
	}

	trimActual := strings.TrimSpace(actual)
	trimExpected := expectationMessageDef.ExpectedStringTrim()
	isEqual := trimActual == trimExpected
	isFailed := !isEqual

	expectationMessageDef.ActualProcessed = trimActual
	expectationMessageDef.ExpectedProcessed = trimExpected
	expectationMessageDef.PrintIfFailed(
		isPrintOnFail,
		isFailed,
		actual,
	)

	return isEqual
}

func isStrMsgNonWhiteSortedEqualInternal(
	isPrintOnFail bool,
	actual string,
	expectationMessageDef *errcore.ExpectationMessageDef,
) bool {
	actualSortedDefault := GetAssert.SortedMessage(
		false,
		strings.TrimSpace(actual),
		commonJoiner,
	)

	expectedSortedDefault := GetAssert.SortedMessage(
		false,
		expectationMessageDef.ExpectedStringTrim(),
		commonJoiner,
	)

	isEqual := actualSortedDefault == expectedSortedDefault
	isFailed := !isEqual

	// Exception case for mutation, because test updates it
	expectationMessageDef.ActualProcessed = actualSortedDefault
	expectationMessageDef.ExpectedProcessed = expectedSortedDefault
	expectationMessageDef.PrintIfFailed(
		isPrintOnFail,
		isFailed,
		actual,
	)

	return isEqual
}
