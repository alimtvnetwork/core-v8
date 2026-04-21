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

package pathinternaltests

import (
	"os"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/internal/pathinternal"
)

// ── Clean ──

func Test_Clean_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": pathinternal.Clean("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "Clean returns empty -- empty", actual)
}

func Test_Clean_Valid(t *testing.T) {
	// Arrange
	result := pathinternal.Clean("/tmp/test/../other")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Clean returns non-empty -- valid", actual)
}

// ── Join ──

func Test_Join(t *testing.T) {
	// Arrange
	result := pathinternal.Join("a", "b", "c")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Join returns correct value -- with args", actual)
}

// ── JoinTemp ──

func Test_JoinTemp(t *testing.T) {
	// Arrange
	result := pathinternal.JoinTemp("subdir", "file.txt")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "JoinTemp returns correct value -- with args", actual)
}

// ── ParentDir ──

func Test_ParentDir_Empty(t *testing.T) {
	// Act
	actual := args.Map{"result": pathinternal.ParentDir("")}

	// Assert
	expected := args.Map{"result": ""}
	expected.ShouldBeEqual(t, 0, "ParentDir returns empty -- empty", actual)
}

func Test_ParentDir_Valid(t *testing.T) {
	// Arrange
	result := pathinternal.ParentDir("/tmp/test/file.txt")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "ParentDir returns non-empty -- valid", actual)
}

// ── Relative ──

func Test_Relative(t *testing.T) {
	// Arrange
	result := pathinternal.Relative("/tmp", "/tmp/sub/file.txt")

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "Relative returns correct value -- with args", actual)
}

// ── IsPathExists ──

func Test_IsPathExists_True(t *testing.T) {
	// Act
	actual := args.Map{"exists": pathinternal.IsPathExists(os.TempDir())}

	// Assert
	expected := args.Map{"exists": true}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns non-empty -- true", actual)
}

func Test_IsPathExists_False(t *testing.T) {
	// Act
	actual := args.Map{"exists": pathinternal.IsPathExists("/nonexistent/path/xyz")}

	// Assert
	expected := args.Map{"exists": false}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns non-empty -- false", actual)
}

// ── GetTemp ──

func Test_GetTemp(t *testing.T) {
	// Arrange
	result := pathinternal.GetTemp()

	// Act
	actual := args.Map{"notEmpty": result != ""}

	// Assert
	expected := args.Map{"notEmpty": true}
	expected.ShouldBeEqual(t, 0, "GetTemp returns correct value -- with args", actual)
}

// ── RemoveDirIf ──

func Test_RemoveDirIf_FalseCondition(t *testing.T) {
	// Arrange
	err := pathinternal.RemoveDirIf(false, "/tmp/nonexist", "test")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf returns non-empty -- false condition", actual)
}

func Test_RemoveDirIf_NonExistentDir(t *testing.T) {
	// Arrange
	err := pathinternal.RemoveDirIf(true, "/tmp/nonexistent_dir_xyz_test", "test")

	// Act
	actual := args.Map{"noErr": err == nil}

	// Assert
	expected := args.Map{"noErr": true}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf returns non-empty -- non-existent dir", actual)
}
