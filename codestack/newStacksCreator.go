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

type newStacksCreator struct{}

func (it newStacksCreator) All(
	isSkipInvalid,
	isBreakOnceInvalid bool,
	startSkipIndex, // should start from 1
	stackCount int,
) TraceCollection {
	traces := New.traces.Cap(stackCount + constants.Capacity2)

	return *traces.AddsUsingSkip(
		isSkipInvalid,
		isBreakOnceInvalid,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func (it newStacksCreator) Default(
	startSkipIndex,
	stackCount int,
) TraceCollection {
	return it.All(
		true,
		true,
		startSkipIndex+defaultInternalSkip,
		stackCount,
	)
}

func (it newStacksCreator) DefaultCount(
	startSkipIndex int,
) TraceCollection {
	return it.All(
		true,
		true,
		startSkipIndex,
		DefaultStackCount,
	)
}

func (it newStacksCreator) SkipOne() TraceCollection {
	return it.All(
		true,
		true,
		Skip1+defaultInternalSkip,
		DefaultStackCount,
	)
}

func (it newStacksCreator) SkipNone() TraceCollection {
	return it.All(
		true,
		true,
		defaultInternalSkip,
		DefaultStackCount,
	)
}
