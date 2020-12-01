package computer

import "testing"

func TestArgs(t *testing.T) {
	if Args("10002") != "001" {
		t.Fatalf("expected: %s, got: %s", "001", Args("10002"))
	}
}

func TestOpcode(t *testing.T) {
	if Opcode("10002") != "02" {
		t.Fatalf("expected: %s, got: %s", "02", Opcode("10002"))
	}
}
