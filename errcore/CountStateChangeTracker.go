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

package errcore

type CountStateChangeTracker struct {
	lengthGetter
	initLength int
}

func NewCountStateChangeTracker(
	lengthGetter lengthGetter,
) CountStateChangeTracker {
	return CountStateChangeTracker{
		lengthGetter: lengthGetter,
		initLength:   lengthGetter.Length(),
	}
}

func (it CountStateChangeTracker) IsSameStateUsingCount(
	currentCount int,
) bool {
	return currentCount == it.initLength
}

// IsSameState returns true if the current length matches the initial length.
// IsValid and IsSuccess are aliases with identical behavior.
func (it CountStateChangeTracker) IsSameState() bool {
	return it.lengthGetter.Length() == it.initLength
}

// IsValid is an alias for IsSameState.
func (it CountStateChangeTracker) IsValid() bool {
	return it.IsSameState()
}

// IsSuccess is an alias for IsSameState.
func (it CountStateChangeTracker) IsSuccess() bool {
	return it.IsSameState()
}

// HasChanges returns true if the current length differs from the initial length.
// IsFailed is an alias with identical behavior.
func (it CountStateChangeTracker) HasChanges() bool {
	return it.lengthGetter.Length() != it.initLength
}

// IsFailed is an alias for HasChanges.
func (it CountStateChangeTracker) IsFailed() bool {
	return it.HasChanges()
}
