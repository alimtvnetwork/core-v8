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
	"github.com/alimtvnetwork/core/constants"
)

type IdentifiersWithGlobals struct {
	IdentifierWithIsGlobals []IdentifierWithIsGlobal `json:"IdentifierWithIsGlobals"`
}

func EmptyIdentifiersWithGlobals() *IdentifiersWithGlobals {
	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: []IdentifierWithIsGlobal{},
	}
}

func NewIdentifiersWithGlobals(
	isGlobal bool,
	ids ...string,
) *IdentifiersWithGlobals {
	slice := make(
		[]IdentifierWithIsGlobal,
		len(ids))

	if len(ids) == 0 {
		return &IdentifiersWithGlobals{
			IdentifierWithIsGlobals: slice,
		}
	}

	for i, id := range ids {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{
				Id: id,
			},
			IsGlobal: isGlobal,
		}
	}

	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: slice,
	}
}

func (receiver *IdentifiersWithGlobals) Length() int {
	if receiver == nil {
		return 0
	}

	return len(receiver.IdentifierWithIsGlobals)
}

func (receiver *IdentifiersWithGlobals) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *IdentifiersWithGlobals) IndexOf(id string) int {
	if id == constants.EmptyString || receiver.IsEmpty() {
		return constants.InvalidNotFoundCase
	}

	for index, identifierWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		if identifierWithIsGlobal.Id == id {
			return index
		}
	}

	return constants.InvalidNotFoundCase
}

func (receiver *IdentifiersWithGlobals) GetById(id string) *IdentifierWithIsGlobal {
	if id == constants.EmptyString || receiver.IsEmpty() {
		return nil
	}

	for i := range receiver.IdentifierWithIsGlobals {
		if receiver.IdentifierWithIsGlobals[i].Id == id {
			return &receiver.IdentifierWithIsGlobals[i]
		}
	}

	return nil
}

func (receiver *IdentifiersWithGlobals) Add(
	isGlobal bool,
	id string,
) *IdentifiersWithGlobals {
	if id == constants.EmptyString {
		return receiver
	}

	receiver.IdentifierWithIsGlobals = append(
		receiver.IdentifierWithIsGlobals,
		*NewIdentifierWithIsGlobal(id, isGlobal))

	return receiver
}

func (receiver *IdentifiersWithGlobals) Adds(
	isGlobal bool,
	ids ...string,
) *IdentifiersWithGlobals {
	if len(ids) == 0 {
		return receiver
	}

	for _, id := range ids {
		receiver.IdentifierWithIsGlobals = append(
			receiver.IdentifierWithIsGlobals,
			*NewIdentifierWithIsGlobal(id, isGlobal))
	}

	return receiver
}

func (receiver *IdentifiersWithGlobals) HasAnyItem() bool {
	return receiver.Length() > 0
}

func (receiver *IdentifiersWithGlobals) Clone() *IdentifiersWithGlobals {
	length := receiver.Length()

	slice := make(
		[]IdentifierWithIsGlobal,
		length)

	if length == 0 {
		return &IdentifiersWithGlobals{
			IdentifierWithIsGlobals: slice,
		}
	}

	for i, idWithIsGlobal := range receiver.IdentifierWithIsGlobals {
		slice[i] = *idWithIsGlobal.Clone()
	}

	return &IdentifiersWithGlobals{
		IdentifierWithIsGlobals: slice,
	}
}
