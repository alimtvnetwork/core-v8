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

import "github.com/alimtvnetwork/core-v8/coretests"

// CaseTitle returns the test case title (Title or When fallback).
//
// Implements TypedTestCaseWrapper[TInput, TExpect].CaseTitle.
func (it *GenericGherkins[TInput, TExpect]) CaseTitle() string {
	if it.Title != "" {
		return it.Title
	}

	return it.When
}

// TypedInput returns the typed input value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedInput.
func (it *GenericGherkins[TInput, TExpect]) TypedInput() TInput {
	return it.Input
}

// TypedExpected returns the typed expected value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedExpected.
func (it *GenericGherkins[TInput, TExpect]) TypedExpected() TExpect {
	return it.Expected
}

// TypedActual returns the typed actual value.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].TypedActual.
func (it *GenericGherkins[TInput, TExpect]) TypedActual() TExpect {
	return it.Actual
}

// SetTypedActual sets the typed actual value after the Act phase.
//
// Implements TypedTestCaseWrapper[TInput, TExpect].SetTypedActual.
func (it *GenericGherkins[TInput, TExpect]) SetTypedActual(actual TExpect) {
	it.Actual = actual
}

// AsTypedTestCaseWrapper returns the GenericGherkins as a
// TypedTestCaseWrapper[TInput, TExpect] interface.
//
// Implements TypedTestCaseWrapperContractsBinder[TInput, TExpect].
func (it *GenericGherkins[TInput, TExpect]) AsTypedTestCaseWrapper() coretests.TypedTestCaseWrapper[TInput, TExpect] {
	return it
}
