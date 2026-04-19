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
	"github.com/alimtvnetwork/core/coretests/args"
)

// GenericGherkins is a typed test case representation using
// the Gherkin (Given/When/Then) pattern.
//
// Use this when you want compile-time type safety for Input/Expected
// fields instead of extracting values from args.Map at runtime.
//
// TInput  — type of the test input (e.g., string for regex pattern)
// TExpect — type of the expected result (e.g., bool for IsMatch)
type GenericGherkins[TInput, TExpect any] struct {
	// Title is the test case header / scenario name.
	Title string

	// Feature describes the feature being tested.
	Feature string

	// Given describes the precondition.
	Given string

	// When describes the scenario.
	When string

	// Then describes the expected outcome.
	Then string

	// Input is the typed input value for the test.
	Input TInput

	// Expected is the typed expected result.
	Expected TExpect

	// Actual is the typed actual result, set after the Act phase.
	Actual TExpect

	// IsMatching indicates whether a match is expected.
	// Use for validation/matching tests (e.g., regex, search).
	IsMatching bool

	// ExpectedLines holds the expected string output lines for assertion.
	// Used by ShouldBeEqualUsingExpected when comparing act lines.
	ExpectedLines []string

	// ExtraArgs provides optional dynamic key-value pairs
	// for overflow beyond the typed fields.
	ExtraArgs args.Map
}
