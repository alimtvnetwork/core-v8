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

package messages

const (
	HyphenedRwxRwxRwxLengthMustBe10 = "`-rwxrwxrwx` length must be 10. Reference : " +
		"https://ss64.com/bash/chmod.html"
	RwxRwxRwxLengthMustBe9                        = "`rwxrwxrwx` length must be 9."
	ModeCharShouldBeAllNumbersAndWithin0To7       = "mode char should be all digits and under 0 to 7"
	DynamicFailedToParseToFloat64BecauseNull      = "dynamic datatype failed to parse to float64 because it is nil."
	FailedToCompileChmodhelperVarWrapperToWrapper = "Failed to compile chmodhelper.VarWrapper" +
		" to failedToExecute.Wrapper."
	FailedToGetFileModeRwx                             = "Failed to get the existing filemode (Rwx...)."
	PathNotExist                                       = "Path doesn't exist"
	CannotVerifyEmptyContentsWhereValidatorsArePresent = "cannot verify empty text contents where validators are present"
)
