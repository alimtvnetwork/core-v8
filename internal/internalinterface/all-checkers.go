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

package internalinterface

type IsIdentifierEqualer interface {
	IsIdentifier(identifier string) bool
}

type IsIdEqualer interface {
	IsId(id string) bool
}

type IsIdUnsignedIntegerEqualer interface {
	IsId(id uint) bool
}

type HasErrorChecker interface {
	HasError() bool
}

type IsValidChecker interface {
	// IsValid similar or alias for IsSuccessChecker
	IsValid() bool
}

type IsInvalidChecker interface {
	IsInvalid() bool
}
type IsSuccessChecker interface {
	// IsSuccess No error
	IsSuccess() bool
}

type IsFailedChecker interface {
	// IsFailed has error or any other issues, or alias for HasIssues or HasError
	IsFailed() bool
}

type IsSuccessValidator interface {
	IsValidChecker
	IsSuccessChecker
	IsFailedChecker
}

type IsEmptyChecker interface {
	IsEmpty() bool
}

type IsEmptyErrorChecker interface {
	IsEmptyError() bool
}

type IsErrorEqualsChecker interface {
	IsErrorEquals(err error) bool
	IsErrorMessageEqual(msg string) bool
	IsErrorMessage(msg string, isCaseSensitive bool) bool
	IsErrorMessageContains(
		msg string,
		isCaseSensitive bool,
	) bool
}

type HasErrorOrHasAnyErrorChecker interface {
	HasError() bool
	HasAnyError() bool
}
