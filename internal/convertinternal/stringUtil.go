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

package convertinternal

import (
	"fmt"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
)

type stringUtil struct{}

func (it stringUtil) PrependWithSpaces(
	joiner string,
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, joiner)
}

func (it stringUtil) PrependWithSpacesDefault(
	spaceCountLines int,
	existingLines []string,
	prependingLinesSpaceCount int,
	prependingLines ...string,
) string {
	toSlice := Util.Strings.PrependWithSpaces(
		spaceCountLines,
		existingLines,
		prependingLinesSpaceCount,
		prependingLines...,
	)

	return strings.Join(toSlice, constants.NewLineUnix)
}

func (it stringUtil) IndexToPosition(
	index int,
) string {
	position := index + 1

	switch position {
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	default:
		return fmt.Sprintf(
			"%dth",
			position,
		)
	}
}

func (it stringUtil) PascalCase(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	lines := strings.Split(name, "_")

	for i, s := range lines {
		lines[i] = it.pascalCaseWord(s)
	}

	return strings.Join(lines, "")
}

func (it stringUtil) pascalCaseWord(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	allRunes := []rune(name)
	firstChar := allRunes[0]
	firstCharStr := string(firstChar)
	firstCharUpper := strings.ToUpper(firstCharStr)

	if len(allRunes) == 1 {
		return firstCharUpper
	}

	return firstCharUpper + string(allRunes[1:])
}

func (it stringUtil) CamelCase(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	lines := strings.Split(name, "_")

	for i, s := range lines {
		if i == 0 {
			lines[i] = it.camelCaseWord(s)
		} else {
			lines[i] = it.pascalCaseWord(s)
		}
	}

	return strings.Join(lines, "")
}

func (it stringUtil) camelCaseWord(
	name string,
) string {
	if len(name) == 0 {
		return ""
	}

	allRunes := []rune(name)
	firstChar := allRunes[0]
	firstCharStr := string(firstChar)
	firstCharLower := strings.ToLower(firstCharStr)

	if len(allRunes) == 1 {
		return firstCharLower
	}

	return firstCharLower + string(allRunes[1:])
}
