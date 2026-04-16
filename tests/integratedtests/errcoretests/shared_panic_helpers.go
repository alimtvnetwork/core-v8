package errcoretests

// callPanicsErrcore calls fn and returns true if it panicked.
// Moved here from ErrorHandling_test.go so split-recovery subfolders can see it.
func callPanicsErrcore(fn func()) bool {
	didPanic := false
	func() {
		defer func() {
			if r := recover(); r != nil {
				didPanic = true
			}
		}()
		fn()
	}()
	return didPanic
}
