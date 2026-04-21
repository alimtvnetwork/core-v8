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

import "github.com/alimtvnetwork/core-v8/constants"

type newKeyValuesCreator struct{}

func (it *newKeyValuesCreator) Cap(
	capacity int,
) *KeyValueCollection {
	collection := make(
		[]KeyValuePair,
		constants.Zero,
		capacity,
	)

	return &KeyValueCollection{
		KeyValuePairs: collection,
	}
}

func (it *newKeyValuesCreator) Empty() *KeyValueCollection {
	collection := make(
		[]KeyValuePair,
		constants.Zero,
	)

	return &KeyValueCollection{
		KeyValuePairs: collection,
	}
}

func (it *newKeyValuesCreator) UsingMap(
	input map[string]string,
) *KeyValueCollection {
	length := len(input)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	return keyValCollection.AddMap(input)
}

func (it *newKeyValuesCreator) UsingKeyValuePairs(
	keyValPairs ...KeyValuePair,
) *KeyValueCollection {
	length := len(keyValPairs)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	return keyValCollection.Adds(keyValPairs...)
}

func (it *newKeyValuesCreator) UsingKeyValueStrings(
	keys, values []string,
) *KeyValueCollection {
	length := len(keys)
	keyValCollection := it.Cap(length)

	if length == 0 {
		return keyValCollection
	}

	for i, key := range keys {
		keyValCollection.Add(key, values[i])
	}

	return keyValCollection
}
