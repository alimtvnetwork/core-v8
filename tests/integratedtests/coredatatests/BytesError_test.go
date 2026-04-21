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

package coredatatests

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/coredata"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ===== BytesError Tests =====

// Note: String nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_String_WithBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("hello"),
	}

	got := be.String()

	// Act
	actual := args.Map{"result": got != "hello"}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BytesError.String() should match", actual)
}

func Test_BytesError_String_EmptyBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	got := be.String()

	// Act
	actual := args.Map{"result": got != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BytesError.String() on empty =, want empty", actual)
}

func Test_BytesError_String_CachesResult(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("cached"),
	}

	first := be.String()
	second := be.String()

	// Act
	actual := args.Map{"result": first != second}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected cached string to be identical on second call", actual)
}

func Test_BytesError_HasError_True(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Error: errors.New("some error"),
	}

	// Act
	actual := args.Map{"result": be.HasError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasError() = true", actual)
}

func Test_BytesError_HasError_False(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	// Act
	actual := args.Map{"result": be.HasError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasError() = false", actual)
}

// Note: HasError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmptyError_True(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	// Act
	actual := args.Map{"result": be.IsEmptyError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyError() = true", actual)
}

// Note: IsEmptyError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmptyError_False(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Error: errors.New("err"),
	}

	// Act
	actual := args.Map{"result": be.IsEmptyError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEmptyError() = false", actual)
}

func Test_BytesError_HasBytes_True(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("data"),
	}

	// Act
	actual := args.Map{"result": be.HasBytes()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected HasBytes() = true", actual)
}

func Test_BytesError_HasBytes_False_NilBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	// Act
	actual := args.Map{"result": be.HasBytes()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasBytes() = false for nil bytes", actual)
}

func Test_BytesError_HasBytes_False_EmptyBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte{},
	}

	// Act
	actual := args.Map{"result": be.HasBytes()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasBytes() = false for empty bytes", actual)
}

func Test_BytesError_HasBytes_False_EmptyJson(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("{}"),
	}

	// Act
	actual := args.Map{"result": be.HasBytes()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasBytes() = false for empty JSON {}", actual)
}

func Test_BytesError_HasBytes_False_WithError(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("data"),
		Error: errors.New("err"),
	}

	// Act
	actual := args.Map{"result": be.HasBytes()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected HasBytes() = false when error is present", actual)
}

// Note: Length nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_Length_WithBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("hello"),
	}

	got := be.Length()

	// Act
	actual := args.Map{"result": got != 5}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "BytesError.Length() =, want 5", actual)
}

// Note: IsEmpty nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_IsEmpty_EmptyBytes(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte{},
	}

	// Act
	actual := args.Map{"result": be.IsEmpty()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected IsEmpty() = true for empty bytes", actual)
}

func Test_BytesError_IsEmpty_False(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Bytes: []byte("data"),
	}

	// Act
	actual := args.Map{"result": be.IsEmpty()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected IsEmpty() = false for non-empty bytes", actual)
}

func Test_BytesError_HandleError_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	// Should not panic
	be.HandleError()
}

// Note: HandleError nil receiver test migrated to BytesError_NilReceiver_testcases.go

func Test_BytesError_NilReceiver(t *testing.T) {
	for caseIndex, tc := range bytesErrorNilReceiverTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}

func Test_BytesError_HandleError_Panics(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HandleError to panic", actual)
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleError()
}

func Test_BytesError_HandleErrorWithMsg_NoError(t *testing.T) {
	be := &coredata.BytesError{}

	// Should not panic
	be.HandleErrorWithMsg("prefix: ")
}

func Test_BytesError_HandleErrorWithMsg_PanicsWithMsg(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()
		if r == nil {

	// Act
			actual := args.Map{"result": false}

	// Assert
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected HandleErrorWithMsg to panic", actual)
			return
		}

		msg, ok := r.(string)
		if !ok {
			actual := args.Map{"result": false}
			expected := args.Map{"result": true}
			expected.ShouldBeEqual(t, 0, "expected panic value to be string", actual)
			return
		}

		actual := args.Map{"result": msg != "prefix: boom"}
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "panic message should match", actual)
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleErrorWithMsg("prefix: ")
}

func Test_BytesError_HandleErrorWithMsg_PanicsEmptyMsg(t *testing.T) {
	// Arrange
	defer func() {
		r := recover()

	// Act
		actual := args.Map{"result": r == nil}

	// Assert
		expected := args.Map{"result": false}
		expected.ShouldBeEqual(t, 0, "expected HandleErrorWithMsg to panic", actual)
	}()

	be := &coredata.BytesError{
		Error: errors.New("boom"),
	}
	be.HandleErrorWithMsg("")
}

func Test_BytesError_CombineErrorWithRef_NoError(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	got := be.CombineErrorWithRef("ref1")

	// Act
	actual := args.Map{"result": got != ""}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRef on no error =, want empty", actual)
}

func Test_BytesError_CombineErrorWithRefError_NoError(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{}

	got := be.CombineErrorWithRefError("ref1")

	// Act
	actual := args.Map{"result": got != nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "CombineErrorWithRefError on no error =, want nil", actual)
}

func Test_BytesError_CombineErrorWithRefError_WithError(t *testing.T) {
	// Arrange
	be := &coredata.BytesError{
		Error: errors.New("something failed"),
	}

	got := be.CombineErrorWithRefError("ref1", "ref2")

	// Act
	actual := args.Map{"result": got == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil error from CombineErrorWithRefError", actual)
}
