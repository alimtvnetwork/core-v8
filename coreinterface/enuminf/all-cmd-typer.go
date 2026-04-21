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

package enuminf

import "github.com/alimtvnetwork/core-v8/internal/internalinterface/internalenuminf"

type BaseCmdTyper interface {
	BasicEnumer
	internalenuminf.BaseCmdTyper
}

type ToNameLower interface {
	internalenuminf.ToNameLower
}

type ListCmdTyper interface {
	BasicEnumer
	internalenuminf.ListCmdTyper
}

type CreateCmdTyper interface {
	BasicEnumer
	internalenuminf.CreateCmdTyper
}

type AddCmdTyper interface {
	BasicEnumer
	internalenuminf.AddCmdTyper
}

type RemoveCmdTyper interface {
	BasicEnumer
	internalenuminf.RemoveCmdTyper
}

type ImportCmdTyper interface {
	BasicEnumer
	internalenuminf.ImportCmdTyper
}

type ExportCmdTyper interface {
	BasicEnumer
	internalenuminf.ExportCmdTyper
}

type BackupCmdTyper interface {
	BasicEnumer
	internalenuminf.BackupCmdTyper
}

type ImportExportBackupCmdTyper interface {
	BasicEnumer
	internalenuminf.ImportExportBackupCmdTyper
}

type HelpCmdTyper interface {
	BasicEnumer
	internalenuminf.HelpCmdTyper
}

type LogCmdTyper interface {
	BasicEnumer
	internalenuminf.LogCmdTyper
}

type StatusCmdTyper interface {
	BasicEnumer
	internalenuminf.StatusCmdTyper
}

type InstallCmdTyper interface {
	BasicEnumer
	internalenuminf.InstallCmdTyper
}

type ReinstallCmdTyper interface {
	BasicEnumer
	internalenuminf.ReinstallCmdTyper
}

type CleanupCmdTyper interface {
	BasicEnumer
	internalenuminf.CleanupCmdTyper
}

type UninstallCmdTyper interface {
	BasicEnumer
	internalenuminf.CleanupCmdTyper
}

type ApplyInstallFixCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyInstallFixCmdTyper
}

type RevertCmdTyper interface {
	BasicEnumer
	internalenuminf.RevertCmdTyper
}

type HistoriesCmdTyper interface {
	BasicEnumer
	internalenuminf.HistoriesCmdTyper
}

type CompressCmdTyper interface {
	BasicEnumer
	internalenuminf.CompressCmdTyper
}

type DecompressCmdTyper interface {
	BasicEnumer
	internalenuminf.DecompressCmdTyper
}

type DownloadCmdTyper interface {
	BasicEnumer
	internalenuminf.DownloadCmdTyper
}

type DownloadDecompressCmdTyper interface {
	BasicEnumer
	internalenuminf.DownloadDecompressCmdTyper
}

type ChangePortCmdTyper interface {
	BasicEnumer
	internalenuminf.ChangePortCmdTyper
}

type SwitchPortCmdTyper interface {
	BasicEnumer
	internalenuminf.SwitchPortCmdTyper
}

type SwitchOrChangePortCmdTyper interface {
	ChangePortCmdTyper
	SwitchPortCmdTyper
}

type WhichPortCmdTyper interface {
	BasicEnumer
	internalenuminf.WhichPortCmdTyper
}

type ChangeDirCmdTyper interface {
	BasicEnumer
	internalenuminf.ChangeDirCmdTyper
}

type ApplyDefaultChmodCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyDefaultChmodCmdTyper
}

type ApplyDefaultConfigureCmdTyper interface {
	BasicEnumer
	internalenuminf.ApplyDefaultConfigureCmdTyper
}

type SyncCmdTyper interface {
	BasicEnumer
	internalenuminf.SyncCmdTyper
}
