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

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
)

// String returns a Gherkins-formatted representation using index 0.
func (it *GenericGherkins[TInput, TExpect]) String() string {
	return it.ToString(constants.Zero)
}

// ToString returns a Gherkins-formatted string with the given test index.
func (it *GenericGherkins[TInput, TExpect]) ToString(testIndex int) string {
	return errcore.GherkinsString(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
	)
}

// GetWithExpectation returns a Gherkins string that includes
// the Actual and Expected values for diagnostic output.
func (it *GenericGherkins[TInput, TExpect]) GetWithExpectation(
	testIndex int,
) string {
	return errcore.GherkinsStringWithExpectation(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
		it.Actual,
		it.Expected,
	)
}

// GetMessageConditional returns the Gherkins string with or without
// expectation details based on the isExpectation flag.
func (it *GenericGherkins[TInput, TExpect]) GetMessageConditional(
	isExpectation bool,
	testIndex int,
) string {
	if isExpectation {
		return it.GetWithExpectation(testIndex)
	}

	return it.ToString(testIndex)
}

// FullString returns a verbose multi-line representation of all fields
// for debugging purposes.
func (it *GenericGherkins[TInput, TExpect]) FullString() string {
	if it == nil {
		return "<nil GenericGherkins>"
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Title:      %s\n", it.Title))
	sb.WriteString(fmt.Sprintf("Feature:    %s\n", it.Feature))
	sb.WriteString(fmt.Sprintf("Given:      %s\n", it.Given))
	sb.WriteString(fmt.Sprintf("When:       %s\n", it.When))
	sb.WriteString(fmt.Sprintf("Then:       %s\n", it.Then))
	sb.WriteString(fmt.Sprintf("Input:      %v\n", it.Input))
	sb.WriteString(fmt.Sprintf("Expected:   %v\n", it.Expected))
	sb.WriteString(fmt.Sprintf("Actual:     %v\n", it.Actual))
	sb.WriteString(fmt.Sprintf("IsMatching: %v\n", it.IsMatching))

	if len(it.ExtraArgs) > 0 {
		sb.WriteString(fmt.Sprintf("ExtraArgs:  %v\n", it.ExtraArgs))
	}

	return sb.String()
}
