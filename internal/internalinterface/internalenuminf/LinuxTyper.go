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

package internalenuminf

import "fmt"

type LinuxTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsUnknown() bool
	IsUbuntuServer() bool
	IsUbuntuServer18() bool
	IsUbuntuServer19() bool
	IsUbuntuServer20() bool
	IsUbuntuServer21() bool
	IsUbuntuDesktop() bool
	IsCentos() bool
	IsCentos7() bool
	IsCentos8() bool
	IsCentos9() bool
	IsDebianServer() bool
	IsDebianDesktop() bool
	IsDocker() bool
	IsDockerUbuntuServer() bool
	IsDockerUbuntuServer20() bool
	IsDockerUbuntuServer21() bool
	IsDockerCentos9() bool
	IsAndroid() bool
	fmt.Stringer
}
