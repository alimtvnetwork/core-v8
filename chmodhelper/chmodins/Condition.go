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

type Condition struct {
	IsSkipOnInvalid   bool `json:"IsSkipOnInvalid"`
	IsContinueOnError bool `json:"IsContinueOnError"`
	IsRecursive       bool `json:"IsRecursive"`
}

func DefaultAllTrueCondition() *Condition {
	return &Condition{
		IsSkipOnInvalid:   true,
		IsContinueOnError: true,
		IsRecursive:       true,
	}
}

func DefaultAllFalseCondition() *Condition {
	return &Condition{
		IsSkipOnInvalid:   false,
		IsContinueOnError: false,
		IsRecursive:       false,
	}
}

// DefaultAllFalseExceptRecurse only IsRecursive will be true
func DefaultAllFalseExceptRecurse() *Condition {
	return &Condition{
		IsSkipOnInvalid:   false,
		IsContinueOnError: false,
		IsRecursive:       true,
	}
}

// IsExitOnInvalid
//
//	returns true Condition is null or invert of IsSkipOnInvalid
func (it *Condition) IsExitOnInvalid() bool {
	return it == nil || !it.IsSkipOnInvalid
}

// IsCollectErrorOnInvalid
//
//	returns true Condition is null or invert of IsSkipOnInvalid
func (it *Condition) IsCollectErrorOnInvalid() bool {
	return it == nil || !it.IsSkipOnInvalid
}

func (it *Condition) Clone() *Condition {
	if it == nil {
		return nil
	}

	return &Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}
}

func (it Condition) CloneNonPtr() Condition {
	return Condition{
		IsSkipOnInvalid:   it.IsSkipOnInvalid,
		IsContinueOnError: it.IsContinueOnError,
		IsRecursive:       it.IsRecursive,
	}
}
