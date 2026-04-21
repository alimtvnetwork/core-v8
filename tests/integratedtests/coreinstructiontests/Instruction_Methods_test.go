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

package coreinstructiontests

import (
	"github.com/alimtvnetwork/core-v8/coreinstruction"
	"regexp"
	"testing"
	"github.com/alimtvnetwork/core-v8/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsContinueOnError — IsExitOnError nil receiver branch (line 7)
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseIsContinueOnError_IsExitOnError_Nil(t *testing.T) {
	// Arrange
	var b *coreinstruction.BaseIsContinueOnError

	// Act
	actual := args.Map{"result": b.IsExitOnError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for nil receiver", actual)
}

func Test_BaseIsContinueOnError_IsExitOnError_ContinueTrue(t *testing.T) {
	// Arrange
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: true}

	// Act
	actual := args.Map{"result": b.IsExitOnError()}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false when IsContinueOnError is true", actual)
}

func Test_BaseIsContinueOnError_IsExitOnError_ContinueFalse(t *testing.T) {
	// Arrange
	b := &coreinstruction.BaseIsContinueOnError{IsContinueOnError: false}

	// Act
	actual := args.Map{"result": b.IsExitOnError()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true when IsContinueOnError is false", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseIsSecure — NewSecure, NewPlain, IsPlainText, IsIncludePayload (lines 7-25)
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseIsSecure_NewSecure(t *testing.T) {
	// Arrange
	s := coreinstruction.NewSecure()

	// Act
	actual := args.Map{"result": s.IsSecure}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected secure", actual)
	actual = args.Map{"result": s.IsPlainText()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not plain text", actual)
	actual = args.Map{"result": s.IsIncludePayload()}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not include payload", actual)
}

func Test_BaseIsSecure_NewPlain(t *testing.T) {
	// Arrange
	s := coreinstruction.NewPlain()

	// Act
	actual := args.Map{"result": s.IsSecure}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected not secure", actual)
	actual = args.Map{"result": s.IsPlainText()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected plain text", actual)
	actual = args.Map{"result": s.IsIncludePayload()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected include payload", actual)
}

func Test_BaseIsSecure_NilReceiver(t *testing.T) {
	// Arrange
	var s *coreinstruction.BaseIsSecure

	// Act
	actual := args.Map{"result": s.IsPlainText()}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected plain text for nil", actual)
	actual = args.Map{"result": s.IsIncludePayload()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected include payload for nil", actual)
}

// ══════════════════════════════════════════════════════════════════════════════
// BaseTags — NewTagsPtr, TagsLength nil, TagsHashset, IsAnyTagMatchesRegex
// ══════════════════════════════════════════════════════════════════════════════

func Test_BaseTags_NewTagsPtr_Empty(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTagsPtr([]string{})

	// Act
	actual := args.Map{"result": bt == nil}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected non-nil", actual)
	actual = args.Map{"result": bt.TagsLength() != 0}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0 tags", actual)
}

func Test_BaseTags_NewTagsPtr_NonEmpty(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTagsPtr([]string{"a", "b"})

	// Act
	actual := args.Map{"result": bt.TagsLength() != 2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 2 tags", actual)
}

func Test_BaseTags_TagsLength_NilTags(t *testing.T) {
	// Arrange
	bt := coreinstruction.BaseTags{Tags: nil}

	// Act
	actual := args.Map{"result": bt.TagsLength() != 0}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected 0", actual)
}

func Test_BaseTags_TagsHashset_Cached(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{"x", "y"})
	h1 := bt.TagsHashset()
	h2 := bt.TagsHashset()

	// Act
	actual := args.Map{"result": h1 != h2}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected same pointer (cached)", actual)
}

func Test_BaseTags_HasAllTags(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{"a", "b", "c"})

	// Act
	actual := args.Map{"result": bt.HasAllTags("a", "b")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": bt.HasAllTags("a", "z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
	// Empty tags => true
	actual = args.Map{"result": bt.HasAllTags()}
	expected = args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true for empty", actual)
}

func Test_BaseTags_HasAnyTags(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{"a", "b"})

	// Act
	actual := args.Map{"result": bt.HasAnyTags("a", "z")}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected true", actual)
	actual = args.Map{"result": bt.HasAnyTags("x", "z")}
	expected = args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false", actual)
}

func Test_BaseTags_IsAnyTagMatchesRegex_Empty(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{})
	r := regexp.MustCompile(`.*`)

	// Act
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected false for empty tags", actual)
}

func Test_BaseTags_IsAnyTagMatchesRegex_Match(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{"hello-world", "foo"})
	r := regexp.MustCompile(`^hello`)

	// Act
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}

	// Assert
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "expected match", actual)
}

func Test_BaseTags_IsAnyTagMatchesRegex_NoMatch(t *testing.T) {
	// Arrange
	bt := coreinstruction.NewTags([]string{"foo", "bar"})
	r := regexp.MustCompile(`^hello`)

	// Act
	actual := args.Map{"result": bt.IsAnyTagMatchesRegex(r)}

	// Assert
	expected := args.Map{"result": false}
	expected.ShouldBeEqual(t, 0, "expected no match", actual)
}
