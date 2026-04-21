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

package coreapi

import (
	"github.com/alimtvnetwork/core-v8/reqtype"
)

type RequestAttribute struct {
	Url           string `json:"Url,omitempty"`
	Host          string `json:"Host,omitempty"`
	ResourceName  string `json:"ResourceName,omitempty"`
	ActionName    string `json:"ActionName,omitempty"`
	Identifier    string `json:"Identifier,omitempty"`
	OptionalAuth  string `json:"OptionalAuth,omitempty"`
	ErrorJson     string `json:"ErrorJson,omitempty"`
	RequestType   reqtype.Request
	IsValid       bool
	HasError      bool
	SearchRequest *SearchRequest `json:"SearchRequest,omitempty"`
	PageRequest   *PageRequest   `json:"PageRequest,omitempty"`
}

func (it *RequestAttribute) HasSearchRequest() bool {
	return it != nil && it.SearchRequest != nil
}

func (it *RequestAttribute) HasPageRequest() bool {
	return it != nil && it.PageRequest != nil
}

func (it *RequestAttribute) IsEmpty() bool {
	return it == nil
}

func (it *RequestAttribute) IsAnyNull() bool {
	return it == nil
}

func (it *RequestAttribute) IsPageRequestEmpty() bool {
	return it == nil || it.PageRequest == nil
}

func (it *RequestAttribute) IsSearchRequestEmpty() bool {
	return it == nil || it.SearchRequest == nil
}

func (it *RequestAttribute) Clone() *RequestAttribute {
	if it == nil {
		return nil
	}

	return &RequestAttribute{
		Url:           it.Url,
		Host:          it.Host,
		ResourceName:  it.ResourceName,
		RequestType:   it.RequestType,
		IsValid:       it.IsValid,
		SearchRequest: it.SearchRequest.Clone(),
		PageRequest:   it.PageRequest.Clone(),
	}
}
