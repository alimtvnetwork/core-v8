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

package errcore

import (
	"fmt"
)

// PrintDiffOnMismatch prints a formatted diff diagnostic block
// only when actual and expected lines differ.
//
// It prints:
//   - A header with case index and title
//   - Optional context lines (e.g., "  InitValue: hello")
//   - The standard line-by-line diff via PrintLineDiff
//   - A footer closing the block
//
// contextLines are printed as-is between the header and the diff.
// Each context line should be pre-formatted (e.g., fmt.Sprintf("  Key: %v", val)).
func PrintDiffOnMismatch(
	caseIndex int,
	title string,
	actLines []string,
	expectedLines []string,
	contextLines ...string,
) {
	if !HasAnyMismatchOnLines(actLines, expectedLines) {
		return
	}

	header := fmt.Sprintf("\n=== Diff (Case %d: %s) ===\n", caseIndex, title)
	fmt.Print(header)

	for _, cl := range contextLines {
		fmt.Println(cl)
	}

	PrintLineDiff(caseIndex, title, actLines, expectedLines)

}
