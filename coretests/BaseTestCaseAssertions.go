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
	"errors"
	"fmt"
	"log/slog"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// =============================================================================
// Assertion Methods
// =============================================================================

// ShouldBe
//
// Disabled testcases will not be executed.
func (it *BaseTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual any,
) {
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.ShouldBeExplicit(
		true,
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected(),
	)
}

func (it *BaseTestCase) noPrintAssert(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual any,
) {
	toTile := it.FormTitle(caseIndex)

	it.SetActual(actual)

	convey.Convey(
		toTile, t, func() {
			convey.SoMsg(
				toTile,
				actual,
				assert,
				it.ExpectedInput,
			)
		},
	)
}

func (it *BaseTestCase) ShouldBeExplicit(
	isValidateType bool,
	caseIndex int,
	t *testing.T,
	title string,
	actual any,
	assert convey.Assertion,
	expected any,
) {
	if it.IsEnable.IsFalse() {
		it.noPrintAssert(caseIndex, t, assert, actual)

		return
	}

	it.SetActual(actual)
	headerTitle := it.CustomTitle(caseIndex, title)
	actualLines := GetAssert.ToStrings(actual)
	expectedLines := GetAssert.ToStrings(expected)
	compare := assert(actualLines, expectedLines)
	isFailed := compare != ""

	convey.Convey(
		headerTitle, t, func() {
			if isFailed {
				toString := it.LinesString(caseIndex)

				slog.Warn("test case mismatch", "caseIndex", caseIndex, "details", toString)
			}

			convey.SoMsg(
				headerTitle,
				actualLines,
				assert,
				expectedLines,
			)
		},
	)

	isSkipTypeValidation := !isValidateType

	if isSkipTypeValidation {
		return
	}

	it.TypeShouldMatch(t, caseIndex, title)
}

func (it *BaseTestCase) TypeShouldMatch(
	t *testing.T,
	caseIndex int,
	title string,
) {
	err := it.TypeValidationError()

	if err == nil {
		return
	}

	errHeader := fmt.Sprintf(
		"%d : %s - type verification failed",
		caseIndex,
		title,
	)

	var finalError error

	if err != nil {
		finalError = errors.New(errHeader + err.Error())
	}

	convey.Convey(
		errHeader, t, func() {
			convey.So(
				finalError,
				convey.ShouldBeNil,
			)
		},
	)
}
