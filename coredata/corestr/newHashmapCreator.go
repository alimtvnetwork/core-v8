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
	"github.com/alimtvnetwork/core/converters"
)

type newHashmapCreator struct{}

func (it *newHashmapCreator) Empty() *Hashmap {
	return it.Cap(constants.Zero)
}

func (it *newHashmapCreator) Cap(length int) *Hashmap {
	hashset := make(map[string]string, length)

	return &Hashmap{
		items:         hashset,
		hasMapUpdated: false,
	}
}

func (it *newHashmapCreator) KeyAnyValues(
	keyAnyValues ...KeyAnyValuePair,
) *Hashmap {
	if len(keyAnyValues) == 0 {
		return it.Cap(defaultHashsetItems)
	}

	length := len(keyAnyValues)
	hashMap := it.Cap(length + constants.ArbitraryCapacity10)

	return hashMap.AddOrUpdateKeyAnyValues(keyAnyValues...)
}

func (it *newHashmapCreator) KeyValues(
	keyValues ...KeyValuePair,
) *Hashmap {
	if len(keyValues) == 0 {
		return it.Cap(defaultHashsetItems)
	}

	length := len(keyValues)
	hashMap := it.Cap(length + constants.ArbitraryCapacity10)

	return hashMap.AddOrUpdateKeyValues(keyValues...)
}

func (it *newHashmapCreator) KeyValuesCollection(
	keys, values *Collection,
) *Hashmap {
	if keys == nil || keys.IsEmpty() {
		return it.Empty()
	}

	itemsMap := converters.KeyValuesTo.ToMap(
		keys.List(),
		values.List(),
	)

	return it.UsingMap(itemsMap)
}

func (it *newHashmapCreator) KeyValuesStrings(
	keys, values []string,
) *Hashmap {
	if len(keys) == 0 {
		return it.Empty()
	}

	itemsMap := converters.KeyValuesTo.ToMap(
		keys,
		values,
	)

	return it.UsingMap(itemsMap)
}

func (it *newHashmapCreator) UsingMap(
	itemsMap map[string]string,
) *Hashmap {
	return &Hashmap{
		items:         itemsMap,
		hasMapUpdated: true,
	}
}

// UsingMapOptions
// isMakeClone : copies itemsMap or else use the same one as pointer assign.
func (it *newHashmapCreator) UsingMapOptions(
	isMakeClone bool,
	addCapacity int,
	itemsMap map[string]string,
) *Hashmap {
	if len(itemsMap) == 0 {
		return it.Cap(addCapacity)
	}

	length := len(itemsMap)

	if isMakeClone {
		hashMap := it.Cap(length + addCapacity)

		return hashMap.AddOrUpdateMap(itemsMap)
	}

	// no clone
	return &Hashmap{
		items:         itemsMap,
		hasMapUpdated: true,
	}
}

// MapWithCap always returns the clone of the items.
func (it *newHashmapCreator) MapWithCap(
	addCapacity int,
	itemsMap map[string]string,
) *Hashmap {
	if len(itemsMap) == 0 {
		return it.Cap(addCapacity)
	}

	if addCapacity == 0 {
		return it.UsingMap(itemsMap)
	}

	length := len(itemsMap)
	hashMap := it.Cap(length + addCapacity)

	return hashMap.AddOrUpdateMap(itemsMap)
}
