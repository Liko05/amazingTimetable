package table

import (
	"testing"
)

func TestTable_Hash(t *testing.T) {
	table := Table{}
	table.CreateDefault()

	hash := table.Hash()

	if hash != 1033733252 {
		t.Errorf("Hash() = %v, want %v", hash, "1033733252")
	}
}
