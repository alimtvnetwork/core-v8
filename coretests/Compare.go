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

package coretests

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/alimtvnetwork/core-v8/constants"
)

type Compare struct {
	sortedString   *string
	sortedStrings  []string
	MatchingLength int // 0/-1 means all, number means, that number must match
	StringContains string
}

func (it *Compare) SortedStrings() []string {
	if it.sortedStrings != nil {
		return it.sortedStrings
	}

	it.sortedStrings = GetAssert.SortedArray(
		false,
		true,
		strings.TrimSpace(it.StringContains),
	)

	return it.sortedStrings
}

func (it *Compare) SortedString() string {
	if it.sortedString != nil {
		return *it.sortedString
	}

	sortedStrings := it.SortedStrings()
	sortedString := strings.Join(
		sortedStrings,
		commonJoiner,
	)

	it.sortedString = &sortedString

	return *it.sortedString
}

func (it *Compare) GetPrintMessage(index int) string {
	return fmt.Sprintf(
		"\n\tIndex:%d\n\tString IsContains:%s\n\tString Processed:%s",
		index,
		it.StringContains,
		it.SortedString(),
	)
}

func (it *Compare) IsMatch(
	isPrint bool,
	index int,
	instruction *ComparingInstruction,
) bool {
	actualHashset := instruction.ActualHashset()
	sortedStrings := it.SortedStrings()

	// all
	if it.MatchingLength <= constants.Zero {
		isMatch := actualHashset.HasAll(sortedStrings...)

		if !isMatch && isPrint {
			compiledMessage := it.GetPrintMessage(index)

			slog.Warn("compare mismatch", "message", compiledMessage)
		}

		return isMatch
	}

	foundMatches := 0

	for _, item := range sortedStrings {
		if actualHashset.Has(item) {
			foundMatches++
		}
	}

	isMatch := foundMatches >= it.MatchingLength
	if !isMatch && isPrint {
		compiledMessage := it.GetPrintMessage(index)

		slog.Warn("compare mismatch", "message", compiledMessage)
	}

	return isMatch
}
