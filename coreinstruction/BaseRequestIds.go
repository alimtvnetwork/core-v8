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

import "github.com/alimtvnetwork/core/constants"

type BaseRequestIds struct {
	RequestIds []IdentifierWithIsGlobal `json:"RequestIds,omitempty"`
}

func NewBaseRequestIds(
	isGlobal bool,
	ids ...string,
) *BaseRequestIds {
	return &BaseRequestIds{
		RequestIds: NewRequestIds(
			isGlobal,
			ids...),
	}
}

func NewRequestIds(
	isGlobal bool,
	ids ...string,
) []IdentifierWithIsGlobal {
	slice := make([]IdentifierWithIsGlobal, len(ids))
	if len(ids) == 0 {
		return slice
	}

	for i, id := range ids {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{Id: id},
			IsGlobal:       isGlobal,
		}
	}

	return slice
}

func NewRequestId(
	isGlobal bool,
	id string,
) *IdentifierWithIsGlobal {
	return &IdentifierWithIsGlobal{
		BaseIdentifier: BaseIdentifier{Id: id},
		IsGlobal:       isGlobal,
	}
}

func (b *BaseRequestIds) RequestIdsLength() int {
	if b == nil || b.RequestIds == nil {
		return constants.Zero
	}

	return len(b.RequestIds)
}

func (b *BaseRequestIds) AddReqId(
	requestId IdentifierWithIsGlobal,
) *BaseRequestIds {
	b.RequestIds = append(
		b.RequestIds,
		requestId)

	return b
}

func (b *BaseRequestIds) AddIds(
	isGlobal bool,
	ids ...string,
) *BaseRequestIds {
	if len(ids) == 0 {
		return b
	}

	for _, id := range ids {
		b.RequestIds = append(b.RequestIds, IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{id},
			IsGlobal:       isGlobal,
		})
	}

	return b
}

func (b *BaseRequestIds) IsEmptyRequestIds() bool {
	return b.RequestIdsLength() == 0
}

func (b *BaseRequestIds) HasRequestIds() bool {
	return b != nil && b.RequestIds != nil && len(b.RequestIds) > 0
}

func (b *BaseRequestIds) Clone() *BaseRequestIds {
	if b == nil {
		return nil
	}

	length := b.RequestIdsLength()
	slice := make(
		[]IdentifierWithIsGlobal,
		length)

	if length == 0 {
		return &BaseRequestIds{
			RequestIds: slice,
		}
	}

	for i, reqId := range b.RequestIds {
		slice[i] = IdentifierWithIsGlobal{
			BaseIdentifier: BaseIdentifier{Id: reqId.Id},
			IsGlobal:       reqId.IsGlobal,
		}
	}

	return &BaseRequestIds{
		RequestIds: slice,
	}
}
