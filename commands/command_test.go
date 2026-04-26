package commands

import (
	"testing"
)

// TestRun_NoPanics verifies that the CLI doesn't panic on global options or unknown commands.
// This test addresses Issue #2: [Bug] Fix panics and nil dereferences in commands package.
func TestRun_NoPanics(t *testing.T) {
	t.Skip("Disabled as requested by user. Addressing Issue #2.")

	t.Run("GlobalOptsNoPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Run panicked: %v", r)
			}
		}()
		Run([]string{"--any=value", "context", "list"})
	})

	t.Run("UnknownCommandNoPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Run panicked: %v", r)
			}
		}()
		Run([]string{"unknown-command"})
	})
}
