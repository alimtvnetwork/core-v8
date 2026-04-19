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

package internalinterface

type CountStateTracker interface {
	IsSameStateUsingCount(
		currentCount int,
	) bool
	IsSameState() bool
	IsSuccessValidator
	HasChangesChecker
}

type HasChangesChecker interface {
	HasChanges() bool
}

type DynamicDiffChangesGetter interface {
	DynamicDiffChanges() ([]byte, error)
}

type DynamicChangeStateDetector interface {
	IsDynamicEqual(rawBytes []byte) bool
	DynamicDiffChangesGetter
	HasChangesChecker
	IsSuccessValidator
	ChangesLogger
	MustChangesLogger
}

type ChangesLogger interface {
	LogChanges() error
}

type MustChangesLogger interface {
	LogChangesMust()
}

type AcceptRejectOrSkipper interface {
	IsAccept() bool
	IsReject() bool
	IsSkip() bool
}

type YesNoAsker interface {
	IsYes() bool
	IsNo() bool
	IsAsk() bool
	IsIndeterminate() bool
}

type YesNoAcceptRejecter interface {
	YesNoAsker
	AcceptRejectOrSkipper
}
