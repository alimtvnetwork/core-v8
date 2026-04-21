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

package args

import (
	"errors"
	"testing"

	"github.com/alimtvnetwork/core-v8/errcore"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// ShouldBeEqual compares this Map (expected) against actual Map.
//
// Both maps are compiled to sorted "key : value" string lines using
// CompileToStrings(), then compared line-by-line.
//
// On mismatch, a copy-pasteable Go literal block is printed showing
// each entry on its own indexed line in Go literal format.
//
// Usage:
//
//	expected := args.Map{"result": 42}
//	expected.ShouldBeEqual(t, 0, "TestTitle", actual)
func (it Map) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	title string,
	actual Map,
) {
	t.Helper()

	actualLines := actual.CompileToStrings()
	expectedLines := it.CompileToStrings()

	hasMismatch := errcore.HasAnyMismatchOnLines(actualLines, expectedLines)

	var validationErr error

	if hasMismatch {
		errcore.PrintDiffOnMismatch(caseIndex, title, actualLines, expectedLines)

		mapErrMsg := errcore.MapMismatchError(
			t.Name(),
			caseIndex,
			title,
			actual.GoLiteralLines(),
			it.GoLiteralLines(),
		)

		validationErr = errors.New(mapErrMsg)
	}

	convey.Convey(
		title, t, func() {
			convey.So(
				validationErr,
				should.BeNil,
			)
		},
	)
}
