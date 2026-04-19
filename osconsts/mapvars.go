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

package osconsts

var (
	X32ArchitecturesMap = map[string]bool{
		"386":         true,
		"arm":         true,
		"armbe":       true,
		"mips":        true,
		"amd64p32":    true,
		"mips64p32":   true,
		"mips64p32le": true,
		"ppc":         true,
		"riscv":       true,
		"s390":        true,
		"sparc":       true,
	}

	X64ArchitecturesMap = map[string]bool{
		"amd64":    true,
		"arm64":    true,
		"ppc64":    true,
		"ppc64le":  true,
		"mips64":   true,
		"mips64le": true,
		"riscv64":  true,
		"s390x":    true,
		"wasm":     true,
		"arm64be":  true,
		"sparc64":  true,
	}

	UnixGroupsMap = map[string]bool{
		"android":   true,
		"darwin":    true,
		"dragonfly": true,
		"freebsd":   true,
		"linux":     true,
		"nacl":      true,
		"netbsd":    true,
		"openBSD":   true,
		"plan9":     true,
		"solaris":   true,
	}
)
