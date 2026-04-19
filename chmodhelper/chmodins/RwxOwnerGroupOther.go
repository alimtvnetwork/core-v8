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

package chmodins

import "github.com/alimtvnetwork/core/constants"

// RwxOwnerGroupOther
//
// Owner, Group, Other:
// String Index Values
//   - 0: 'r'/'*'/'-'
//   - 1: 'w'/'*'/'-'
//   - 2: 'x'/'*'/'-'
//
// Examples can be :
//   - "rwx" or
//   - "*wx" or
//   - "rw*" or
//   - "***"
//
// Length must be 3. Not more not less.
type RwxOwnerGroupOther struct {
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Owner string `json:"Owner"`
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Group string `json:"Group"`
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Other string `json:"Other"`
}

// NewRwxOwnerGroupOther
//
// # Each arg ownerRwx, groupRwx, otherRwx should have
//
// Index Values
//   - 0: 'r'/'*'/'-'
//   - 1: 'w'/'*'/'-'
//   - 2: 'x'/'*'/'-'
//
// Examples can be :
//   - "rwx" or
//   - "*wx" or
//   - "rw*" or
//   - "***"
//
// Length must be 3. Not more not less.
func NewRwxOwnerGroupOther(
	ownerRwx,
	groupRwx,
	otherRwx string,
) *RwxOwnerGroupOther {
	return &RwxOwnerGroupOther{
		Owner: ownerRwx,
		Group: groupRwx,
		Other: otherRwx,
	}
}

func (receiver *RwxOwnerGroupOther) IsOwner(rwx string) bool {
	return receiver.Owner == rwx
}

func (receiver *RwxOwnerGroupOther) IsGroup(rwx string) bool {
	return receiver.Group == rwx
}

func (receiver *RwxOwnerGroupOther) IsOther(rwx string) bool {
	return receiver.Other == rwx
}

func (receiver *RwxOwnerGroupOther) ExpandCharOwner() (r, w, x byte) {
	return expandCharsRwx(receiver.Owner)
}

func (receiver *RwxOwnerGroupOther) ExpandCharGroup() (r, w, x byte) {
	return expandCharsRwx(receiver.Group)
}

func (receiver *RwxOwnerGroupOther) ExpandCharOther() (r, w, x byte) {
	return expandCharsRwx(receiver.Other)
}

func (receiver *RwxOwnerGroupOther) Is(
	ownerRwx,
	groupRwx,
	otherRwx string,
) bool {
	return receiver.IsOwner(ownerRwx) &&
		receiver.IsGroup(groupRwx) &&
		receiver.IsOther(otherRwx)
}

func (receiver *RwxOwnerGroupOther) IsEqual(another *RwxOwnerGroupOther) bool {
	if another == nil && receiver == nil {
		return true
	}

	if another == nil || receiver == nil {
		return false
	}

	return receiver.Owner == another.Owner &&
		receiver.Group == another.Group &&
		receiver.Other == another.Other
}

func (receiver *RwxOwnerGroupOther) ToString(isIncludeHyphen bool) string {
	if isIncludeHyphen {
		return receiver.String()
	}

	return receiver.Owner +
		receiver.Group +
		receiver.Other
}

// String : Includes hyphen in-front
// constants.Hyphen +
//
//	receiver.Owner +
//	receiver.Group +
//	receiver.Other
func (receiver *RwxOwnerGroupOther) String() string {
	return constants.Hyphen +
		receiver.Owner +
		receiver.Group +
		receiver.Other
}

func (receiver *RwxOwnerGroupOther) Clone() *RwxOwnerGroupOther {
	if receiver == nil {
		return nil
	}

	return &RwxOwnerGroupOther{
		Owner: receiver.Owner,
		Group: receiver.Group,
		Other: receiver.Other,
	}
}
