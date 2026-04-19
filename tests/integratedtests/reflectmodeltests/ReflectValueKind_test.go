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

package reflectmodeltests

import (
	"testing"

	"github.com/alimtvnetwork/core/reflectcore/reflectmodel"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ===== ReflectValueKind Tests =====

func Test_InvalidReflectValueKindModel_FromReflectValueKind(t *testing.T) {
	// Arrange
	rvk := reflectmodel.InvalidReflectValueKindModel("test error")

	// Act
	actual := args.Map{"result": rvk.IsValid}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsValid = false", actual)

	actual = args.Map{"result": rvk.HasError()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasError() = true", actual)

	actual = args.Map{"result": rvk.Error.Error() != "test error"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Error should match", actual)
}

// Note: All nil receiver tests migrated to ReflectValueKind_NilReceiver_testcases.go

func Test_ReflectValueKind_NilReceiver_Reflectvaluekind(t *testing.T) {
	for caseIndex, tc := range reflectValueKindNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_ReflectValueKind_IsInvalid_NotValid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	// Act
	actual := args.Map{"result": rvk.IsInvalid()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsInvalid() = true when IsValid=false", actual)
}

func Test_ReflectValueKind_IsEmptyError_NoError(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{}

	// Act
	actual := args.Map{"result": rvk.IsEmptyError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyError() = true when no error", actual)
}

func Test_ReflectValueKind_PkgPath_NotValid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.PkgPath()

	// Act
	actual := args.Map{"result": got != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected PkgPath() = empty when IsValid=false", actual)
}

func Test_ReflectValueKind_TypeName_NotValid(t *testing.T) {
	// Arrange
	rvk := &reflectmodel.ReflectValueKind{IsValid: false}

	got := rvk.TypeName()

	// Act
	actual := args.Map{"result": got != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected TypeName() = empty when IsValid=false", actual)
}
