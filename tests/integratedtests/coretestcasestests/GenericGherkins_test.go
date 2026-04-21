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

package coretestcasestests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
)

func Test_GenericGherkins_IsFailedToMatch_WhenMatching(t *testing.T) {
	tc := isFailedToMatchWhenMatchingTestCase

	// Act
	result := tc.IsFailedToMatch()

	// Assert
	tc.ShouldMatchExpectedFirst(t, result)
}

func Test_GenericGherkins_IsFailedToMatch_WhenNotMatching(t *testing.T) {
	tc := isFailedToMatchWhenNotMatchingTestCase

	// Act
	result := tc.IsFailedToMatch()

	// Assert
	tc.ShouldMatchExpectedFirst(t, result)
}

func Test_GenericGherkins_CompareWith_Equal_FromGenericGherkins(t *testing.T) {
	tc := compareWithEqualTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)
	expectedDiff, _ := tc.ExtraArgs.GetAsString("expectedDiff")

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	tc.ShouldMatchExpectedFirst(t, isEqual)
	actual := args.Map{"diff": diff}
	expected := args.Map{"diff": expectedDiff}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_CompareWith_DiffTitle(t *testing.T) {
	tc := compareWithDiffTitleTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)
	expectedDiff, _ := tc.ExtraArgs.GetAsString("expectedDiff")

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	tc.ShouldMatchExpectedFirst(t, isEqual)
	actual := args.Map{"diff": diff}
	expected := args.Map{"diff": expectedDiff}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_CompareWith_DiffInput(t *testing.T) {
	tc := compareWithDiffInputTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)
	expectedDiff, _ := tc.ExtraArgs.GetAsString("expectedDiff")

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	tc.ShouldMatchExpectedFirst(t, isEqual)
	actual := args.Map{"diff": diff}
	expected := args.Map{"diff": expectedDiff}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_CompareWith_BothNil_FromGenericGherkins(t *testing.T) {
	tc := compareWithBothNilTestCase

	// Arrange
	var a *coretestcases.StringBoolGherkins
	var b *coretestcases.StringBoolGherkins
	expectedDiff, _ := tc.ExtraArgs.GetAsString("expectedDiff")

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	tc.ShouldMatchExpectedFirst(t, isEqual)
	actual := args.Map{"diff": diff}
	expected := args.Map{"diff": expectedDiff}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_CompareWith_OneNil_FromGenericGherkins(t *testing.T) {
	tc := compareWithOneNilTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	var b *coretestcases.StringBoolGherkins
	expectedDiff, _ := tc.ExtraArgs.GetAsString("expectedDiff")

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	tc.ShouldMatchExpectedFirst(t, isEqual)
	actual := args.Map{"diff": diff}
	expected := args.Map{"diff": expectedDiff}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_FullString_Basic(t *testing.T) {
	tc := fullStringBasicTestCase

	// Arrange
	g := tc.GetExtra("subject").(*coretestcases.StringBoolGherkins)

	// Act
	result := g.FullString()
	actLines := strings.Split(strings.TrimRight(result, "\n"), "\n")

	// Assert — build args.Map with lineCount + indexed keys
	actual := args.Map{
		"expectedLineCount": fmt.Sprintf("%d", len(actLines)),
	}
	for i, line := range actLines {
		actual[fmt.Sprintf("line%d", i)] = line
	}

	expected := args.Map{}
	for k, v := range tc.ExtraArgs {
		if k != "subject" {
			expected[k] = v
		}
	}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_FullString_Nil_FromGenericGherkins(t *testing.T) {
	tc := fullStringNilTestCase

	// Arrange
	var g *coretestcases.StringBoolGherkins

	// Act
	result := g.FullString()

	// Assert
	expectedResult, _ := tc.ExtraArgs.GetAsString("expectedResult")
	actual := args.Map{"result": result}
	expected := args.Map{"result": expectedResult}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_ShouldBeEqualArgs_Passing(t *testing.T) {
	tc := shouldBeEqualPassingTestCase

	// Act + Assert — ExpectedLines kept for ShouldBeEqualArgs compatibility
	tc.ShouldBeEqualArgsFirst(
		t,
		"line-a",
		"line-b",
	)
}

func Test_GenericGherkins_CaseTitle_UsesTitle(t *testing.T) {
	tc := caseTitleUseTitleTestCase

	// Act
	result := tc.CaseTitle()

	// Assert
	expectedResult, _ := tc.ExtraArgs.GetAsString("expectedResult")
	actual := args.Map{"result": result}
	expected := args.Map{"result": expectedResult}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.Title, ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}

func Test_GenericGherkins_CaseTitle_FallsBackToWhen(t *testing.T) {
	tc := caseTitleFallbackToWhenTestCase

	// Act
	result := tc.CaseTitle()

	// Assert
	expectedResult, _ := tc.ExtraArgs.GetAsString("expectedResult")
	actual := args.Map{"result": result}
	expected := args.Map{"result": expectedResult}
	verifyCaseV1 := coretestcases.CaseV1{Title: tc.CaseTitle(), ExpectedInput: expected}
	verifyCaseV1.ShouldBeEqualMapFirst(t, actual)
}
