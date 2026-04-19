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

import "github.com/alimtvnetwork/core/constants"

type newHashsetsCollectionCreator struct{}

func (it *newHashsetsCollectionCreator) Empty() *HashsetsCollection {
	collection := make(
		[]*Hashset,
		constants.Zero,
		constants.Zero)

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) UsingHashsets(
	hashsets ...Hashset,
) *HashsetsCollection {
	length := len(hashsets)

	if length == 0 {
		return it.Empty()
	}

	collection := make(
		[]*Hashset,
		length,
		length+constants.ArbitraryCapacity10)

	for i := range hashsets {
		collection[i] = &hashsets[i]
	}

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) UsingHashsetsPointers(
	hashsets ...*Hashset,
) *HashsetsCollection {
	if len(hashsets) == 0 {
		return it.Empty()
	}

	return &HashsetsCollection{
		items: hashsets,
	}
}

func (it *newHashsetsCollectionCreator) LenCap(
	length, capacity int,
) *HashsetsCollection {
	collection := make([]*Hashset, length, capacity)

	return &HashsetsCollection{
		items: collection,
	}
}

func (it *newHashsetsCollectionCreator) Cap(
	capacity int,
) *HashsetsCollection {
	collection := make([]*Hashset, constants.Zero, capacity)

	return &HashsetsCollection{
		items: collection,
	}
}
