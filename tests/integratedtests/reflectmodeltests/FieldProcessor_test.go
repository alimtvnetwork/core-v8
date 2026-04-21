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
	"reflect"
	"testing"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ===== FieldProcessor Tests =====

func Test_FieldProcessor_IsFieldType_Match(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Name", 0)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual = args.Map{"result": fp.IsFieldType(reflect.TypeOf(""))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFieldType(string) = true for Name field", actual)
}

func Test_FieldProcessor_IsFieldType_NoMatch(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Name", 0)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual = args.Map{"result": fp.IsFieldType(reflect.TypeOf(0))}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsFieldType(int) = false for Name (string) field", actual)
}

// Note: IsFieldType nil receiver test migrated to FieldProcessor_NilReceiver_testcases.go

func Test_FieldProcessor_IsFieldKind_Match(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Age", 1)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Age", actual)

	actual = args.Map{"result": fp.IsFieldKind(reflect.Int)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsFieldKind(Int) = true for Age field", actual)
}

func Test_FieldProcessor_IsFieldKind_NoMatch(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Age", 1)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Age", actual)

	actual = args.Map{"result": fp.IsFieldKind(reflect.String)}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsFieldKind(String) = false for Age (int) field", actual)
}

// Note: IsFieldKind nil receiver test migrated to FieldProcessor_NilReceiver_testcases.go

func Test_FieldProcessor_NilReceiver(t *testing.T) {
	for caseIndex, tc := range fieldProcessorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_FieldProcessor_BoolField(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Active", 2)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Active", actual)

	actual = args.Map{"result": fp.IsFieldKind(reflect.Bool)}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Active field to be Bool kind", actual)

	actual = args.Map{"result": fp.IsFieldType(reflect.TypeOf(true))}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected Active field to match bool type", actual)
}

func Test_FieldProcessor_FieldData(t *testing.T) {
	// Arrange
	fp := newFieldProcessor("Name", 0)

	// Act
	actual := args.Map{"result": fp == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "failed to create FieldProcessor for Name", actual)

	actual = args.Map{"result": fp.Name != "Name"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Name should match", actual)

	actual = args.Map{"result": fp.Index != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Index =, want 0", actual)

	actual = args.Map{"result": fp.Field.Name != "Name"}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "Field.Name should match", actual)
}
