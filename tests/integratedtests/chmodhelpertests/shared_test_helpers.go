package chmodhelpertests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
)

// covTempDir creates a temporary directory for test use.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covTempDir(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

// covWriteFile writes a file in a directory and returns the full path.
// Moved here from Coverage_test.go so split-recovery subfolders can see it.
func covWriteFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	p := filepath.Join(dir, name)
	err := os.WriteFile(p, []byte(content), 0644)
	convey.Convey("covWriteFile creates "+name, t, func() {
		convey.So(err, should.BeNil)
	})
	return p
}

// newTestRW creates a SimpleFileReaderWriter for testing.
// Moved here from Coverage12_SimpleFileRW_test.go so split-recovery subfolders can see it.
func newTestRW(dir, file string) chmodhelper.SimpleFileReaderWriter {
	return chmodhelper.SimpleFileReaderWriter{
		ChmodDir:               0755,
		ChmodFile:              0644,
		ParentDir:              dir,
		FilePath:               filepath.Join(dir, file),
		IsMustChmodApplyOnFile: true,
		IsApplyChmodOnMismatch: true,
	}
}
