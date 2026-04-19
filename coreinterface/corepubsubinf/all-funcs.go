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
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/coreinterface"
	"github.com/alimtvnetwork/core/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core/coreinterface/loggerinf"
	"github.com/alimtvnetwork/core/coreinterface/serializerinf"
)

type (
	SubscribeFunc func(
		subscriptionRecorder SubscriptionRecorder,
	)

	DirectSingleLogModelerSubscribeFunc func(modeler loggerinf.SingleLogModeler)

	JsonResultSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		jsonResult *corejson.Result,
	)

	BooleanSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		isResult bool,
	)

	HashmapSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		hashmap map[string]string,
	)
	BytesSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		rawBytes []byte,
	)
	ModelJsonSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		modelJson []byte,
	)
	AnyItemSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		anyItem any,
	)

	StringSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		message string,
	)

	SimpleBytesResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.SimpleBytesResulter,
	)

	BaseJsonResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.BaseJsonResulter,
	)

	JsonResulterSubscribeFunc func(
		categoryRevealer coreinterface.CategoryRevealer,
		result serializerinf.JsonResulter,
	)

	DirectBytesSubscribeFunc func(
		rawBytes []byte,
	)

	DirectModelJsonSubscribeFunc func(
		modelJson []byte,
	)

	DirectAnyItemSubscribeFunc func(
		anyItem any,
	)

	DirectJsonResultSubscribeFunc func(
		jsonResult *corejson.Result,
	)

	DirectStringSubscribeFunc func(
		message string,
	)

	DirectBasicErrorSubscribeFunc func(
		basicErrorWrapper errcoreinf.BasicErrWrapper,
	)

	DirectBaseErrorOrCollectionWrapperSubscribeFunc func(
		basicErrorWrapper errcoreinf.BaseErrorOrCollectionWrapper,
	)

	DirectBooleanSubscribeFunc func(
		isResult bool,
	)

	DirectSimpleBytesResulterSubscribeFunc func(
		result serializerinf.SimpleBytesResulter,
	)

	DirectBaseJsonResulterSubscribeFunc func(
		result serializerinf.BaseJsonResulter,
	)

	DirectJsonResulterSubscribeFunc func(
		result serializerinf.JsonResulter,
	)

	SimpleSubscribeFunc                 func(communicate CommunicateModeler)
	FilterStringSubscribeFunc           func(communicate CommunicateModeler, currentStringValue string)
	LogSubscriberFunc                   func(logRecord BaseLogModeler)
	StartFunc                           func(subscriptionRecorder SubscriptionRecorder)
	CompletionFunc                      func(subscriptionRecorder SubscriptionRecorder)
	StartEndSubscriptionFunc            func(isStart bool, subscriptionRecorder SubscriptionRecorder)
	SimpleCompletionFunc                func(communicate CommunicateModeler)
	CategoryNameAnyItemSubscriptionFunc func(categoryName string, anyItem any)
	FilterAnyItemSubscriptionFunc       func(filter string, anyItem any)
	FilterBytesSubscriptionFunc         func(filter string, rawBytes []byte)
	MatcherFunc                         func() (isMatch bool)
)
