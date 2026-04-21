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

import (
	"github.com/alimtvnetwork/core-v8/internal/osconstsinternal"
)

// GOOS values https://stackoverflow.com/a/20728862
//
//goland:noinspection ALL
const (
	LsbCommand                = osconstsinternal.LsbCommand
	Android                   = osconstsinternal.Android
	DarwinOrMacOs             = osconstsinternal.DarwinOrMacOs
	DragonFly                 = osconstsinternal.DragonFly
	FreeBsd                   = osconstsinternal.FreeBsd
	Linux                     = osconstsinternal.Linux
	Nacl                      = osconstsinternal.Nacl
	NetBsd                    = osconstsinternal.NetBsd
	OpenBsd                   = osconstsinternal.OpenBsd
	Plan9                     = osconstsinternal.Plan9
	Solaris                   = osconstsinternal.Solaris
	Windows                   = osconstsinternal.Windows
	Unknown                   = osconstsinternal.Unknown
	Any                       = osconstsinternal.Any
	Illumos                   = osconstsinternal.Illumos
	IOs                       = osconstsinternal.IOs
	Aix                       = osconstsinternal.Aix
	NewLine                   = osconstsinternal.NewLine
	PathSeparator             = osconstsinternal.PathSeparator
	CurrentOperatingSystem    = osconstsinternal.CurrentOperatingSystem
	CurrentSystemArchitecture = osconstsinternal.CurrentSystemArchitecture
	IsWindows                 = osconstsinternal.IsWindows
	IsLinux                   = osconstsinternal.IsLinux
	IsDarwinOrMacOs           = osconstsinternal.IsDarwinOrMacOs
	IsPlan9                   = osconstsinternal.IsPlan9
	IsSolaris                 = osconstsinternal.IsSolaris
	IsFreebsd                 = osconstsinternal.IsFreebsd
	IsNetBsd                  = osconstsinternal.IsNetBsd
	IsOpenBsd                 = osconstsinternal.IsOpenBsd
	IsDragonFly               = osconstsinternal.IsDragonFly
	IsNacl                    = osconstsinternal.IsNacl
	IsUnixGroup               = osconstsinternal.IsUnixGroup
	WindowsCDrive             = osconstsinternal.WindowsCDrive
	LinuxHome                 = osconstsinternal.LinuxHome
	LinuxHomeSlash            = osconstsinternal.LinuxHomeSlash
	LinuxBin                  = osconstsinternal.LinuxBin
	LinuxPermanentTemp        = osconstsinternal.LinuxPermanentTemp // https://prnt.sc/gW0DA5d4jt6R, unix : /var/tmp/
)
