package table

import (
	"testing"
)

func TestTable_Hash(t *testing.T) {
	table := Table{}
	table.CreateDefault()

	hash := table.Hash()

	if hash != 312455669 {
		t.Errorf("Hash() = %v, want %v", hash, "312455669")
	}
}
