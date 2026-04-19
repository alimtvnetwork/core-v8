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
	"strings"
)

const mapSeparator = "============================>"

// MapMismatchError builds a diagnostic error for map assertion failures.
//
// Each map entry is shown on its own line with tab indentation in Go literal
// format, making the output directly copy-pasteable into _testcases.go.
//
// Output format:
//
//	Test Method : TestFuncName
//	Case        : 1
//	Title       : Case Title
//
//	============================>
//	1) Actual Received (2 entries):
//	    Case Title
//	============================>
//		"containsName": false,
//		"hasError":      false,
//	============================>
//
//	============================>
//	1) Expected Input (1 entries):
//	    Case Title
//	============================>
//		"hasError": false,
//	============================>
func MapMismatchError(
	testName string,
	caseIndex int,
	title string,
	actualGoLiteralLines []string,
	expectedGoLiteralLines []string,
) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("\n    Test Method : %s\n", testName))
	sb.WriteString(fmt.Sprintf("    Case        : %d\n", caseIndex))
	sb.WriteString(fmt.Sprintf("    Title       : %s\n\n", title))

	// Actual block
	sb.WriteString(mapSeparator + "\n")
	sb.WriteString(fmt.Sprintf(
		"%d) Actual Received (%d entries):\n",
		caseIndex,
		len(actualGoLiteralLines),
	))
	sb.WriteString(fmt.Sprintf("    %s\n", title))
	sb.WriteString(mapSeparator + "\n")

	for _, line := range actualGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	sb.WriteString(mapSeparator + "\n")

	// Expected block
	sb.WriteString("\n")
	sb.WriteString(mapSeparator + "\n")
	sb.WriteString(fmt.Sprintf(
		"%d) Expected Input (%d entries):\n",
		caseIndex,
		len(expectedGoLiteralLines),
	))
	sb.WriteString(fmt.Sprintf("    %s\n", title))
	sb.WriteString(mapSeparator + "\n")

	for _, line := range expectedGoLiteralLines {
		sb.WriteString("\t")
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	sb.WriteString(mapSeparator)

	return sb.String()
}
