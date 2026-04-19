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

import (
	"errors"
	"fmt"
)

// CompiledError wraps a main error with an additional message, preserving the error chain.
func CompiledError(mainErr error, additionalMessage string) error {
	if mainErr == nil {
		return nil
	}

	if additionalMessage == "" {
		return mainErr
	}

	return fmt.Errorf("%s: %w",
		additionalMessage, mainErr)
}

// CompiledErrorString returns the compiled error string with main error and additional message.
// Kept for backward compatibility with string-based consumers.
func CompiledErrorString(mainErr error, additionalMessage string) string {
	if mainErr == nil {
		return ""
	}

	compiled := CompiledError(mainErr, additionalMessage)
	if compiled == nil {
		return ""
	}

	return compiled.Error()
}

// JoinErrors is a convenience alias for errors.Join, provided for discoverability
// within the errcore package.
func JoinErrors(errs ...error) error {
	return errors.Join(errs...)
}
