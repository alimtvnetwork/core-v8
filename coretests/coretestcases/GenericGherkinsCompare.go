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
)

// CompareWith performs a structural comparison of two GenericGherkins instances.
//
// Returns isEqual=true if all fields match, otherwise returns a diff string
// describing the first mismatched field.
func (it *GenericGherkins[TInput, TExpect]) CompareWith(
	other *GenericGherkins[TInput, TExpect],
) (isEqual bool, diff string) {
	if it == nil && other == nil {
		return true, ""
	}

	if it == nil || other == nil {
		return false, "one side is nil"
	}

	var diffs []string

	if it.Title != other.Title {
		diffs = append(diffs, fmt.Sprintf("Title: %q != %q", it.Title, other.Title))
	}

	if it.Feature != other.Feature {
		diffs = append(diffs, fmt.Sprintf("Feature: %q != %q", it.Feature, other.Feature))
	}

	if it.Given != other.Given {
		diffs = append(diffs, fmt.Sprintf("Given: %q != %q", it.Given, other.Given))
	}

	if it.When != other.When {
		diffs = append(diffs, fmt.Sprintf("When: %q != %q", it.When, other.When))
	}

	if it.Then != other.Then {
		diffs = append(diffs, fmt.Sprintf("Then: %q != %q", it.Then, other.Then))
	}

	inputStr := fmt.Sprintf("%v", it.Input)
	otherInputStr := fmt.Sprintf("%v", other.Input)
	if inputStr != otherInputStr {
		diffs = append(diffs, fmt.Sprintf("Input: %v != %v", it.Input, other.Input))
	}

	expectedStr := fmt.Sprintf("%v", it.Expected)
	otherExpectedStr := fmt.Sprintf("%v", other.Expected)
	if expectedStr != otherExpectedStr {
		diffs = append(diffs, fmt.Sprintf("Expected: %v != %v", it.Expected, other.Expected))
	}

	actualStr := fmt.Sprintf("%v", it.Actual)
	otherActualStr := fmt.Sprintf("%v", other.Actual)
	if actualStr != otherActualStr {
		diffs = append(diffs, fmt.Sprintf("Actual: %v != %v", it.Actual, other.Actual))
	}

	if it.IsMatching != other.IsMatching {
		diffs = append(diffs, fmt.Sprintf("IsMatching: %v != %v", it.IsMatching, other.IsMatching))
	}

	if len(diffs) == 0 {
		return true, ""
	}

	return false, strings.Join(diffs, "; ")
}
