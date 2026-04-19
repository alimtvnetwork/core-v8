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

package osconstsinternal

import (
	"runtime"

	"github.com/alimtvnetwork/core/constants"
)

// GOOS values https://stackoverflow.com/a/20728862
//
//goland:noinspection ALL
const (
	LsbCommand                = "lsb_release"
	Android                   = "android"
	DarwinOrMacOs             = "darwin"
	DragonFly                 = "dragonfly"
	FreeBsd                   = "freebsd"
	Linux                     = "linux"
	Nacl                      = "nacl"
	NetBsd                    = "netbsd"
	OpenBsd                   = "openbsd"
	Plan9                     = "plan9"
	Solaris                   = "solaris"
	Windows                   = "windows"
	Unknown                   = "Unknown"
	Any                       = "Any"
	Illumos                   = "illumos"
	IOs                       = "ios"
	Aix                       = "aix"
	NewLine                   = constants.NewLine
	PathSeparator             = constants.PathSeparator
	CurrentOperatingSystem    = runtime.GOOS
	CurrentSystemArchitecture = runtime.GOARCH
	IsWindows                 = CurrentOperatingSystem == Windows
	IsLinux                   = CurrentOperatingSystem == Linux
	IsDarwinOrMacOs           = CurrentOperatingSystem == DarwinOrMacOs
	IsPlan9                   = CurrentOperatingSystem == Plan9
	IsSolaris                 = CurrentOperatingSystem == Solaris
	IsFreebsd                 = CurrentOperatingSystem == FreeBsd
	IsNetBsd                  = CurrentOperatingSystem == NetBsd
	IsOpenBsd                 = CurrentOperatingSystem == OpenBsd
	IsDragonFly               = CurrentOperatingSystem == DragonFly
	IsNacl                    = CurrentOperatingSystem == Nacl
	IsUnixGroup               = !IsWindows
	WindowsCDrive             = "C:\\"
	LinuxHome                 = "/home"
	LinuxHomeSlash            = "/home/"
	LinuxBin                  = "/bin"
	LinuxPermanentTemp        = "/var/tmp/"         // https://prnt.sc/gW0DA5d4jt6R, unix : /var/tmp/
	WindowsPermanentTemp      = "c:\\Windows\\Temp" // or %temp% expand
)
