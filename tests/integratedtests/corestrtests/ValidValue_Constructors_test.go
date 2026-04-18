package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

// ══════════════════════════════════════════════════════════════════════════════
// ValidValue — Constructors (NewValidValue, NewValidValueEmpty, InvalidValidValue, UsingAny variants)
// ══════════════════════════════════════════════════════════════════════════════

func Test_ValidValue_NewValidValue_StoresValueAndIsValid(t *testing.T) {
	safeTest(t, "Test_NewValidValue", func() {
		// Arrange
		v := corestr.NewValidValue("hello")

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "hello",
			"valid": true,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "NewValidValue returns correct value -- basic", actual)
	})
}


func Test_ValidValue_NewValidValueEmpty_ReturnsEmptyValid(t *testing.T) {
	safeTest(t, "Test_ValidValue_NewValidValueEmpty_ReturnsEmptyValid", func() {
		// Arrange
		v := corestr.NewValidValueEmpty()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueEmpty returns correct value -- empty", actual)
	})
}


func Test_ValidValue_InvalidValidValue_StoresMessage(t *testing.T) {
	safeTest(t, "Test_ValidValue_InvalidValidValue_StoresMessage", func() {
		// Arrange
		v := corestr.InvalidValidValue("bad input")

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "bad input",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValue returns correct value -- with message", actual)
	})
}


func Test_ValidValue_InvalidValidValueNoMessage_HasEmptyMessage(t *testing.T) {
	safeTest(t, "Test_ValidValue_InvalidValidValueNoMessage_HasEmptyMessage", func() {
		// Arrange
		v := corestr.InvalidValidValueNoMessage()

		// Act
		actual := args.Map{
			"val": v.Value,
			"valid": v.IsValid,
			"msg": v.Message,
		}

		// Assert
		expected := args.Map{
			"val": "",
			"valid": false,
			"msg": "",
		}
		expected.ShouldBeEqual(t, 0, "InvalidValidValueNoMessage returns correct value -- no message", actual)
	})
}


func Test_ValidValue_NewValidValueUsingAny_FromIntInput(t *testing.T) {
	safeTest(t, "Test_ValidValue_NewValidValueUsingAny_FromIntInput", func() {
		// Arrange
		v := corestr.NewValidValueUsingAny(false, true, 42)

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"notEmpty": v.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAny returns correct value -- int input", actual)
	})
}


func Test_ValidValue_NewValidValueUsingAnyAutoValid_NonEmptyKeepsGivenFlag(t *testing.T) {
	safeTest(t, "Test_ValidValue_NewValidValueUsingAnyAutoValid_NonEmptyKeepsGivenFlag", func() {
		// Arrange
		v := corestr.NewValidValueUsingAnyAutoValid(false, "hello")

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"notEmpty": v.Value != "",
		}

		// Assert
		expected := args.Map{
			"valid": false,
			"notEmpty": true,
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- non-empty", actual)
	})
}


func Test_ValidValue_NewValidValueUsingAnyAutoValid_EmptyForcesValid(t *testing.T) {
	safeTest(t, "Test_ValidValue_NewValidValueUsingAnyAutoValid_EmptyForcesValid", func() {
		// Arrange
		v := corestr.NewValidValueUsingAnyAutoValid(false, "")

		// Act
		actual := args.Map{
			"valid": v.IsValid,
			"val": v.Value,
		}

		// Assert
		expected := args.Map{
			"valid": true,
			"val": "",
		}
		expected.ShouldBeEqual(t, 0, "NewValidValueUsingAnyAutoValid returns correct value -- empty input", actual)
	})
}

// ── Bool/Int/Byte/Float converters ──

