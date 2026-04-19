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
	"strings"

	"github.com/alimtvnetwork/core/constants"
)

type newCollectionCreator struct{}

func (it *newCollectionCreator) Empty() *Collection {
	return &Collection{
		items: []string{},
	}
}

func (it *newCollectionCreator) Cap(capacity int) *Collection {
	collection := make([]string, constants.Zero, capacity)

	return &Collection{
		items: collection,
	}
}

func (it *newCollectionCreator) CloneStrings(stringItems []string) *Collection {
	length := len(stringItems)
	slice := make([]string, 0, length+constants.Capacity4)

	collection := &Collection{
		items: slice,
	}

	return collection.AddStrings(stringItems)
}

func (it *newCollectionCreator) Create(stringItems []string) *Collection {
	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) Strings(stringItems []string) *Collection {
	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) StringsOptions(
	isMakeClone bool,
	stringItems []string,
) *Collection {
	if isMakeClone {
		length := len(stringItems)
		slice := make([]string, 0, length+constants.Capacity4)

		collection := &Collection{
			items: slice,
		}

		return collection.AddStrings(stringItems)
	}

	if len(stringItems) == 0 {
		return it.Empty()
	}

	return &Collection{
		items: stringItems,
	}
}

func (it *newCollectionCreator) LineUsingSep(sep, line string) *Collection {
	lines := strings.Split(line, sep)

	return &Collection{
		items: lines,
	}
}

func (it *newCollectionCreator) LineDefault(compiledLine string) *Collection {
	lines := strings.Split(compiledLine, constants.DefaultLine)

	return &Collection{
		items: lines,
	}
}

func (it *newCollectionCreator) StringsPlusCap(
	additionalCapacity int,
	stringItems []string,
) *Collection {
	if additionalCapacity == 0 {
		return it.Strings(stringItems)
	}

	length := len(stringItems)
	collection := it.Cap(length + additionalCapacity)

	return collection.Adds(stringItems...)
}

func (it *newCollectionCreator) CapStrings(
	additionalCap int,
	stringItems []string,
) *Collection {
	if additionalCap == 0 {
		return it.StringsOptions(
			false,
			stringItems,
		)
	}

	length := len(stringItems)
	collection := it.Cap(length + additionalCap)

	return collection.AddStrings(stringItems)
}

func (it *newCollectionCreator) LenCap(length, capacity int) *Collection {
	collection := make([]string, length, capacity)

	return &Collection{
		items: collection,
	}
}
