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

package coretestcases

import "testing"

// ShouldBeEqualFirst asserts using ShouldBeEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeTrimEqualFirst asserts using ShouldBeTrimEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeTrimEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeTrimEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeSortedEqualFirst asserts using ShouldBeSortedEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeSortedEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeSortedEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldContainsFirst asserts using ShouldContains with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldContainsFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldContains(
		t,
		0,
		actualElements...,
	)
}

// ShouldStartsWithFirst asserts using ShouldStartsWith with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldStartsWithFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldStartsWith(
		t,
		0,
		actualElements...,
	)
}

// ShouldEndsWithFirst asserts using ShouldEndsWith with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldEndsWithFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldEndsWith(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeNotEqualFirst asserts using ShouldBeNotEqual with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeNotEqualFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeNotEqual(
		t,
		0,
		actualElements...,
	)
}

// ShouldBeRegexFirst asserts using ShouldBeRegex with caseIndex=0.
// Use for named single test cases (non-loop).
func (it CaseV1) ShouldBeRegexFirst(
	t *testing.T,
	actualElements ...string,
) {
	t.Helper()

	it.ShouldBeRegex(
		t,
		0,
		actualElements...,
	)
}
