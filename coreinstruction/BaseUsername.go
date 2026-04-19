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
	"regexp"
	"strings"

	"github.com/alimtvnetwork/core/internal/strutilinternal"
)

type BaseUsername struct {
	Username string `json:"Username"`
}

func NewUsername(user string) *BaseUsername {
	return &BaseUsername{Username: user}
}

func (it *BaseUsername) UsernameString() string {
	return it.Username
}

func (it *BaseUsername) IsUsernameEmpty() bool {
	return it == nil || it.Username == ""
}

func (it *BaseUsername) IsUsernameWhitespace() bool {
	return it == nil || strutilinternal.IsNullOrEmptyOrWhitespace(&it.Username)
}

func (it *BaseUsername) IsUsername(user string) bool {
	return it.Username == user
}

func (it *BaseUsername) IsUsernameCaseInsensitive(usernameInsensitive string) bool {
	return strings.EqualFold(it.Username, usernameInsensitive)
}

func (it *BaseUsername) IsUsernameContains(usernameContains string) bool {
	return strings.Contains(it.Username, usernameContains)
}

func (it *BaseUsername) IsUsernameRegexMatches(regex *regexp.Regexp) bool {
	return regex.MatchString(it.Username)
}

func (it *BaseUsername) IsEqual(right *BaseUsername) bool {
	if it == nil && right == nil {
		return true
	}

	if it == nil || right == nil {
		return false
	}

	return it.Username == right.Username
}

func (it *BaseUsername) IsNotEqual(right *BaseUsername) bool {
	return !it.IsEqual(right)
}

func (it *BaseUsername) ClonePtr() *BaseUsername {
	if it == nil {
		return nil
	}

	return &BaseUsername{
		Username: it.Username,
	}
}

func (it BaseUsername) Clone() BaseUsername {
	return BaseUsername{
		Username: it.Username,
	}
}
