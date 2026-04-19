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

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/corevalidator"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

// ==========================================
// TextValidators — collection operations
// ==========================================

func Test_TextValidators_NewEmpty(t *testing.T) {
	// Arrange
	tc := textValidatorsNewEmptyTestCase
	v := corevalidator.NewTextValidators(5)

	// Act
	actual := args.Map{
		"isEmpty": v.IsEmpty(),
		"length":  v.Length(),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Add_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsAddTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{"length": v.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Adds_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsAddsTestCase
	v := corevalidator.NewTextValidators(2)
	v.Adds(
		corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal},
		corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal},
	)

	// Act
	actual := args.Map{"length": v.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_Adds_Empty_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsAddsEmptyTestCase
	v := corevalidator.NewTextValidators(2)
	v.Adds()

	// Act
	actual := args.Map{"length": v.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_AddSimple_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsAddSimpleTestCase
	v := corevalidator.NewTextValidators(1)
	v.AddSimple("test", stringcompareas.Contains)

	// Act
	actual := args.Map{"length": v.Length()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_HasIndex_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsHasIndexTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{
		"hasIndex0": v.HasIndex(0),
		"hasIndex1": v.HasIndex(1),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_LastIndex(t *testing.T) {
	// Arrange
	tc := textValidatorsLastIndexTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Add(corevalidator.TextValidator{Search: "b", SearchAs: stringcompareas.Equal})

	// Act
	actual := args.Map{"lastIndex": v.LastIndex()}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.IsMatch
// ==========================================

func Test_TextValidators_IsMatch_EmptyValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsIsMatchEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	// Act
	actual := args.Map{"isMatch": v.IsMatch("anything", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_IsMatch_AllPass_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsIsMatchAllPassTestCase
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("world", stringcompareas.Contains)

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_IsMatch_OneFails_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsIsMatchOneFailsTestCase
	v := corevalidator.NewTextValidators(2)
	v.AddSimple("hello", stringcompareas.Contains)
	v.AddSimple("xyz", stringcompareas.Contains)

	// Act
	actual := args.Map{"isMatch": v.IsMatch("hello world", true)}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.IsMatchMany
// ==========================================

func Test_TextValidators_IsMatchMany_EmptyValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsIsMatchManyEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	// Act
	actual := args.Map{"isMatch": v.IsMatchMany(false, true, "a", "b")}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.VerifyFirstError
// ==========================================

func Test_TextValidators_VerifyFirstError_AllPass(t *testing.T) {
	// Arrange
	tc := textValidatorsVerifyFirstAllPassTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	// Act
	actual := args.Map{"hasError": v.VerifyFirstError(0, "hello", true) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_VerifyFirstError_Fails(t *testing.T) {
	// Arrange
	tc := textValidatorsVerifyFirstFailTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	// Act
	actual := args.Map{"hasError": v.VerifyFirstError(0, "world", true) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_VerifyFirstError_Empty_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsVerifyFirstEmptyTestCase
	v := corevalidator.NewTextValidators(0)

	// Act
	actual := args.Map{"hasError": v.VerifyFirstError(0, "anything", true) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.AllVerifyError
// ==========================================

func Test_TextValidators_AllVerifyError_AllPass(t *testing.T) {
	// Arrange
	tc := textValidatorsAllVerifyPassTestCase
	v := corevalidator.NewTextValidators(1)
	v.Add(corevalidator.TextValidator{
		Search:    "hello",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	// Act
	actual := args.Map{"hasError": v.AllVerifyError(0, "hello", true) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

func Test_TextValidators_AllVerifyError_MultipleFail(t *testing.T) {
	// Arrange
	tc := textValidatorsAllVerifyFailTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{
		Search:    "x",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})
	v.Add(corevalidator.TextValidator{
		Search:    "y",
		SearchAs:  stringcompareas.Equal,
		Condition: corevalidator.DefaultDisabledCoreCondition,
	})

	// Act
	actual := args.Map{"hasError": v.AllVerifyError(0, "z", true) != nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// TextValidators.Dispose
// ==========================================

func Test_TextValidators_Dispose_FromTextValidators(t *testing.T) {
	// Arrange
	tc := textValidatorsDisposeTestCase
	v := corevalidator.NewTextValidators(2)
	v.Add(corevalidator.TextValidator{Search: "a", SearchAs: stringcompareas.Equal})
	v.Dispose()

	// Act
	actual := args.Map{"isNil": v.Items == nil}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// (nil receiver tests migrated to TextValidators_NilReceiver_testcases.go)
