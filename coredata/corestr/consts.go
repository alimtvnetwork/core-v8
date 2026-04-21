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

package corestr

import "github.com/alimtvnetwork/core-v8/constants"

const (
	charCollectionDefaultCapacity               = constants.ArbitraryCapacity10
	emptyChar                              byte = 0
	defaultEachCollectionCapacity               = constants.ArbitraryCapacity30
	defaultHashsetItems                         = 10
	RegularCollectionEfficiencyLimit            = 1000
	DoubleLimit                                 = RegularCollectionEfficiencyLimit * 3
	commonJoiner                                = "\n - "
	NoElements                                  = " {No Element}"
	charCollectionMapLengthFormat               = "\n## Items of `%s`"
	charHashsetMapLengthFormat                  = charCollectionMapLengthFormat
	charCollectionMapSingleItemFormat           = "\n\t%d . `%s` has `%d` items."
	charHashsetMapSingleItemFormat              = charCollectionMapSingleItemFormat
	summaryOfCharCollectionMapLengthFormat      = "# Summary of `%T`, Length (\"%d\") - Sequence `%d`"
	summaryOfCharHashsetMapLengthFormat         = summaryOfCharCollectionMapLengthFormat
	nodesCannotBeNull                           = "node cannot be nil."
	currentNodeCannotBeNull                     = "CurrentNode cannot be nil."
	keyValuePrintFormat                         = "{%s:%s}"
	linkedListCollectionCompareHeaderLeft       = "\n=================================================\n" +
		"-------------------------------------------------\n" +
		"# %s - Length : %d\n" +
		"%s\n" +
		"\n"
	linkedListCollectionCompareHeaderRight = "\n-------------------------------------------------\n" +
		"# %s - Length : %d\n" +
		"%s\n" +
		"----------------\n" +
		"Is Equals: %+v\n" +
		"Length: %d | %d\n"
	expectedLeftRightLength = constants.Two
)
