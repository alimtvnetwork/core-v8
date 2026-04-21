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

package corepayload

import (
	"github.com/alimtvnetwork/core-v8/constants"
	"github.com/alimtvnetwork/core-v8/converters"
)

type SessionInfo struct {
	Id          string `json:"Id,omitempty"`
	User        *User  `json:"User,omitempty"`
	SessionPath string `json:"SessionPath,omitempty"`
}

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it SessionInfo) IdentifierInteger() int {
	if it.Id == "" {
		return constants.InvalidValue
	}

	idInt, _ := converters.StringTo.IntegerWithDefault(
		it.Id,
		constants.InvalidValue,
	)

	return idInt
}

// IdentifierUnsignedInteger
//
// Invalid value returns constants.Zero
func (it SessionInfo) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
}

func (it *SessionInfo) IsEmpty() bool {
	return it == nil || (it.Id == "" && it.User == nil && it.SessionPath == "")
}

func (it *SessionInfo) IsValid() bool {
	return !it.IsEmpty() && it.Id != ""
}

func (it *SessionInfo) IsUserNameEmpty() bool {
	return it == nil || it.User.IsNameEmpty()
}

func (it *SessionInfo) IsUserEmpty() bool {
	return it == nil || it.User.IsEmpty()
}

func (it *SessionInfo) HasUser() bool {
	return it != nil && it.User.IsValidUser()
}

func (it *SessionInfo) IsUsernameEqual(
	name string,
) bool {
	return it != nil &&
		it.User.IsNameEqual(name)
}

func (it SessionInfo) Clone() SessionInfo {
	return SessionInfo{
		Id:          it.Id,
		User:        it.User.ClonePtr(),
		SessionPath: it.SessionPath,
	}
}

func (it *SessionInfo) ClonePtr() *SessionInfo {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it SessionInfo) Ptr() *SessionInfo {
	return &it
}
