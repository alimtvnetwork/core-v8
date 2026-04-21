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

	"github.com/alimtvnetwork/core-v8/constants"
)

type SourceDestination struct {
	Source      string `json:"Source,omitempty"`
	Destination string `json:"Destination,omitempty"`
}

func (it *SourceDestination) IsNull() bool {
	return it == nil
}

func (it *SourceDestination) IsSourceEmpty() bool {
	return it == nil || it.Source == ""
}

func (it *SourceDestination) IsDestinationEmpty() bool {
	return it == nil || it.Destination == ""
}

func (it SourceDestination) String() string {
	return fmt.Sprintf(
		constants.SourceDestinationFormat,
		it.Source,
		it.Destination)
}

func (it SourceDestination) FromName() string {
	return it.Source
}

func (it SourceDestination) ToName() string {
	return it.Destination
}

func (it *SourceDestination) SetFromName(from string) {
	if it == nil {
		return
	}

	it.Source = from
}

func (it *SourceDestination) SetToName(to string) {
	if it == nil {
		return
	}

	it.Destination = to
}

func (it *SourceDestination) FromTo() *FromTo {
	if it == nil {
		return nil
	}

	return &FromTo{
		From: it.Source,
		To:   it.Destination,
	}
}

func (it *SourceDestination) Rename() *Rename {
	if it == nil {
		return nil
	}

	return &Rename{
		Existing: it.Source,
		New:      it.Destination,
	}
}

func (it *SourceDestination) Clone() *SourceDestination {
	if it == nil {
		return nil
	}

	return &SourceDestination{
		Source:      it.Source,
		Destination: it.Destination,
	}
}
