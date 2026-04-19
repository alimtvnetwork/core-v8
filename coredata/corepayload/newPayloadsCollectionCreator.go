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
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/errcore"
)

type newPayloadsCollectionCreator struct{}

func (it newPayloadsCollectionCreator) Empty() *PayloadsCollection {
	return &PayloadsCollection{
		Items: []*PayloadWrapper{},
	}
}

func (it newPayloadsCollectionCreator) Deserialize(
	rawBytes []byte,
) (*PayloadsCollection, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		UsingBytes(
			rawBytes, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadsCollectionCreator) DeserializeMust(
	rawBytes []byte,
) *PayloadsCollection {
	collection, err := it.Deserialize(rawBytes)
	errcore.HandleErr(err)

	return collection
}

func (it newPayloadsCollectionCreator) DeserializeToMany(
	rawBytes []byte,
) (payloadsSlice []*PayloadsCollection, err error) {
	err = corejson.
		Deserialize.
		UsingBytes(
			rawBytes,
			&payloadsSlice)

	if err != nil {
		return nil, err
	}

	return payloadsSlice, nil
}

func (it newPayloadsCollectionCreator) DeserializeUsingJsonResult(
	jsonResult *corejson.Result,
) (*PayloadsCollection, error) {
	empty := it.Empty()

	err := corejson.
		Deserialize.
		Apply(jsonResult, empty)

	if err != nil {
		return nil, err
	}

	return empty, nil
}

func (it newPayloadsCollectionCreator) UsingCap(
	capacity int,
) *PayloadsCollection {
	return &PayloadsCollection{
		Items: make([]*PayloadWrapper, 0, capacity),
	}
}

func (it newPayloadsCollectionCreator) UsingWrappers(
	payloadsWrappers ...*PayloadWrapper,
) *PayloadsCollection {
	if len(payloadsWrappers) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(
		len(payloadsWrappers) +
			constants.Capacity3)

	return collection.AddsPtr(payloadsWrappers...)
}
