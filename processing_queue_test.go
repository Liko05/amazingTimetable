package main

import (
	"testing"
)

func TestProcessingQueue_Push(t *testing.T) {
	q := ProcessingQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if len(q.Queue) != 3 {
		t.Errorf("Expected queue length to be 3, but got %d", len(q.Queue))
	}
}

func TestProcessingQueue_Pop(t *testing.T) {
	q := ProcessingQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)

	if len(q.Queue) != 3 {
		t.Errorf("Expected queue length to be 3, but got %d", len(q.Queue))
	}

	element := q.Pop()
	if element != 1 {
		t.Errorf("Expected element to be 1, but got %d", element)
	}

	if len(q.Queue) != 2 {
		t.Errorf("Expected queue length to be 2, but got %d", len(q.Queue))
	}
}

func TestProcessingQueue_AddIfBetter(t *testing.T) {
	q := ProcessingQueue{}
	q.BestTable = Table{Score: 1}
	q.AddIfBetter(Table{Score: 2})
	q.AddIfBetter(Table{Score: 3})

	if q.BestTable.Score != 3 {
		t.Errorf("Expected BestTable.Score to be 3, but got %d", q.BestTable.Score)
	}
}

func TestProcessingQueue_AddToBestTables(t *testing.T) {
	q := ProcessingQueue{}
	table1, table2, table3 := Table{Score: 1}, Table{Score: 2}, Table{Score: 3}
	table1.CreateDefault()
	table1.Shuffle()

	table2.CreateDefault()
	table2.Shuffle()

	table3.CreateDefault()
	table3.Shuffle()

	q.AddToBestTables(table1)
	q.AddToBestTables(table3)
	q.AddToBestTables(table2)

	if len(q.BestTables) != 3 {
		t.Errorf("Expected BestTables length to be 3, but got %d", len(q.BestTables))
	}
}

func TestProcessingQueue_AddToBestTablesFailsOnMatchingHashes(t *testing.T) {
	q := ProcessingQueue{}
	q.AddToBestTables(Table{Score: 1})
	q.AddToBestTables(Table{Score: 2})
	q.AddToBestTables(Table{Score: 3})

	if len(q.BestTables) != 1 {
		t.Errorf("Expected BestTables length to be 1, but got %d", len(q.BestTables))
	}
}
