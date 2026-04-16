package coreversiontests

import (
	"testing"

	"github.com/alimtvnetwork/core/coreversion"
	"github.com/alimtvnetwork/core/coretests/args"
)

// Test_Cov5_VersionDisplayMajorMinorPatch_InvalidPatch tests the IsPatchInvalid branch.
func Test_VersionDisplayMajorMinorPatch_InvalidPatch(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 1,
		VersionMinor: 2,
		VersionPatch: -1,
	}

	// Act
	result := v.VersionDisplayMajorMinorPatch()

	// Assert
	expectedStr := "v1.2"
	actualCheck := args.Map{"result": result != expectedStr}
	expectedCheck := args.Map{"result": false}
	expectedCheck.ShouldBeEqual(t, 0, "VersionDisplayMajorMinorPatch with invalid patch", actualCheck)
}

// Test_Cov5_Major_Compare tests the Major() comparison method.
func Test_Major_Compare(t *testing.T) {
	// Arrange
	v := coreversion.Version{
		VersionMajor: 3,
		VersionMinor: 0,
	}

	// Act
	result := v.Major(3)

	// Assert
	actual := args.Map{"result": result.IsEqual()}
	expected := args.Map{"result": true}
	expected.ShouldBeEqual(t, 0, "Major(3) on version with major=3 should be Equal", actual)
}
