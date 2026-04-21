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

package coreinterface

import "github.com/alimtvnetwork/core-v8/internal/internalinterface"

type ChangesCommitter interface {
	HasChangesChecker
	ChangeAccepter
	ChangeRejecter
	RemindLaterChangeSkipper
	Commit(option AcceptRejectOrSkipper) error
	CommitMust(option AcceptRejectOrSkipper)
}

type ChangeAccepter interface {
	AcceptChanges() error
	AcceptChangesMust()
}

type ChangeRejecter interface {
	RejectChanges() error
	RejectChangesMust()
}

type RemindLaterChangeSkipper interface {
	SkipChangesRemindLater() error
	SkipChangesRemindLaterMust()
}

type CountStateTracker interface {
	internalinterface.CountStateTracker
}

type DynamicDiffChangesGetter interface {
	internalinterface.DynamicDiffChangesGetter
}

type HasChangesChecker interface {
	internalinterface.HasChangesChecker
}

type DynamicChangeStateDetector interface {
	internalinterface.DynamicChangeStateDetector
}

type ChangesLogger interface {
	internalinterface.ChangesLogger
}

type MustChangesLogger interface {
	internalinterface.MustChangesLogger
}

type YesNoAsker interface {
	internalinterface.YesNoAsker
}

type AcceptRejectOrSkipper interface {
	internalinterface.AcceptRejectOrSkipper
}

type YesNoAcceptRejecter interface {
	internalinterface.YesNoAcceptRejecter
}

type EnhanceYesNoAcceptRejecter interface {
	YesNoAcceptRejecter
	IsAcceptOrReject() bool
	IsNotAcceptOrReject() bool
	IsDefinedAccepted() bool
}

type IsReviewChecker interface {
	IsReview() bool
}
