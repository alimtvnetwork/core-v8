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

type UserInfo struct {
	User       *User `json:"User,omitempty"`       // refers to control panel or any kinds of virtual user
	SystemUser *User `json:"SystemUser,omitempty"` // System or OS user
}

func (it *UserInfo) SetUserSystemUser(
	user, systemUser *User,
) *UserInfo {
	if it == nil {
		return &UserInfo{
			User:       user,
			SystemUser: systemUser,
		}
	}

	it.User = user
	it.SystemUser = systemUser

	return it
}

// SetUser
//
// on null creates new
func (it *UserInfo) SetUser(
	user *User,
) *UserInfo {
	if it == nil {
		return &UserInfo{
			User: user,
		}
	}

	it.User = user

	return it
}

// SetSystemUser
//
// on null creates new
func (it *UserInfo) SetSystemUser(
	systemUser *User,
) *UserInfo {
	if it == nil {
		return &UserInfo{
			SystemUser: systemUser,
		}
	}

	it.SystemUser = systemUser

	return it
}

func (it *UserInfo) HasUser() bool {
	return it != nil && it.User.IsValidUser()
}

func (it *UserInfo) HasSystemUser() bool {
	return it != nil && it.SystemUser.IsValidUser()
}

func (it *UserInfo) IsEmpty() bool {
	return it == nil || it.User.IsEmpty() && it.SystemUser.IsEmpty()
}

func (it *UserInfo) IsUserEmpty() bool {
	return it == nil || it.User.IsEmpty()
}

func (it *UserInfo) IsSystemUserEmpty() bool {
	return it == nil || it.SystemUser.IsEmpty()
}

func (it *UserInfo) Clone() UserInfo {
	return UserInfo{
		User:       it.User.ClonePtr(),
		SystemUser: it.SystemUser.ClonePtr(),
	}
}

func (it *UserInfo) ClonePtr() *UserInfo {
	if it == nil {
		return nil
	}

	return it.Clone().Ptr()
}

func (it UserInfo) Ptr() *UserInfo {
	return &it
}

func (it *UserInfo) ToNonPtr() UserInfo {
	if it == nil {
		return UserInfo{}
	}

	return *it
}
