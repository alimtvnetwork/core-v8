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

package chmodhelper

import (
	"os"

	"github.com/alimtvnetwork/core/constants"
)

// Reference : https://ss64.com/bash/chmod.html
const (
	SingleRwxLengthString                           = "3"
	HyphenedRwxLength                               = constants.ArbitraryCapacity10
	HyphenedRwxLengthString                         = constants.N10String
	FullRwxLengthWithoutHyphenString                = constants.N9String
	FullRwxLengthWithoutHyphen                      = constants.ArbitraryCapacity9
	SingleRwxLength                                 = 3
	ReadValue                                       = 4
	WriteValue                                      = 2
	ExecuteValue                                    = 1
	ReadWriteValue                                  = ReadValue + WriteValue
	ReadExecuteValue                                = ReadValue + ExecuteValue
	WriteExecuteValue                               = WriteValue + ExecuteValue
	ReadWriteExecuteValue                           = ReadValue + WriteValue + ExecuteValue
	OwnerIndex                                      = 0
	GroupIndex                                      = 1
	OtherIndex                                      = 2
	ReadChar                            byte        = 'r'
	NopChar                             byte        = '-'
	WriteChar                           byte        = 'w'
	ExecuteChar                         byte        = 'x'
	AllWildcards                                    = "***"
	pathInvalidMessage                              = "Path invalid or permission issue. path : "
	chmodExpectationFailed                          = "chmod expectation failed"
	messageWithPathWrappedFormat                    = "%s Path: (\"%s\")"     // message, path
	fileDefaultChmod                    os.FileMode = 0644                    // cannot execute by everyone OwnerCanReadWriteGroupOtherCanReadOnly
	dirDefaultChmod                     os.FileMode = 0755                    // can execute by everyone OwnerCanDoAllExecuteGroupOtherCanReadExecute
	fileModeStringFriendlyDisplayFormat             = "{chmod : \"%s (%s)\"}" // chmod - 0777, -rwxrwxrwx
)
