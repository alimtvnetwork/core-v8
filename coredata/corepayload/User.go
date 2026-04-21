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
	"github.com/alimtvnetwork/core-v8/coredata/corejson"
)

type User struct {
	Identifier   string `json:"Identifier,omitempty"`
	Name         string `json:"Name,omitempty"`
	Type         string `json:"Type,omitempty"`
	AuthToken    string `json:"AuthToken,omitempty"`
	PasswordHash string `json:"PasswordHash,omitempty"`
	IsSystemUser bool   `json:"IsSystemUser"`
}

// IdentifierInteger
//
// Invalid value returns constants.InvalidValue
func (it User) IdentifierInteger() int {
	if it.Identifier == "" {
		return constants.InvalidValue
	}

	idInt, _ := converters.StringTo.IntegerWithDefault(
		it.Identifier,
		constants.InvalidValue,
	)

	return idInt
}

// IdentifierUnsignedInteger
//
// Invalid value returns constants.Zero
func (it User) IdentifierUnsignedInteger() uint {
	idInt := it.IdentifierInteger()

	if idInt < 0 {
		return constants.Zero
	}

	return uint(idInt)
}

func (it *User) HasAuthToken() bool {
	return it != nil && it.AuthToken != ""
}

func (it *User) HasPasswordHash() bool {
	return it != nil && it.PasswordHash != ""
}

func (it *User) IsPasswordHashEmpty() bool {
	return it == nil || it.PasswordHash == ""
}

func (it *User) IsAuthTokenEmpty() bool {
	return it == nil || it.AuthToken == ""
}

func (it *User) IsEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *User) IsValidUser() bool {
	return !it.IsEmpty()
}

func (it *User) IsNameEmpty() bool {
	return it == nil || it.Name == ""
}

func (it *User) IsNameEqual(name string) bool {
	return it != nil && it.Name == name
}

func (it *User) IsNotSystemUser() bool {
	return it != nil && !it.IsSystemUser
}

func (it *User) IsVirtualUser() bool {
	return it != nil && !it.IsSystemUser
}

func (it *User) HasType() bool {
	return it != nil && it.Type != ""
}

func (it *User) IsTypeEmpty() bool {
	return it == nil || it.Type == ""
}

func (it User) String() string {
	return it.JsonPtr().JsonString()
}

func (it *User) PrettyJsonString() string {
	return it.JsonPtr().PrettyJsonString()
}

func (it *User) Json() corejson.Result {
	return corejson.New(it)
}

func (it *User) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

func (it *User) Serialize() ([]byte, error) {
	return corejson.Serialize.Raw(it)
}

func (it *User) Deserialize(rawJsonBytes []byte) error {
	return corejson.
		Deserialize.
		UsingBytes(rawJsonBytes, it)
}

func (it User) Clone() User {
	return User{
		Name:         it.Name,
		Type:         it.Type,
		AuthToken:    it.AuthToken,
		PasswordHash: it.PasswordHash,
		IsSystemUser: it.IsSystemUser,
	}
}

func (it *User) ClonePtr() *User {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it User) Ptr() *User {
	return &it
}
