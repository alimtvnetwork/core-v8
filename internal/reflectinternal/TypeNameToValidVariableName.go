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

package reflectinternal

import (
	"strings"
	"unicode"

	"github.com/alimtvnetwork/core-v8/coredata/stringslice"
)

func TypeNameToValidVariableName(currentTypeName string) (validVarName string) {
	if len(currentTypeName) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.Grow(len(currentTypeName))
	splitsByDot := strings.Split(currentTypeName, ".")
	hasDot := len(splitsByDot) > 1

	if hasDot {
		var leftPartSb strings.Builder
		firstPart := splitsByDot[0]

		for i, c := range firstPart {
			if unicode.IsSpace(c) {
				continue
			}

			if c == '[' && firstPart[i+1] == ']' {
				i++
				leftPartSb.WriteString(typeReplacerMap["[]"])
				continue
			}

			if c == '{' && firstPart[i+1] == '}' {
				i++
				continue
			}

			_, has := typeReplacerMap[string(c)]
			if has {
				leftPartSb.WriteString(typeReplacerMap[string(c)])

				continue
			}
		}

		currentTypeName = stringslice.Joins(
			"",
			leftPartSb.String(),
			splitsByDot[len(splitsByDot)-1],
		)
	}

	for i := 0; i < len(currentTypeName); i++ {
		currentChar := currentTypeName[i]

		if unicode.IsSpace(rune(currentChar)) {
			continue
		}

		if currentChar == '[' && currentTypeName[i+1] == ']' {
			i++
			sb.WriteString(typeReplacerMap["[]"])

			continue
		}

		if currentChar == '{' && currentTypeName[i+1] == '}' {
			i++
			continue
		}

		replace, has := typeReplacerMap[string(currentChar)]

		if has {
			sb.WriteString(typeReplacerMap[replace])

			continue
		}

		sb.WriteByte(currentChar)
	}

	return sb.String()
}
