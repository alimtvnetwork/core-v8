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

package corestr

import (
	"github.com/alimtvnetwork/core/constants"
)

type newCollectionsOfCollectionCreator struct{}

func (it *newCollectionsOfCollectionCreator) Cap(
	capacity int,
) *CollectionsOfCollection {
	collection := make([]*Collection, 0, capacity)

	return &CollectionsOfCollection{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionCreator) Empty() *CollectionsOfCollection {
	collection := make([]*Collection, constants.Zero)

	return &CollectionsOfCollection{
		items: collection,
	}
}

func (it *newCollectionsOfCollectionCreator) StringsOfStrings(
	isMakeClone bool,
	stringItems ...[]string,
) *CollectionsOfCollection {
	length := len(stringItems)

	return it.LenCap(
		constants.Zero,
		length,
	).AddsStringsOfStrings(isMakeClone, stringItems...)
}

func (it *newCollectionsOfCollectionCreator) SpreadStrings(
	isMakeClone bool,
	stringItems ...string,
) *CollectionsOfCollection {
	length := len(
		stringItems,
	)

	return it.LenCap(
		constants.Zero,
		length,
	).AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) CloneStrings(
	stringItems []string,
) *CollectionsOfCollection {
	return it.StringsOption(
		true,
		0,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) Strings(
	stringItems []string,
) *CollectionsOfCollection {
	length := len(
		stringItems,
	)
	collection := it.Cap(
		length,
	)

	return collection.AddStrings(
		false,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) StringsOption(
	isMakeClone bool,
	capacity int,
	stringItems []string,
) *CollectionsOfCollection {
	length := len(stringItems)
	collection := it.Cap(
		length + capacity,
	)

	return collection.AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) StringsOptions(
	isMakeClone bool,
	capacity int,
	stringItems []string,
) *CollectionsOfCollection {
	length := len(stringItems)
	collection := it.Cap(length + capacity)

	return collection.AddStrings(
		isMakeClone,
		stringItems,
	)
}

func (it *newCollectionsOfCollectionCreator) LenCap(
	length,
	capacity int,
) *CollectionsOfCollection {
	collection := make(
		[]*Collection,
		length,
		capacity,
	)

	return &CollectionsOfCollection{
		items: collection,
	}
}
