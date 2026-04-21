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

package corefuncs

import (
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
	"github.com/alimtvnetwork/core-v8/coredata/corepayload"
	"github.com/alimtvnetwork/core-v8/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v8/coreinterface/serializerinf"
)

type (
	ExecFunc                    func()
	StringerActionFunc          func() (result string)
	StringerWithErrorActionFunc func() (result string, err error)
	ActionFunc                  func()
	IsApplyFunc                 func() (isSuccess bool)
	IsBooleanFunc               func() bool
	InOutFunc                   func(input any) (output any)
	InOutErrFunc                func(input any) (output any, err error)
	SerializeOutputFunc         func(input any) (serializedBytes []byte, err error)
	SerializerVoidFunc          func() (serializedBytes []byte, err error)
	InActionReturnsErrFunc      func(input any) (err error)
	NamedActionFunc             func(name string)
	ActionReturnsErrorFunc      func() error
	IsSuccessFunc               func() (isSuccess bool)
	IsFailureFunc               func() (isFailed bool)
	// ResultDelegatingFunc
	//
	// resultDelegatedTo can be unmarshal or marshal or reflect set
	ResultDelegatingFunc           func(resultDelegatedTo any) error
	NextReturnErrWrapperFunc       func(nextAction ActionReturnsErrorFunc) error
	NextVoidActionFunc             func(nextAction ExecFunc)
	PayloadProcessorFunc           func(payloads []byte) (err error)
	PayloadToBasicErrProcessorFunc func(payloads []byte) (basicError errcoreinf.BasicErrWrapper)
	SimpleBytesResultProcessorFunc func(simpleBytes serializerinf.SimpleBytesResulter) (basicError errcoreinf.BasicErrWrapper)
	ErrorToBasicError              func(
		errorTyper errcoreinf.BaseErrorTyper,
		err error,
	) (basicError errcoreinf.BasicErrWrapper)
	BaseJsonResultProcessorFunc          func(baseJsonResulter serializerinf.BaseJsonResulter) (basicError errcoreinf.BasicErrWrapper)
	JsonResulterProcessorFunc            func(result serializerinf.JsonResulter) (basicError errcoreinf.BasicErrWrapper)
	JsonResultProcessorFunc              func(result *corejson.Result) (basicError errcoreinf.BasicErrWrapper)
	PayloadWrapperProcessorFunc          func(payloadWrapper *corepayload.PayloadWrapper) (basicError errcoreinf.BasicErrWrapper)
	MultiPayloadsProcessorFunc           func(multiPayloads ...[]byte) (err error)
	BytesCollectionPayloadsProcessorFunc func(collectionOfBytes *corejson.BytesCollection) (err error)
	PayloadToPayloadWrapperFunc          func(payloads []byte) (payloadWrapper *corepayload.PayloadWrapper, err error)
	NextPayloadProcessorLinkerFunc       func(nextLinkerFunc PayloadProcessorFunc) error
)
