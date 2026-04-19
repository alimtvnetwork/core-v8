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
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

type FromTo struct {
	From string `json:"From,omitempty"`
	To   string `json:"To,omitempty"`
}

func (it *FromTo) IsNull() bool {
	return it == nil
}

func (it *FromTo) IsFromEmpty() bool {
	return it == nil || it.From == ""
}

func (it *FromTo) IsToEmpty() bool {
	return it == nil || it.To == ""
}

func (it FromTo) String() string {
	return fmt.Sprintf(
		constants.FromToFormat,
		it.From,
		it.To)
}

func (it FromTo) FromName() string {
	return it.From
}

func (it FromTo) ToName() string {
	return it.To
}

func (it *FromTo) SetFromName(from string) {
	if it == nil {
		return
	}

	it.From = from
}

func (it *FromTo) SetToName(to string) {
	if it == nil {
		return
	}

	it.To = to
}

func (it *FromTo) SourceDestination() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.From,
		Destination: it.To,
	}
}

func (it *FromTo) Rename() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.From,
		New:      it.To,
	}
}

func (it FromTo) Clone() FromTo {
	return FromTo{
		From: it.From,
		To:   it.To,
	}
}

func (it *FromTo) ClonePtr() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.From,
		To:   it.To,
	}
}
