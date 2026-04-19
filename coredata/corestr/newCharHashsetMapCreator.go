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

type newCharHashsetMapCreator struct{}

// Cap
//
// CharHashsetMap.eachHashsetCapacity,
// capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the Empty hashset definition.
func (it *newCharHashsetMapCreator) Cap(
	capacity, selfHashsetCapacity int,
) *CharHashsetMap {
	const limit = constants.ArbitraryCapacity10

	if capacity < limit {
		capacity = limit
	}

	mapElements := make(
		map[byte]*Hashset,
		capacity,
	)

	if selfHashsetCapacity < limit {
		selfHashsetCapacity = limit
	}

	return &CharHashsetMap{
		items:               mapElements,
		eachHashsetCapacity: selfHashsetCapacity,
	}
}

func (it *newCharHashsetMapCreator) CapItems(
	capacity, selfHashsetCapacity int,
	items ...string,
) *CharHashsetMap {
	charHashsetMap := it.Cap(
		capacity, selfHashsetCapacity,
	)

	return charHashsetMap.AddStrings(items...)
}

func (it *newCharHashsetMapCreator) Strings(
	selfHashsetCapacity int,
	items []string,
) *CharHashsetMap {
	if items == nil {
		return it.Cap(
			constants.ArbitraryCapacity5,
			selfHashsetCapacity,
		)
	}

	length := len(items)
	charHashsetMap := it.Cap(
		length,
		selfHashsetCapacity,
	)

	charHashsetMap.AddStrings(items...)

	return charHashsetMap
}
