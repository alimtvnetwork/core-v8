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

type newCharCollectionMapCreator struct{}

// CapSelfCap CharHashsetMap.eachCollectionCapacity,
// capacity minimum 10 will be set if lower than 10 is given.
//
// For lower than 5 use the Empty items definition.
func (it *newCharCollectionMapCreator) CapSelfCap(
	capacity, selfCollectionCapacity int,
) *CharCollectionMap {
	if capacity < charCollectionDefaultCapacity {
		capacity = charCollectionDefaultCapacity
	}

	mapElements := make(map[byte]*Collection, capacity)

	if selfCollectionCapacity < charCollectionDefaultCapacity {
		selfCollectionCapacity = charCollectionDefaultCapacity
	}

	return &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: selfCollectionCapacity,
	}
}

// Empty eachCollectionCapacity = 0
func (it *newCharCollectionMapCreator) Empty() *CharCollectionMap {
	mapElements := make(map[byte]*Collection, constants.Zero)

	return &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: defaultEachCollectionCapacity,
	}
}

func (it *newCharCollectionMapCreator) Items(
	items []string,
) *CharCollectionMap {
	if len(items) == 0 {
		return it.Empty()
	}

	mapElements := make(map[byte]*Collection, len(items))
	charCollectionMap := &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: constants.Zero,
	}

	charCollectionMap.AddStrings(items...)

	return charCollectionMap
}

func (it *newCharCollectionMapCreator) ItemsPtrWithCap(
	additionalCapacity int,
	eachCollectionCap int,
	items []string,
) *CharCollectionMap {
	length := len(items)
	isDefined := length > 0
	if isDefined {
		additionalCapacity += length
	}

	mapElements := make(
		map[byte]*Collection,
		additionalCapacity,
	)

	charCollectionMap := &CharCollectionMap{
		items:                  mapElements,
		eachCollectionCapacity: eachCollectionCap,
	}

	if !isDefined || length == 0 {
		return charCollectionMap
	}

	return charCollectionMap.AddStrings(items...)
}
