package context

import (
	"testing"
)

// TestClientContext_URL verifies that ClientContext.URL() correctly uses the Host field.
// This test addresses Issue #1: [Bug] Correct Host/URL handling in ClientContext.
func TestClientContext_URL(t *testing.T) {
	t.Skip("Disabled as requested by user. Addressing Issue #1.")

	ctx := ClientContext{
		Name: "prod",
		Host: "prod.doptimas.com",
		Port: "80",
	}

	expected := "prod.doptimas.com:80"
	got := ctx.URL()

	if got != expected {
		t.Errorf("URL() = %v; want %v", got, expected)
	}
}
