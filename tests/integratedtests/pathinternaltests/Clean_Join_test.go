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

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/internal/pathinternal"
)

func Test_Clean(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"empty":       pathinternal.Clean(""),
		"slashNorm":   pathinternal.Clean("/a//b///c"),
		"backslash":   pathinternal.Clean("a\\b\\c"),
		"doubleBack":  pathinternal.Clean("a\\\\b"),
		"normal":      pathinternal.Clean("/a/b/c"),
	}
	expected := args.Map{
		"empty":       "",
		"slashNorm":   "/a/b/c",
		"backslash":   "a/b/c",
		"doubleBack":  "a/b",
		"normal":      "/a/b/c",
	}
	expected.ShouldBeEqual(t, 0, "Clean returns correct value -- with args", actual)
}

func Test_Join_FromCleanJoin(t *testing.T) {
	// Arrange & Act
	result := pathinternal.Join("a", "b", "c")

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Join returns correct value -- with args", actual)
}

func Test_ParentDir(t *testing.T) {
	// Arrange & Act
	result := pathinternal.ParentDir("/a/b/c")
	emptyResult := pathinternal.ParentDir("")

	// Assert
	actual := args.Map{
		"notEmpty":    result != "",
		"emptyResult": emptyResult,
	}
	expected := args.Map{
		"notEmpty":    true,
		"emptyResult": "",
	}
	expected.ShouldBeEqual(t, 0, "ParentDir returns correct value -- with args", actual)
}

func Test_Relative_FromCleanJoin(t *testing.T) {
	// Arrange & Act
	result := pathinternal.Relative("/a/b", "/a/b/c/d")

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "Relative returns correct value -- with args", actual)
}

func Test_IsPathExists(t *testing.T) {
	// Arrange & Act
	actual := args.Map{
		"tempExists":  pathinternal.IsPathExists(os.TempDir()),
		"fakeNotExist": pathinternal.IsPathExists("/surely/does/not/exist/12345"),
	}
	expected := args.Map{
		"tempExists":  true,
		"fakeNotExist": false,
	}
	expected.ShouldBeEqual(t, 0, "IsPathExists returns correct value -- with args", actual)
}

func Test_GetTemp_FromCleanJoin(t *testing.T) {
	// Arrange & Act
	result := pathinternal.GetTemp()

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "GetTemp returns correct value -- with args", actual)
}

func Test_JoinTemp_FromCleanJoin(t *testing.T) {
	// Arrange & Act
	result := pathinternal.JoinTemp("subdir", "file.txt")

	// Assert
	actual := args.Map{
		"notEmpty": result != "",
	}
	expected := args.Map{
		"notEmpty": true,
	}
	expected.ShouldBeEqual(t, 0, "JoinTemp returns correct value -- with args", actual)
}

func Test_RemoveDirIf_NonExistent(t *testing.T) {
	// Arrange & Act — removing non-existent dir with flag=false should be no-op
	err := pathinternal.RemoveDirIf(false, "/nonexistent/path/12345", "test")

	// Assert
	actual := args.Map{
		"noError": err == nil,
	}
	expected := args.Map{
		"noError": true,
	}
	expected.ShouldBeEqual(t, 0, "RemoveDirIf_NonExistent returns correct value -- with args", actual)
}
