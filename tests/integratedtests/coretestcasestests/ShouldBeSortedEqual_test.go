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
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/coretests/results"
	"github.com/alimtvnetwork/core-v8/corevalidator"
	"github.com/alimtvnetwork/core-v8/enums/stringcompareas"
	"github.com/alimtvnetwork/core-v8/issetter"
)

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldBeSortedEqual — covers CaseV1.go L484-492
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldBeSortedEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "sorted equal direct",
		ExpectedInput: []string{"a", "b"},
	}
	c.ShouldBeSortedEqual(t, 0, "a", "b")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldContains — covers CaseV1.go L498-505
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldContains(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "contains direct",
		ExpectedInput: "hello",
	}
	c.ShouldContains(t, 0, "hello world")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldStartsWith — covers CaseV1.go L511-518
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldStartsWith(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "starts with direct",
		ExpectedInput: "hello",
	}
	c.ShouldStartsWith(t, 0, "hello world")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldEndsWith — covers CaseV1.go L524-531
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldEndsWith(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "ends with direct",
		ExpectedInput: "world",
	}
	c.ShouldEndsWith(t, 0, "hello world")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldBeNotEqual — covers CaseV1.go L537-544
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldBeNotEqual(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "not equal direct",
		ExpectedInput: "hello",
	}
	c.ShouldBeNotEqual(t, 0, "world")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldBeRegex — covers CaseV1.go L554-561
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldBeRegex(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "regex direct",
		ExpectedInput: "hel.*",
	}
	c.ShouldBeRegex(t, 0, "hello world")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldBeTrimRegex — covers CaseV1.go L571-579
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldBeTrimRegex(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "trim regex direct",
		ExpectedInput: "hel.*",
	}
	c.ShouldBeTrimRegex(t, 0, "  hello world  ")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 ShouldHaveNoError — covers CaseV1.go L586-601
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldHaveNoError(t *testing.T) {
	c := coretestcases.CaseV1{
		Title: "no error direct",
	}
	c.ShouldHaveNoError(t, "additional", 0, nil)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 PrepareTitle — covers CaseV1.go L636-643
// ══════════════════════════════════════════════════════════════════════════════

func Test_PrepareTitle(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "title"}
	title := c.PrepareTitle(1, "suffix")

	// Act
	actual := args.Map{"hasTitle": title != ""}

	// Assert
	expected := args.Map{"hasTitle": true}
	expected.ShouldBeEqual(t, 0, "PrepareTitle returns non-empty", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 AsSimpleTestCaseWrapperContractsBinder — covers CaseV1.go L649-651
// ══════════════════════════════════════════════════════════════════════════════

func Test_AsSimpleTestCaseWrapperContractsBinder(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{Title: "binder"}
	binder := c.AsSimpleTestCaseWrapperContractsBinder()

	// Act
	actual := args.Map{"notNil": binder != nil}

	// Assert
	expected := args.Map{"notNil": true}
	expected.ShouldBeEqual(t, 0, "ContractsBinder returns non-nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyTypeOfMatch — covers CaseV1.go L89-110
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyTypeOfMatch(t *testing.T) {
	vt := coretests.NewVerifyTypeOf("actual")
	c := coretestcases.CaseV1{
		Title:         "type match",
		ExpectedInput: "hello",
		IsEnable:      issetter.True,
		VerifyTypeOf:  vt,
	}
	c.VerifyTypeOfMatch(t, 0, "actual")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyTypeOfMust — covers CaseV1.go L120-141
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyTypeOfMust(t *testing.T) {
	vt := coretests.NewVerifyTypeOf("actual")
	c := coretestcases.CaseV1{
		Title:         "type must",
		ExpectedInput: "hello",
		IsEnable:      issetter.True,
		VerifyTypeOf:  vt,
	}
	c.VerifyTypeOfMust(t, 0, "actual")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyType — covers CaseV1.go L151-172
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyType(t *testing.T) {
	vt := coretests.NewVerifyTypeOf("actual")
	c := coretestcases.CaseV1{
		Title:         "verify type",
		ExpectedInput: "hello",
		IsEnable:      issetter.True,
		VerifyTypeOf:  vt,
	}
	c.VerifyType(t, 0, "actual")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyTypeMust — covers CaseV1.go L182-203
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyTypeMust(t *testing.T) {
	vt := coretests.NewVerifyTypeOf("actual")
	c := coretestcases.CaseV1{
		Title:         "verify type must",
		ExpectedInput: "hello",
		IsEnable:      issetter.True,
		VerifyTypeOf:  vt,
	}
	c.VerifyTypeMust(t, 0, "actual")
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 SliceValidatorCondition — covers CaseV1.go L241-256
// ══════════════════════════════════════════════════════════════════════════════

func Test_SliceValidatorCondition(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "slice validator condition",
		ExpectedInput: "hello",
	}
	sv := c.SliceValidatorCondition(
		stringcompareas.Equal,
		corevalidator.DefaultDisabledCoreCondition,
		[]string{"hello"},
	)

	// Act
	actual := args.Map{"hasActual": len(sv.ActualLines) > 0}

	// Assert
	expected := args.Map{"hasActual": true}
	expected.ShouldBeEqual(t, 0, "SliceValidatorCondition populated", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyAllEqualCondition — covers CaseV1.go L217-228
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyAllEqualCondition(t *testing.T) {
	// Arrange
	c := coretestcases.CaseV1{
		Title:         "verify all equal condition",
		ExpectedInput: "hello",
	}
	err := c.VerifyAllEqualCondition(0, corevalidator.DefaultDisabledCoreCondition, "hello")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyAllEqualCondition passes", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 VerifyError with type verification — covers CaseV1.go L349-366
// ══════════════════════════════════════════════════════════════════════════════

func Test_VerifyError_WithTypeVerify(t *testing.T) {
	// Act
	actualStr := "hello"
	vt := &coretests.VerifyTypeOf{
		IsVerify:      issetter.True,
		ArrangeInput:  reflect.TypeOf(&actualStr),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf(""),
	}
	c := coretestcases.CaseV1{
		Title:         "verify error with type",
		ExpectedInput: "hello",
		IsEnable:      issetter.True,
		VerifyTypeOf:  vt,
	}
	err := c.VerifyError(0, stringcompareas.Equal, "hello")
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "VerifyError with type nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseNilSafe — covers CaseNilSafe.go + CaseNilSafeAssertHelper.go
// ══════════════════════════════════════════════════════════════════════════════

func Test_CaseNilSafe_ShouldBeSafe_FromShouldBeSortedEqualI(t *testing.T) {
	tc := coretestcases.CaseNilSafe{
		Title: "ClonePtr nil safe direct",
		Func:  (*coretests.DraftType).ClonePtr,
		Expected: results.ResultAny{
			Panicked: false,
		},
		CompareFields: []string{"panicked"},
	}

	// Assert
	tc.ShouldBeSafe(t, 0)
}

// ══════════════════════════════════════════════════════════════════════════════
// CaseV1 First variants — covers CaseV1FirstAssertions.go
// ══════════════════════════════════════════════════════════════════════════════

func Test_ShouldContainsFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "contains first direct",
		ExpectedInput: "hel",
	}
	c.ShouldContainsFirst(t, "hello world")
}

func Test_ShouldStartsWithFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "starts with first direct",
		ExpectedInput: "hello",
	}
	c.ShouldStartsWithFirst(t, "hello world")
}

func Test_ShouldBeRegexFirst(t *testing.T) {
	c := coretestcases.CaseV1{
		Title:         "regex first direct",
		ExpectedInput: "hel.*",
	}
	c.ShouldBeRegexFirst(t, "hello world")
}
