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

package corepayload

import (
	"github.com/alimtvnetwork/core-v8/coredata/coredynamic"
	"github.com/alimtvnetwork/core-v8/coreinterface/errcoreinf"
	"github.com/alimtvnetwork/core-v8/coreinterface/payloadinf"
	"github.com/alimtvnetwork/core-v8/errcore"
)

// AttributesSetters.go — Mutating methods extracted from Attributes.go

func (it *Attributes) HandleErr() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) HandleError() {
	if it.HasError() {
		it.BasicErrWrapper.HandleError()
	}
}

func (it *Attributes) MustBeEmptyError() {
	if it.IsEmptyError() {
		return
	}

	panic(it.Error())
}

// SetAuthInfo
//
//	On nil create new attributes
func (it *Attributes) SetAuthInfo(authInfo *AuthInfo) *Attributes {
	if it == nil {
		return New.
			Attributes.
			UsingAuthInfo(authInfo)
	}

	it.AuthInfo = authInfo

	return it
}

func (it *Attributes) SetUserInfo(
	userInfo *UserInfo,
) *Attributes {
	if it == nil {
		return &Attributes{
			AuthInfo: &AuthInfo{
				UserInfo: userInfo,
			},
		}
	}

	it.AuthInfo.SetUserInfo(userInfo)

	return it
}

func (it *Attributes) AddNewStringKeyValueOnly(key, value string) (isAdded bool) {
	if it == nil || it.KeyValuePairs == nil {
		return false
	}

	it.KeyValuePairs.
		AddOrUpdate(key, value)

	return true
}

func (it *Attributes) AddNewAnyKeyValueOnly(
	key string, value any,
) (isAdded bool) {
	if it == nil || it.AnyKeyValuePairs == nil {
		return false
	}

	it.AnyKeyValuePairs.Add(key, value)

	return true
}

func (it *Attributes) AnyKeyReflectSetTo(
	key string,
	toPtr any,
) error {
	if it == nil || it.KeyValuePairs == nil {
		return errcore.
			CannotBeNilOrEmptyType.ErrorNoRefs(
			"KeyValuePairs is nil")
	}

	return it.AnyKeyValuePairs.ReflectSetTo(key, toPtr)
}

func (it *Attributes) ReflectSetTo(
	toPointer any,
) error {
	return coredynamic.ReflectSetFromTo(it, toPointer)
}

func (it *Attributes) AddOrUpdateString(
	key, value string,
) (isNewlyAdded bool) {
	if it == nil || it.KeyValuePairs == nil {
		return false
	}

	return it.
		KeyValuePairs.
		AddOrUpdate(key, value)
}

func (it *Attributes) AddOrUpdateAnyItem(
	key string,
	anyItem any,
) (isNewlyAdded bool) {
	if it == nil || it.AnyKeyValuePairs == nil {
		return false
	}

	return it.
		AnyKeyValuePairs.
		Add(key, anyItem)
}

// SetBasicErr
//
//	on nil creates and attach new error and returns the attributes
func (it *Attributes) SetBasicErr(
	basicErr errcoreinf.BasicErrWrapper,
) payloadinf.AttributesBinder {
	if it == nil {
		return New.Attributes.UsingBasicError(
			basicErr)
	}

	it.BasicErrWrapper = basicErr

	return it
}

func (it *Attributes) Clear() {
	if it == nil {
		return
	}

	it.KeyValuePairs.Clear()
	it.AnyKeyValuePairs.Clear()
	it.DynamicPayloads = []byte{}
}

func (it *Attributes) Dispose() {
	if it == nil {
		return
	}

	it.Clear()
}
