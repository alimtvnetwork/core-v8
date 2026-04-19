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

package codestack

import "github.com/alimtvnetwork/core/constants"

type stacksTo struct{}

func (it stacksTo) Bytes(stackSkipIndex int) []byte {
	return New.
		StackTrace.
		DefaultCount(stackSkipIndex + defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) BytesDefault() []byte {
	return New.
		StackTrace.
		DefaultCount(defaultInternalSkip).
		StackTracesBytes()
}

func (it stacksTo) String(
	startSkipIndex, count int,
) string {
	stacks := it.newStacks(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func (it stacksTo) newStacks(startSkipIndex int, count int) TraceCollection {
	stacks := New.
		StackTrace.
		Default(startSkipIndex+defaultInternalSkip, count)

	return stacks
}

func (it stacksTo) newStacksDefault(startSkipIndex int) TraceCollection {
	stacks := New.
		StackTrace.
		DefaultCount(startSkipIndex + defaultInternalSkip)

	return stacks
}

func (it stacksTo) StringUsingFmt(
	formatter Formatter,
	startSkipIndex, count int,
) string {
	stacks := it.newStacks(
		startSkipIndex+defaultInternalSkip,
		count,
	)

	toString := stacks.JoinUsingFmt(
		formatter,
		constants.NewLineSpaceHyphenSpace,
	)
	stacks.Dispose()

	return toString
}

func (it stacksTo) JsonString(
	startSkipIndex int,
) string {
	stacks := it.newStacksDefault(
		startSkipIndex + defaultInternalSkip,
	)

	json := stacks.JsonPtr()
	stacks.Dispose()
	json.HandleError()

	return json.JsonString()
}

func (it stacksTo) JsonStringDefault() string {
	return it.JsonString(defaultInternalSkip)
}

func (it stacksTo) StringNoCount(
	startSkipIndex int,
) string {
	stacks := it.newStacksDefault(
		startSkipIndex + defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}

func (it stacksTo) StringDefault() string {
	stacks := it.newStacksDefault(
		defaultInternalSkip,
	)

	toString := stacks.CodeStacksString()
	stacks.Dispose()

	return toString
}
