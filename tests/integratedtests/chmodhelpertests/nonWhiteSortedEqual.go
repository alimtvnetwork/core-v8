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

package chmodhelpertests

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// isStackTraceLine returns true if the line is a stack trace artifact
// that should be stripped before comparison.
// See issues/chmodhelpertests-stack-trace-mismatch.md
func isStackTraceLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}
	if trimmed == "Stack-Trace:" {
		return true
	}
	// File path references like "- /Users/.../file.go:40"
	if strings.HasPrefix(trimmed, "- /") {
		return true
	}
	// Function name references like "- ErrorRefOnly" or "- getVerifyRwxInternalError"
	if strings.HasPrefix(trimmed, "- ") && !strings.Contains(trimmed, " ") {
		// Single-token after "- " means it's a function/source ref, not error content
		return true
	}
	return false
}

// isStackTraceNormalizedLine checks if a SORTED-TOKEN line is a stack trace artifact.
// After tokens are sorted, "- ErrorRefOnly" becomes "- ErrorRefOnly" (same)
// but we need to also check the sorted form.
func isStackTraceNormalizedLine(sortedLine string) bool {
	trimmed := strings.TrimSpace(sortedLine)
	if trimmed == "" {
		return false
	}
	if trimmed == "Stack-Trace:" {
		return true
	}
	// After sorting, file path lines start with "- /" still
	if strings.HasPrefix(trimmed, "- /") {
		return true
	}
	// Check for single-word function refs: after sorting "- ErrorRefOnly" stays same
	// Also handle "- getVerifyRwxInternalError" etc.
	if strings.HasPrefix(trimmed, "- ") {
		rest := strings.TrimPrefix(trimmed, "- ")
		// If the rest has no spaces, it's a single function name reference
		if !strings.Contains(rest, " ") {
			return true
		}
	}
	// Check if it starts with known function names (sorted tokens may reorder)
	if strings.Contains(trimmed, "ErrorRefOnly") && !strings.Contains(trimmed, "access") {
		return true
	}
	if strings.Contains(trimmed, "getVerifyRwxInternalError") && !strings.Contains(trimmed, "access") {
		return true
	}
	return false
}

func nonWhiteSortedLines(s string) []string {
	if s == "" {
		return []string{""}
	}

	lines := strings.Split(strings.TrimSpace(s), "\n")
	var filtered []string

	for _, line := range lines {
		// First check raw line
		if isStackTraceLine(line) {
			continue
		}
		tokens := strings.Fields(line)
		sort.Strings(tokens)
		sortedLine := strings.Join(tokens, " ")
		// Also check after normalization
		if isStackTraceNormalizedLine(sortedLine) {
			continue
		}
		filtered = append(filtered, sortedLine)
	}

	if len(filtered) == 0 {
		return []string{""}
	}

	sort.Strings(filtered)

	return filtered
}

func assertNonWhiteSortedEqual(
	t *testing.T,
	testCase coretestcases.CaseV1,
	caseIndex int,
	actualErr error,
) {
	actStr := ""
	if actualErr != nil {
		actStr = actualErr.Error()
	}

	// Fix: handle both string and []string ExpectedInput types
	// See issues/chmodhelpertests-type-assertion-panic.md
	var expectedStr string
	switch v := testCase.ExpectedInput.(type) {
	case []string:
		expectedStr = strings.Join(v, "\n")
	case string:
		expectedStr = v
	default:
		expectedStr = fmt.Sprintf("%v", testCase.ExpectedInput)
	}

	actNorm := nonWhiteSortedLines(actStr)
	expNorm := nonWhiteSortedLines(expectedStr)

	normalizedCase := testCase
	normalizedCase.ExpectedInput = expNorm

	normalizedCase.ShouldBeEqual(t, caseIndex, actNorm...)
}
