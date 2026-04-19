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

package corejson

import (
	"github.com/alimtvnetwork/core/constants"
)

type newResultsPtrCollectionCreator struct{}

// UnmarshalUsingBytes
//
//	Aka. alias for DeserializeUsingBytes
//
//	Should be used when ResultsPtrCollection itself is Serialized
//	and save to somewhere and then unmarshal or deserialize
func (it newResultsPtrCollectionCreator) UnmarshalUsingBytes(
	deserializingBytes []byte,
) (*ResultsPtrCollection, error) {
	return it.DeserializeUsingBytes(deserializingBytes)
}

// DeserializeUsingBytes
//
//	Should be used when ResultsPtrCollection itself is Serialized
//	and save to somewhere and then unmarshal or deserialize
func (it newResultsPtrCollectionCreator) DeserializeUsingBytes(
	deserializingBytes []byte,
) (*ResultsPtrCollection, error) {
	empty := it.Empty()

	err := Deserialize.
		UsingBytes(deserializingBytes, empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

func (it newResultsPtrCollectionCreator) DeserializeUsingResult(
	jsonResult *Result,
) (*ResultsPtrCollection, error) {
	if jsonResult.HasIssuesOrEmpty() {
		return nil, jsonResult.MeaningfulError()
	}

	empty := it.Empty()

	err := Deserialize.
		UsingBytes(jsonResult.SafeBytes(), empty)

	if err == nil {
		return empty, nil
	}

	return nil, err
}

func (it newResultsPtrCollectionCreator) Empty() *ResultsPtrCollection {
	list := make([]*Result, 0)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) UsingCap(
	capacity int,
) *ResultsPtrCollection {
	list := make([]*Result, 0, capacity)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) Default() *ResultsPtrCollection {
	list := make([]*Result, 0, constants.Capacity8)

	return &ResultsPtrCollection{
		Items: list,
	}
}

func (it newResultsPtrCollectionCreator) AnyItemsPlusCap(
	capacity int,
	anyItems ...any,
) *ResultsPtrCollection {
	length := capacity + len(anyItems)

	if length == 0 || len(anyItems) == 0 {
		return it.UsingCap(length)
	}

	collection := it.UsingCap(length)

	return collection.AddAnyItems(
		anyItems...)
}

func (it newResultsPtrCollectionCreator) AnyItems(
	anyItems ...any,
) *ResultsPtrCollection {
	return it.AnyItemsPlusCap(
		0,
		anyItems...)
}

func (it newResultsPtrCollectionCreator) UsingResultsPlusCap(
	addCapacity int,
	results ...*Result,
) *ResultsPtrCollection {
	length := addCapacity + len(results)

	if length == 0 || len(results) == 0 {
		return it.UsingCap(length)
	}

	list := it.UsingCap(length)

	return list.
		AddNonNilItemsPtr(results...)
}

func (it newResultsPtrCollectionCreator) UsingResults(
	results ...*Result,
) *ResultsPtrCollection {
	return it.UsingResultsPlusCap(
		0,
		results...)
}

func (it newResultsPtrCollectionCreator) JsonersPlusCap(
	isIgnoreNilOrErr bool,
	capacity int,
	jsoners ...Jsoner,
) *ResultsPtrCollection {
	length := capacity + len(jsoners)

	if length == 0 || len(jsoners) == 0 {
		return it.UsingCap(length)
	}

	collection := it.UsingCap(length)

	return collection.AddJsoners(
		isIgnoreNilOrErr,
		jsoners...)
}

func (it newResultsPtrCollectionCreator) Jsoners(
	jsoners ...Jsoner,
) *ResultsPtrCollection {
	return it.JsonersPlusCap(
		true,
		0,
		jsoners...)
}

func (it newResultsPtrCollectionCreator) Serializers(
	serializers ...bytesSerializer,
) *ResultsPtrCollection {
	if len(serializers) == 0 {
		return it.Empty()
	}

	collection := it.UsingCap(len(serializers))

	for _, serializer := range serializers {
		collection.AddSerializer(serializer)
	}

	return collection
}
