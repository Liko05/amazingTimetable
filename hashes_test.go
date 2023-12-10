package main

import "testing"

func TestTable_Hash(t *testing.T) {
	table := Table{}
	table.CreateDefault()

	hash := table.Hash()

	if hash != "b8c798a74f2bce3c7d6d1dbb05cc196f" {
		t.Errorf("Hash() = %v, want %v", hash, "b8c798a74f2bce3c7d6d1dbb05cc196f")
	}
}
