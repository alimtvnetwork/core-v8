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

package corepubsubinf

import (
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/enuminf"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core/coreinterface/loggerinf"
	"github.com/alimtvnetwork/core/coreinterface/pathextendinf"
	"github.com/alimtvnetwork/core/internal/internalinterface"
)

type IdAsStringer interface {
	internalinterface.IdAsStringer
}

type SubscriptionMainRecorder interface {
	IdAsStringer
	TableName() string

	IsEmpty() bool

	pathextendinf.PathExtenderGetter

	HasRecordError() bool
	SetRecordError() bool
	IsArchivedRecord() bool
	IsCompletedRecord() bool
	IsMigratedRecord() bool
	CompletionTyper() enuminf.CompletionStateTyper

	// DefaultDelayMillis
	//
	//  Chmod delay in milliseconds
	DefaultDelayMillis() int
}

type BaseLogModeler interface {
	enuminf.LoggerTyperGetter
	enuminf.EventTyperGetter
	errcoreinf.BasicErrorTyperGetter
	errcoreinf.ErrorStringGetter
	coreinterface.StackTracesBytesGetter
	coreinterface.JsonErrorBytesGetter
	IsEmpty() bool
	LogMessage() string
}

type CommunicateModeler interface {
	BaseLogModeler() BaseLogModeler
	PersistentId() uint
	IdAsStringer
	TableName() string

	SetCallerFileLineUsingStackSkip(
		stackSkip int,
	)

	loggerinf.SingleLogModeler
}

type SubscriptionRecorder interface {
	MainRecord() SubscriptionMainRecorder
	CommunicateRecord() CommunicateModeler
}
