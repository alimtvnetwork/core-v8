package corestrtests

import (
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args")

// ══════════════════════════════════════════════════════════════════════════════
// Hashmap — Segment 5a
// ══════════════════════════════════════════════════════════════════════════════

func Test_HashmapDataModel_Length_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Length", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"len": d.Length(),
			"empty": d.IsEmpty(),
			"hasAny": d.HasAnyItem(),
			"last": d.LastIndex(),
		}

		// Assert
		expected := args.Map{
			"len": 1,
			"empty": false,
			"hasAny": true,
			"last": 0,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff -- basic props", actual)
	})
}

func Test_HashmapDataModel_Length_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Length_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": d.Length()}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff nil -- 0", actual)
	})
}

func Test_HashmapDataModel_Raw_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Raw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": len(d.Raw())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Raw -- 1 item", actual)
	})
}

func Test_HashmapDataModel_Raw_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Raw_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": len(d.Raw())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Raw nil -- empty", actual)
	})
}

func Test_HashmapDataModel_MapAnyItems_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_MapAnyItems", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"len": len(d.MapAnyItems())}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "MapAnyItems -- 1 item", actual)
	})
}

func Test_HashmapDataModel_MapAnyItems_Nil_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_MapAnyItems_Nil", func() {
		// Arrange
		var d *corestr.HashmapDiff

		// Act
		actual := args.Map{"len": len(d.MapAnyItems())}

		// Assert
		expected := args.Map{"len": 0}
		expected.ShouldBeEqual(t, 0, "MapAnyItems nil -- empty", actual)
	})
}

func Test_HashmapDataModel_AllKeysSorted_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_AllKeysSorted", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"b": "2", "a": "1"})
		keys := d.AllKeysSorted()

		// Act
		actual := args.Map{"first": keys[0]}

		// Assert
		expected := args.Map{"first": "a"}
		expected.ShouldBeEqual(t, 0, "AllKeysSorted -- sorted", actual)
	})
}

func Test_HashmapDataModel_IsRawEqual_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_IsRawEqual", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{
			"eq":  d.IsRawEqual(map[string]string{"a": "1"}),
			"neq": d.IsRawEqual(map[string]string{"a": "2"}),
		}

		// Assert
		expected := args.Map{
			"eq": true,
			"neq": false,
		}
		expected.ShouldBeEqual(t, 0, "IsRawEqual -- match and mismatch", actual)
	})
}

func Test_HashmapDataModel_HasAnyChanges_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HasAnyChanges", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"has": d.HasAnyChanges(map[string]string{"a": "2"})}

		// Assert
		expected := args.Map{"has": true}
		expected.ShouldBeEqual(t, 0, "HasAnyChanges -- true", actual)
	})
}

func Test_HashmapDataModel_DiffRaw_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_DiffRaw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.DiffRaw(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": len(diff) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "DiffRaw -- has diff", actual)
	})
}

func Test_HashmapDataModel_HashmapDiffUsingRaw_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HashmapDiffUsingRaw", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.HashmapDiffUsingRaw(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": diff.HasAnyItem()}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw -- has diff", actual)
	})
}

func Test_HashmapDataModel_HashmapDiffUsingRaw_NoDiff_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_HashmapDiffUsingRaw_NoDiff", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		diff := d.HashmapDiffUsingRaw(map[string]string{"a": "1"})

		// Act
		actual := args.Map{"empty": diff.IsEmpty()}

		// Assert
		expected := args.Map{"empty": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiffUsingRaw no diff -- empty", actual)
	})
}

func Test_HashmapDataModel_DiffJsonMessage_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_DiffJsonMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.DiffJsonMessage(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "DiffJsonMessage -- non-empty", actual)
	})
}

func Test_HashmapDataModel_ShouldDiffMessage_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_ShouldDiffMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.ShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "ShouldDiffMessage -- non-empty", actual)
	})
}

func Test_HashmapDataModel_LogShouldDiffMessage_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_LogShouldDiffMessage", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		msg := d.LogShouldDiffMessage("test", map[string]string{"a": "2"})

		// Act
		actual := args.Map{"nonEmpty": msg != ""}

		// Assert
		expected := args.Map{"nonEmpty": true}
		expected.ShouldBeEqual(t, 0, "LogShouldDiffMessage -- non-empty", actual)
	})
}

func Test_HashmapDataModel_ToStringsSliceOfDiffMap_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_ToStringsSliceOfDiffMap", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.ToStringsSliceOfDiffMap(map[string]string{"a": "2"})

		// Act
		actual := args.Map{"hasItems": len(result) > 0}

		// Assert
		expected := args.Map{"hasItems": true}
		expected.ShouldBeEqual(t, 0, "ToStringsSliceOfDiffMap -- non-empty", actual)
	})
}

func Test_HashmapDataModel_RawMapStringAnyDiff_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_RawMapStringAnyDiff", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		result := d.RawMapStringAnyDiff()

		// Act
		actual := args.Map{"len": len(result)}

		// Assert
		expected := args.Map{"len": 1}
		expected.ShouldBeEqual(t, 0, "RawMapStringAnyDiff -- 1 item", actual)
	})
}

func Test_HashmapDataModel_Serialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Serialize", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		b, err := d.Serialize()

		// Act
		actual := args.Map{
			"noErr": err == nil,
			"hasBytes": len(b) > 0,
		}

		// Assert
		expected := args.Map{
			"noErr": true,
			"hasBytes": true,
		}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Serialize -- success", actual)
	})
}

func Test_HashmapDataModel_Deserialize_FromSeg5(t *testing.T) {
	safeTest(t, "Test_Seg5_HMD_Deserialize", func() {
		// Arrange
		d := corestr.HashmapDiff(map[string]string{"a": "1"})
		var dest map[string]string
		err := d.Deserialize(&dest)

		// Act
		actual := args.Map{"noErr": err == nil}

		// Assert
		expected := args.Map{"noErr": true}
		expected.ShouldBeEqual(t, 0, "HashmapDiff Deserialize -- success", actual)
	})
}
