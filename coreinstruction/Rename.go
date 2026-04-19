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

type Rename struct {
	Existing string `json:"Existing,omitempty"`
	New      string `json:"New,omitempty"`
}

func (it *Rename) IsNull() bool {
	return it == nil
}

func (it *Rename) IsExistingEmpty() bool {
	return it == nil || it.Existing == ""
}

func (it *Rename) IsNewEmpty() bool {
	return it == nil || it.New == ""
}

func (it Rename) String() string {
	return fmt.Sprintf(
		constants.RenameFormat,
		it.Existing,
		it.New)
}

func (it *Rename) SourceDestination() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.Existing,
		Destination: it.New,
	}
}

func (it Rename) FromName() string {
	return it.Existing
}

func (it Rename) ExistingName() string {
	return it.Existing
}

func (it Rename) NewName() string {
	return it.New
}

func (it Rename) ToName() string {
	return it.New
}

func (it *Rename) SetFromName(from string) {
	if it == nil {
		return
	}

	it.Existing = from
}

func (it *Rename) SetToName(to string) {
	if it == nil {
		return
	}

	it.New = to
}

func (it *Rename) FromTo() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.Existing,
		To:   it.New,
	}
}

func (it *Rename) Clone() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.Existing,
		New:      it.New,
	}
}
