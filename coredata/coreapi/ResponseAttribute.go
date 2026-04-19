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
	"github.com/alimtvnetwork/core/reqtype"
)

type ResponseAttribute struct {
	ResponseOfRequestType reqtype.Request
	Count                 int
	HasAnyRecord          bool
	NextPageRequestUrl    string
	StepsPerformed        []string `json:"StepsPerformed,omitempty"`
	DebugInfos            []string `json:"DebugInfos,omitempty"`
	HttpCode              int
	HttpMethod            reqtype.Request
	IsValid               bool
	Message               string `json:"Message,omitempty"`
}

func (receiver *ResponseAttribute) Clone() *ResponseAttribute {
	if receiver == nil {
		return nil
	}

	var steps []string
	if len(receiver.StepsPerformed) > 0 {
		steps = make([]string, len(receiver.StepsPerformed))
		copy(steps, receiver.StepsPerformed)
	}

	var debugs []string
	if len(receiver.DebugInfos) > 0 {
		debugs = make([]string, len(receiver.DebugInfos))
		copy(debugs, receiver.DebugInfos)
	}

	return &ResponseAttribute{
		ResponseOfRequestType: receiver.ResponseOfRequestType,
		Count:                 receiver.Count,
		HasAnyRecord:          receiver.HasAnyRecord,
		NextPageRequestUrl:    receiver.NextPageRequestUrl,
		StepsPerformed:        steps,
		DebugInfos:            debugs,
		HttpCode:              receiver.HttpCode,
		HttpMethod:            receiver.HttpMethod,
		IsValid:               receiver.IsValid,
		Message:               receiver.Message,
	}
}
