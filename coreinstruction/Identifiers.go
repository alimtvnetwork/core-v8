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

package coreinstruction

import (
	"github.com/alimtvnetwork/core-v8/constants"
)

type Identifiers struct {
	Ids []BaseIdentifier `json:"Ids,omitempty"`
}

func EmptyIdentifiers() *Identifiers {
	return &Identifiers{
		Ids: []BaseIdentifier{},
	}
}

func NewIdentifiersUsingCap(
	capacity int,
) *Identifiers {
	slice := make(
		[]BaseIdentifier,
		0,
		capacity)

	return &Identifiers{Ids: slice}
}

func NewIdentifiers(
	ids ...string,
) *Identifiers {
	slice := make(
		[]BaseIdentifier,
		len(ids))

	if len(ids) == 0 {
		return &Identifiers{
			Ids: []BaseIdentifier{},
		}
	}

	for i, id := range ids {
		slice[i] = BaseIdentifier{
			Id: id,
		}
	}

	return &Identifiers{
		Ids: slice,
	}
}

func (it *Identifiers) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Ids)
}

func (it *Identifiers) IsEmpty() bool {
	return it.Length() == 0
}

func (it *Identifiers) IndexOf(id string) int {
	if id == constants.EmptyString || it.IsEmpty() {
		return constants.InvalidNotFoundCase
	}

	for index, baseIdentifier := range it.Ids {
		if baseIdentifier.Id == id {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}

func (it *Identifiers) GetById(id string) *BaseIdentifier {
	if id == constants.EmptyString || it.IsEmpty() {
		return nil
	}

	for i := range it.Ids {
		if it.Ids[i].Id == id {
			return &it.Ids[i]
		}
	}

	return nil
}

func (it *Identifiers) Add(
	id string,
) *Identifiers {
	if id == constants.EmptyString {
		return it
	}

	it.Ids = append(
		it.Ids,
		BaseIdentifier{Id: id})

	return it
}

func (it *Identifiers) Adds(
	ids ...string,
) *Identifiers {
	if len(ids) == 0 {
		return it
	}

	for _, id := range ids {
		it.Ids = append(
			it.Ids,
			BaseIdentifier{Id: id})
	}

	return it
}

func (it *Identifiers) HasAnyItem() bool {
	return it.Length() > 0
}

func (it *Identifiers) Clone() *Identifiers {
	length := it.Length()

	slice := make(
		[]BaseIdentifier,
		length)

	if length == 0 {
		return &Identifiers{
			Ids: slice,
		}
	}

	for i, baseIdentifier := range it.Ids {
		slice[i] = *baseIdentifier.Clone()
	}

	return &Identifiers{
		Ids: slice,
	}
}
