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

package coreteststests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core-v8/coretests"
	"github.com/alimtvnetwork/core-v8/coretests/args"
	"github.com/alimtvnetwork/core-v8/coretests/coretestcases"
	"github.com/alimtvnetwork/core-v8/issetter"
	should "github.com/smarty/assertions/should"
)

// ── DraftType.IsEqual: inner fields mismatch (line 148-150) ──

func Test_DraftType_IsEqual_InnerFieldF2Mismatch(t *testing.T) {
	// Arrange
	a := &coretests.DraftType{
		SampleString1: "s1",
		SampleString2: "s2",
		SampleInteger: 1,
	}
	b := &coretests.DraftType{
		SampleString1: "s1",
		SampleString2: "s2",
		SampleInteger: 1,
	}

	a.SetF2Integer(10)
	b.SetF2Integer(20)

	// Act
	result := a.IsEqual(true, b)

	// Assert
	actual := args.Map{
		"equal": fmt.Sprintf("%v", result),
	}
	tc := coretestcases.CaseV1{
		Title: "IsEqual returns false -- inner field f2Integer mismatch",
		ExpectedInput: args.Map{
			"equal": "false",
		},
	}
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ── BaseTestCase.ShouldBe: disabled case triggers noPrintAssert (BaseTestCase) ──

func Test_BaseTestCase_ShouldBe_DisabledCase(t *testing.T) {
	// Arrange
	tc := coretests.BaseTestCase{
		IsEnable: issetter.False,
	}
	tc.Title = "disabled case"
	tc.ExpectedInput = "expected"

	// Act & Assert — ShouldBe with disabled calls BaseTestCase.noPrintAssert
	tc.ShouldBe(
		0,
		t,
		should.Equal,
		"expected",
	)
}

// ── BaseTestCase.ShouldBeExplicit: failure branch (lines 88-92) ──
// This branch logs when assertion comparison string is non-empty (mismatch).
// Triggering it requires passing mismatched values, which causes GoConvey to
// report a test failure. This is an assertion-helper diagnostic path that only
// activates during actual test failures. Accepted gap — testing it would require
// intentionally failing a test.

// ── SimpleTestCase.noPrintAssert (lines 89-102) ──
// Dead code: SimpleTestCase has no IsEnable field and no code path calls
// noPrintAssert. It's an orphaned method. Accepted gap.

// ── DraftType.JsonString/JsonBytes panic (lines 174-175, 184-185) ──
// json.Marshal on DraftType (simple struct with basic types) cannot fail.
// Defensive panic branches are unreachable. Accepted gap.

// ── SkipOnUnix (line 12-14) ──
// Platform-dependent: only executes on Unix systems.
// Accepted gap on Windows test runners.
