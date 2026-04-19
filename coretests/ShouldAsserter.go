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

package coretests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type V2ShouldAsserter interface {
	ShouldBeSimpleAsserter
	V2ShouldBeExplicitAsserter
}

type ShouldBeSimpleAsserter interface {
	ShouldBe(
		caseIndex int,
		t *testing.T,
		assert convey.Assertion,
		actual any,
	)
}

type ShouldBeEqualAsserter interface {
	ShouldBeEqual(
		caseIndex int,
		t *testing.T,
		actual any,
	)
}

type ShouldHaveNoErrorAsserter interface {
	ShouldHaveNoError(
		caseIndex int,
		t *testing.T,
		err error,
	)
}

type ShouldContainsAsserter interface {
	ShouldContains(
		caseIndex int,
		t *testing.T,
		actual any,
	)
}

type V2ShouldBeExplicitAsserter interface {
	ShouldBeExplicit(
		isValidateType bool,
		caseIndex int,
		t *testing.T,
		title string,
		actual any,
		assert convey.Assertion,
		expected any,
	)
}

type V1ShouldBeExplicitAsserter interface {
	ShouldBeExplicit(
		caseIndex int,
		t *testing.T,
		title string,
		actual any,
		assert convey.Assertion,
		expected any,
	)
}

type V1ShouldAsserter interface {
	ShouldBeSimpleAsserter
	V1ShouldBeExplicitAsserter
}
