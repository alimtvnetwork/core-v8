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

package codestack

import (
	"runtime"
)

type currentNameOf struct{}

func (it currentNameOf) Method() (methodName string) {
	_, _, methodName = it.AllStackSkip(defaultInternalSkip)

	return methodName
}

func (it currentNameOf) MethodByFullName(fullName string) (packageName string) {
	_, _, methodName := it.All(
		fullName,
	)

	return methodName
}

func (it currentNameOf) All(fullFuncName string) (fullMethodName, packageName, methodName string) {
	if fullFuncName == "" {
		return "", "", ""
	}

	return getFuncEverything(fullFuncName)
}

func (it currentNameOf) AllStackSkip(stackSkipIndex int) (fullMethodName, packageName, methodName string) {
	pc, _, _, _ := runtime.Caller(stackSkipIndex + defaultInternalSkip)
	funcInfo := runtime.FuncForPC(pc)
	fullFuncName := funcInfo.Name()

	return it.All(fullFuncName)
}
func (it currentNameOf) MethodStackSkip(stackSkipIndex int) (methodName string) {
	_, _, methodName = it.AllStackSkip(
		stackSkipIndex + defaultInternalSkip,
	)

	return methodName
}

func (it currentNameOf) JoinPackageNameWithRelative(
	fullNameExtractPackageName, relativeNamesJoin string,
) (packageName string) {
	_, packageName, _ = it.All(
		fullNameExtractPackageName,
	)

	return packageName + "." + relativeNamesJoin
}

func (it currentNameOf) Package() (packageName string) {
	_, packageName, _ = it.AllStackSkip(
		defaultInternalSkip,
	)

	return packageName
}

func (it currentNameOf) PackageByFullName(fullName string) (packageName string) {
	_, packageName, _ = it.All(
		fullName,
	)

	return packageName
}

func (it currentNameOf) PackageStackSkip(stackSkipIndex int) (packageName string) {
	_, packageName, _ = it.AllStackSkip(
		stackSkipIndex + defaultInternalSkip,
	)

	return packageName
}

func (it currentNameOf) CurrentFuncFullPath(fullName string) (packageName string) {
	fullMethodNameOf, _, _ := NameOf.All(
		fullName,
	)

	return fullMethodNameOf
}
